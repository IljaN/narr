package main

import (
	"context"
	"fmt"
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
)

func main() {

	ctx := context.Background()

	// Use the DevTools HTTP/JSON API to manage targets (e.g. pages, webworkers).
	devt := devtool.New("http://127.0.0.1:9222")
	pt, err := devt.Get(ctx, devtool.Page)
	if err != nil {
		pt, err = devt.Create(ctx)
		if err != nil {
			log.Panic(err)
		}
	}

	// Initiate a new RPC connection to the Chrome DevTools Protocol target.
	conn, err := rpcc.DialContext(ctx, pt.WebSocketDebuggerURL)
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close() // Leaving connections open will leak memory.

	c := cdp.NewClient(conn)

	responseReceived, err := c.Network.ResponseReceived(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Create the Network.ResponseReceived() event
	navArgs := page.NewNavigateArgs("https://www.netflix.com/watch/81476645?trackId=272554856")
	_, err = c.Page.Navigate(ctx, navArgs)
	if err != nil {
		log.Panic(err)
	}

	if err = c.Network.Enable(ctx, network.NewEnableArgs()); err != nil {
		log.Fatal(err)
	}

	defer responseReceived.Close()

	// initial queue pool
	q := queue.NewPool(6)
	// shutdown the service and notify all the worker
	// wait all jobs are complete.
	defer q.Release()

	for {
		select {
		case <-responseReceived.Ready():
			ev, err := responseReceived.Recv()
			if err != nil {
				log.Fatal(err)
			}

			// Ignore non-audio urls
			if !isAudioURL(ev.Response.URL) {
				continue
			}

			go func(srcUrl, tgtPath string) {
				if err := q.QueueTask(func(ctx context.Context) error {
					downloadAudio(srcUrl, tgtPath)

					return nil
				}); err != nil {
					panic(err)
				}
			}(ev.Response.URL, "DL-"+strconv.Itoa(rand.Int()))
		}
	}
}

// Audio resources have the path format /range/0-nnnn...
func isAudioURL(u string) bool {
	return strings.Contains(u, "/range/0-")
}

func downloadAudio(fromUrl, toPath string) {
	fmt.Println("Downloading " + fromUrl)
	// We need to remove the path from the audio url to get a downloadable url
	u, err := url.Parse(fromUrl)
	if err != nil {
		log.Fatal(err)
	}

	u.Path = ""

	out, err := os.Create(toPath)

	if err != nil {
		log.Fatal(err)
	}

	defer out.Close()

	resp, err := http.Get(u.String())
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	n, err := io.Copy(out, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Done, got %d bytes", n)
}
