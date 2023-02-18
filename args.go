package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
	"os"
)

var ErrNotNetflixUrl = fmt.Errorf("VideoURL must be a netflix url")
var ErrDownloadDirInvalid = fmt.Errorf("DownloadDir is invalid")

func mustParse(a *Args) {
	arg.MustParse(a)

	argParser := arg.MustParse(a)
	if err := validateArgs(a); err != nil {
		argParser.Fail(err.Error())
	}
}

func validateArgs(a *Args) error {
	if !isNetflixUrl(a) {
		return ErrNotNetflixUrl
	}

	if err := processDownloadDir(a); err != nil {
		return ErrDownloadDirInvalid
	}

	return nil
}

func processDownloadDir(a *Args) error {
	var path = a.DownloadDir
	var err error
	if path != "." {
		a.DownloadDir = path
		return nil
	}

	if cwd, err := os.Getwd(); err == nil {
		a.DownloadDir = cwd
		return nil
	}

	return err
}

func isNetflixUrl(a *Args) bool {
	return a.VideoURL.Host == "www.netflix.com"
}
