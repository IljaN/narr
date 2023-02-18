package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestToDownloadPath(t *testing.T) {
	downloadDir := "/tmp/foo"

	videoUrl := "https://www.netflix.com/watch/81405170?trackId=254015180"
	downloadPath := toDownloadPath(videoUrl, downloadDir)
	assert.True(t, strings.HasPrefix(downloadPath, "/tmp/foo/81405170-254015180-"))

	videoUrl = "https://www.netflix.com"
	downloadPath = toDownloadPath(videoUrl, downloadDir)
	assert.True(t, strings.HasPrefix(downloadPath, "/tmp/foo/DL-"))
}
