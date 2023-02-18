package main

import (
	"github.com/alexflint/go-arg"
	"github.com/stretchr/testify/assert"
	"net/url"
	"os"
	"testing"
)

func TestArgs(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	// Test args with multiple sub-tests
	cases := []struct {
		name    string
		args    []string
		expArgs *Args
		expErr  error
	}{
		{
			name: "WithDefault",
			args: []string{"cmd", "https://www.netflix.com/watch/81405170?trackId=254015180"},
			expArgs: &Args{
				VideoURL:    parseUrl(t, "https://www.netflix.com/watch/81405170?trackId=254015180"),
				DownloadDir: cwd(t),
				ChromeURL:   parseUrl(t, "http://127.0.0.1:9222"),
			},
		},
		{
			name: "WithCustom",
			args: []string{"cmd", "-c", "http://127.0.0.1:8888", "https://www.netflix.com/watch/81405170?trackId=254015180", "/tmp"},
			expArgs: &Args{
				VideoURL:    parseUrl(t, "https://www.netflix.com/watch/81405170?trackId=254015180"),
				DownloadDir: "/tmp",
				ChromeURL:   parseUrl(t, "http://127.0.0.1:8888"),
			},
		},
		{
			name: "InvalidVideoUrl",
			args: []string{"cmd", "https://www.notnetflix.com/foo"},
			expArgs: &Args{
				VideoURL:    parseUrl(t, "https://www.notnetflix.com/foo"),
				DownloadDir: cwd(t),
				ChromeURL:   parseUrl(t, "http://127.0.0.1:9222"),
			},
			expErr: ErrNotNetflixUrl,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			os.Args = c.args
			a := &Args{}
			_ = arg.Parse(a)
			validateErr := validateArgs(a)
			if validateErr != nil && c.expErr != nil {
				assert.Equal(t, c.expErr.Error(), validateErr.Error())
				return
			}

			if c.expArgs.VideoURL != nil {
				assert.Equal(t, c.expArgs.VideoURL.String(), a.VideoURL.String())
			}

			if c.expArgs.ChromeURL != nil {
				assert.Equal(t, c.expArgs.ChromeURL.String(), a.ChromeURL.String())
			}
			assert.Equal(t, c.expArgs.DownloadDir, a.DownloadDir)

		})
	}

}

func parseUrl(t *testing.T, urlString string) *url.URL {
	u, err := url.Parse(urlString)
	if err != nil {
		t.Fatal(err)
	}

	return u
}

func cwd(t *testing.T) string {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	return wd
}
