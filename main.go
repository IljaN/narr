package main

import (
	"context"
	"fmt"
	"github.com/cenkalti/backoff/v4"
	"github.com/golang-queue/queue"
	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/devtool"
	"github.com/mafredri/cdp/protocol/network"
	"github.com/mafredri/cdp/protocol/page"
	"github.com/mafredri/cdp/rpcc"
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

	// Enable event stream
	if err = chrome.Network.Enable(ctx, network.NewEnableArgs()); err != nil {
		log.Fatal(err)
	}

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

	for u := range listen(responseReceived) {
		if !isAudioURL(u) {
			continue
		}

		err := enqueueDownload(q, toDownloadableURL(u), toDownloadPath(args.VideoURL, args.DownloadDir))
		if err != nil {
			log.Println(err)
		}
	}
}

// listen to all responses received by the current tab and send us their URLs.
func listen(responseReceived network.ResponseReceivedClient) chan string {
	urls := make(chan string)
	go func() {
		for {
			select {
			case <-responseReceived.Ready():
				ev, err := responseReceived.Recv()
				if err != nil {
					log.Fatal(err)
				}

				urls <- ev.Response.URL
			}
		}
	}()

	return urls
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

// Audio resources have the path format /range/0-nnnn...
func isAudioURL(u string) bool {
	return strings.Contains(u, "/range/0-")
}

// toDownloadableURL removes the path from the url to make the resource downloadable. In our case the path
// always contains a download-range in bytes which we can discard. See isAudioURL.
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
func toDownloadPath(videoURL *url.URL, downloadDir string) string {
	if strings.HasPrefix(videoURL.Path, "/watch") && videoURL.Query().Has("trackId") {
		videoId := strings.TrimLeft(videoURL.Path, "/watch/")
		trackId := videoURL.Query().Get("trackId")
		return downloadDir + "/" + videoId + "-" + trackId + "-" + strconv.Itoa(rand.Int())
	}

	return downloadDir + "/" + "DL-" + strconv.Itoa(rand.Int())
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
	log.Println("Downloading " + fromUrl)
	out, err := os.Create(toPath)

	if err != nil {
		return err
	}

	defer out.Close()

	resp, err := http.Get(fromUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	n, err := io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	log.Printf("Finished %s as  %s, got %d bytes", fromUrl, toPath, n)

	return nil
}
