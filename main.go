package main

import (
	"context"
	"log"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
)

type Args struct {
	VideoURL    *url.URL `arg:"positional,required" help:"url of the video to download audio from. Must be a netflix url. e.g. https://www.netflix.com/watch/12345678?trackId=12345678"`
	DownloadDir string   `arg:"positional" default:"." help:"directory where to download the audio files. Defaults to current working directory."`
	ChromeURL   *url.URL `arg:"-c, --chrome-url" default:"http://127.0.0.1:9222" help:"url of the chrome debugger."`
}

var Version string

func (Args) Version() string {
	return Version
}

func main() {
	args := &Args{}
	mustParse(args)

	// Connect to Chrome debugger, retry until success
	ctx := context.Background()
	chromeURL := args.ChromeURL.String()
	chrome := tryConnectToChromeUntilSuccess(ctx, chromeURL)

	log.Printf("Ꙫ Sucessfully connected to Chrome at %s", chromeURL)

	// Create task queue for downloading
	q := NewDownloadQueue()
	defer q.Release()

	// Listen for download status updates
	q.OnStatusReceived(func(status DownloadStatus) {
		task := status.Task()
		switch s := status.(type) {
		case Queuing:
			break
		case Begin:
			log.Printf("▼ [%s] Downloading %s to %s", s.TaskId(), task.VideoUrl, task.FullFilePath)
		case Finished:
			log.Printf("✓ [%s] Finished %s  ⟾  %s (got %d bytes in %f)", s.TaskId(), task.VideoUrl, task.FullFilePath, s.BytesReceived(), s.Duration().Seconds())
		}
	})

	nflx := NewNFLX(chrome)

	// Navigate to the initial url
	err := nflx.NavigateTo(ctx, args.VideoURL.String())
	if err != nil {
		log.Fatal(err)
	}

	// Listen for received media urls and queue them for download. Also listen for navigated events to update the current url
	var browserURL = args.VideoURL.String()
	for events := range nflx.Listen(ctx) {
		switch events.evType {
		case MediaUrlReceivedEvent:
			err := q.QueueDownload(DownloadTask{
				SrcURL:      toDownloadableURL(string(events.payload)),
				VideoUrl:    browserURL,
				DownloadDir: args.DownloadDir,
			})

			if err != nil {
				log.Println(err)
			}
		case NavigatedEvent:
			log.Printf("ᐅ Navigate to %s \n", events.payload)
			browserURL = string(events.payload)
		}
	}
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
func toDownloadPath(videoURL string, downloadDir string, fi probeInfo) string {
	u, err := url.Parse(videoURL)
	if err != nil {
		log.Fatal(err)
	}

	var audCodec = ".aac"
	if fi.isXHEAAC {
		audCodec = ".xhe-aac"
	}

	if strings.HasPrefix(u.Path, "/watch") && u.Query().Has("trackId") {
		videoId := strings.TrimLeft(u.Path, "/watch/")
		trackId := u.Query().Get("trackId")
		return downloadDir + "/" + videoId + "-" + trackId + "-" + strconv.Itoa(rand.Int()) + audCodec + ".mp4a"
	}

	return downloadDir + "/" + "DL-" + strconv.Itoa(rand.Int()) + audCodec + ".mp4a"
}
