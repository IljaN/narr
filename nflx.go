package main

import (
	"context"
	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/protocol/network"
	"github.com/mafredri/cdp/protocol/page"
	"golang.org/x/sync/errgroup"
	"log"
	"strings"
)

type NFlX struct {
	chrome *cdp.Client
}

func NewNFLX(chrome *cdp.Client) *NFlX {
	return &NFlX{
		chrome: chrome,
	}
}

type event struct {
	evType  int
	payload []byte
}

const MediaUrlReceivedEvent = 0
const NavigatedEvent = 1

func (n *NFlX) Listen(ctx context.Context) chan event {
	c := n.chrome // Listen to response received events
	responseReceived, err := c.Network.ResponseReceived(ctx)
	if err != nil {
		log.Fatal(err)
	}

	navigated, err := c.Page.NavigatedWithinDocument(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if err = runBatch(
		// Enable all the domain events that we're interested in.
		func() error { return c.Network.Enable(ctx, network.NewEnableArgs()) },
		func() error { return c.Page.Enable(ctx) },
	); err != nil {
		log.Fatal(err)
	}

	ev := func(navigated page.NavigatedWithinDocumentClient, responseReceived network.ResponseReceivedClient) chan event {
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

					events <- event{NavigatedEvent, []byte(ev.URL)}

				case <-responseReceived.Ready():
					ev, err := responseReceived.Recv()
					if err != nil {
						log.Fatal(err)
					}

					if isMediaURL(ev.Response.URL) {
						events <- event{MediaUrlReceivedEvent, []byte(ev.Response.URL)}
					}
				}
			}
		}()
		return events
	}(navigated, responseReceived)

	return ev
}

// Media resources have the path format /range/0-nnnn...
func isMediaURL(u string) bool {
	return strings.Contains(u, "/range/0-")
}

func (n *NFlX) NavigateTo(ctx context.Context, url string) error {
	navArgs := page.NewNavigateArgs(url)
	_, err := n.chrome.Page.Navigate(ctx, navArgs)
	if err != nil {
		return err
	}

	return nil
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
