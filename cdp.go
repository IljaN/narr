package main

import (
	"context"
	"fmt"
	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/devtool"
	"github.com/mafredri/cdp/rpcc"
	"log"
	"time"
)

func tryConnectToChromeUntilSuccess(ctx context.Context, url string) *cdp.Client {
	var err error
	var client *cdp.Client
	for {
		client, err = connectToChromeDebugger(ctx, url)
		if err == nil {
			return client
		}
		log.Print(fmt.Errorf("can't connect to %s Chrome must be started in debug mode. Retrying every 5 seconds. %w", url, err))
		time.Sleep(5 * time.Second)
	}
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
