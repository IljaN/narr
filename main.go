package main

import (
	"context"
	"fmt"
	"github.com/alfg/mp4"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type Args struct {
	VideoURL    *url.URL `arg:"positional,required" help:"url of the video to download audio from. Must be a netflix url. e.g. https://www.netflix.com/watch/12345678?trackId=12345678"`
	DownloadDir string   `arg:"positional" default:"." help:"directory where to download the audio files. Defaults to current working directory."`
	ChromeURL   *url.URL `arg:"-c, --chrome-url" default:"http://127.0.0.1:9222" help:"url of the chrome debugger."`
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
		case "mediaUrlReceived":
			err := q.QueueDownload(DownloadTask{
				SrcURL:   toDownloadableURL(string(events.payload)),
				ToPath:   toDownloadPath(browserURL, args.DownloadDir),
				VideoUrl: browserURL,
			})
			if err != nil {
				log.Println(err)
			}
		case "navigated":
			log.Printf("🗺Navigate to %s \n", events.payload)
			browserURL = string(events.payload)
		}
	}
}

// download downloads the audio file from the given url and saves it to the given path. As Netflix has no metadata for
// its media files, we need to probe the file format to determine if it's an audio file or a video file. This is done by
// reading the first 3000 bytes of the file and checking if it's an audio file. If it's a video file, we skip it.
func download(fromUrl, toPath, videoUrl string) error {
	resp, err := http.Get(fromUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Probe file format
	header := make([]byte, 3000)
	isAudio, err := probeFileFormat(resp.Body, header)
	if err != nil {
		return fmt.Errorf("error probing file format: %w", err)
	}

	if !isAudio {
		return nil
	}
	log.Printf("▼ Downloading %s  ⟾  %s", videoUrl, toPath)

	out, err := os.Create(toPath)

	if err != nil {
		return err
	}

	defer out.Close()

	_, err = out.Write(header)
	if err != nil {
		return err
	}

	n, err := io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	log.Printf("✓ Finished    %s  ⟾  %s, got %d bytes", videoUrl, toPath, n)

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
