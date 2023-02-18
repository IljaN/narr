package main

import (
	"context"
	"fmt"
	"github.com/alfg/mp4"
	"github.com/cenkalti/backoff/v4"
	"github.com/golang-queue/queue"
	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/devtool"
	"github.com/mafredri/cdp/protocol/network"
	"github.com/mafredri/cdp/protocol/page"
	"github.com/mafredri/cdp/rpcc"
	"golang.org/x/sync/errgroup"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type Args struct {
	VideoURL    *url.URL `arg:"positional,required" help:"url of the video to download audio from. Must be a netflix url. e.g. https://www.netflix.com/watch/12345678?trackId=12345678"`
	DownloadDir string   `arg:"positional" default:"." help:"directory where to download the audio files. Defaults to current working directory."`
	ChromeURL   *url.URL `arg:"-c, --chrome-url" default:"http://127.0.0.1:9222" help:"url of the chrome debugger."`
}

func main() {
	args := &Args{}
	mustParse(args)

	ctx := context.Background()
	var chrome *cdp.Client

	retryFunc := func() error {
		var err error
		chromeURL := args.ChromeURL.String()
		chrome, err = connectToChromeDebugger(ctx, chromeURL)
		if err != nil {
			log.Print(fmt.Errorf("can't connect to %s. Chrome must be started in debug mode. %w", chromeURL, err))
		}
		return err
	}

	err := backoff.Retry(retryFunc, backoff.NewConstantBackOff(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	// Listen to response received events
	responseReceived, err := chrome.Network.ResponseReceived(ctx)
	if err != nil {
		log.Fatal(err)
	}

	navigated, err := chrome.Page.NavigatedWithinDocument(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if err = runBatch(
		// Enable all the domain events that we're interested in.
		func() error { return chrome.Network.Enable(ctx, network.NewEnableArgs()) },
		func() error { return chrome.Page.Enable(ctx) },
	); err != nil {
		log.Fatal(err)
		return
	}

	defer navigated.Close()

	// Open netflix tab
	navArgs := page.NewNavigateArgs(args.VideoURL.String())
	_, err = chrome.Page.Navigate(ctx, navArgs)
	if err != nil {
		log.Fatal(err)
	}

	defer responseReceived.Close()

	// Initial queue pool for download jobs
	q := queue.NewPool(8)
	defer q.Release()

	var videoURL = args.VideoURL.String()
	for events := range listen(navigated, responseReceived) {
		if events.evType == "mediaUrlReceived" {
			err := enqueueDownload(q, toDownloadableURL(string(events.payload)), toDownloadPath(videoURL, args.DownloadDir))
			if err != nil {
				log.Println(err)
			}
		}

		if events.evType == "navigated" {
			log.Printf("Navigate to %s \n", events.payload)
			videoURL = string(events.payload)
		}
	}

}

type event struct {
	evType  string
	payload []byte
}

// listen to all responses received by the current tab and send us their URLs.
func listen(navigated page.NavigatedWithinDocumentClient, responseReceived network.ResponseReceivedClient) chan event {
	events := make(chan event)

	go func() {
		defer navigated.Close()
		defer responseReceived.Close()
		for {
			select {
			case <-navigated.Ready():
				ev, err := navigated.Recv()
				if err != nil {
					log.Fatal(err)
				}

				events <- event{
					evType:  "navigated",
					payload: []byte(ev.URL),
				}

			case <-responseReceived.Ready():
				ev, err := responseReceived.Recv()
				if err != nil {
					log.Fatal(err)
				}

				if isMediaURL(ev.Response.URL) {
					events <- event{
						evType:  "mediaUrlReceived",
						payload: []byte(ev.Response.URL),
					}
				}
			}
		}
	}()

	return events
}

// connectToChromeDebugger establishes a debugging session on a remote chrome instance. Chrome must be already started in debug-mode.
// See https://blog.chromium.org/2011/05/remote-debugging-with-chrome-developer.html for more details
func connectToChromeDebugger(ctx context.Context, url string) (*cdp.Client, error) {
	// Use the DevTools HTTP/JSON API to manage targets (e.g. pages, webworkers).
	devt := devtool.New(url)
	pt, err := devt.Get(ctx, devtool.Page)
	if err != nil {
		pt, err = devt.Create(ctx)
		if err != nil {
			return nil, err
		}
	}

	// Initiate a new RPC connection to the Chrome DevTools Protocol target.
	conn, err := rpcc.DialContext(ctx, pt.WebSocketDebuggerURL)
	if err != nil {
		return nil, err
	}

	return cdp.NewClient(conn), nil
}

func enqueueDownload(q *queue.Queue, fromURL, toPath string) error {
	go func(s, t string) {
		err := q.QueueTask(func(ctx context.Context) error {
			return download(fromURL, toPath)
		})
		if err != nil {
			return
		}

	}(fromURL, toPath)

	return nil
}

func download(fromUrl, toPath string) error {
	resp, err := http.Get(fromUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	header := make([]byte, 3000)
	isAudio, err := probeFileFormat(resp.Body, header)
	if err != nil {
		return fmt.Errorf("error probing file format: %w", err)
	}

	if !isAudio {
		log.Printf("skipping video file %s", fromUrl)
		return nil
	}

	log.Printf("Downloading %s", fromUrl)
	out, err := os.Create(toPath)

	if err != nil {
		return err
	}

	defer out.Close()

	nHdr, err := out.Write(header)
	if err != nil {
		return err
	}

	log.Printf("Written %d bytes of header data", nHdr)

	n, err := io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	log.Printf("Finished %s as  %s, got %d bytes", fromUrl, toPath, n)

	return nil
}

// probeFileFormat reads the first 30k bytes of the file and checks if it is an audio file.
func probeFileFormat(body io.Reader, header []byte) (isAudio bool, err error) {
	_, err = body.Read(header)
	if err != nil {
		return false, fmt.Errorf("error reading header: %w", err)
	}

	rd, err := mp4.OpenFromBytes(header)
	if err != nil {
		return false, err
	}

	return rd.Ftyp.MajorBrand == "mp42", nil
}

// toDownloadableURL removes the path from the payload to make the resource downloadable. In our case the path
// always contains a download-range in bytes which we can discard. See isMediaURL.
func toDownloadableURL(audioURL string) string {
	// We need to remove the path from the audio url to get a downloadable url
	u, err := url.Parse(audioURL)
	if err != nil {
		log.Fatal(err)
	}

	u.Path = ""
	return u.String()
}

// toDownloadPath returns the path where to download the audio file. The path is composed of the video id, the track
// id and a random number.
func toDownloadPath(videoURL string, downloadDir string) string {
	u, err := url.Parse(videoURL)
	if err != nil {
		log.Fatal(err)
	}
	if strings.HasPrefix(u.Path, "/watch") && u.Query().Has("trackId") {
		videoId := strings.TrimLeft(u.Path, "/watch/")
		trackId := u.Query().Get("trackId")
		return downloadDir + "/" + videoId + "-" + trackId + "-" + strconv.Itoa(rand.Int())
	}

	return downloadDir + "/" + "DL-" + strconv.Itoa(rand.Int())
}

// Media resources have the path format /range/0-nnnn...
func isMediaURL(u string) bool {
	return strings.Contains(u, "/range/0-")
}

// runBatchFunc is the function signature for runBatch.
type runBatchFunc func() error

// runBatch runs all functions simultaneously and waits until
// execution has completed or an error is encountered.
func runBatch(fn ...runBatchFunc) error {
	eg := errgroup.Group{}
	for _, f := range fn {
		eg.Go(f)
	}
	return eg.Wait()
}
