package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestToDownloadPath(t *testing.T) {
	downloadDir := "/tmp/foo"

	videoUrl := "https://www.netflix.com/watch/81405170?trackId=254015180"
	downloadPath := toDownloadPath(videoUrl, downloadDir, probeInfo{
		isAudio:    true,
		majorBrand: "mp42",
		isXHEAAC:   false,
	})
	assert.True(t, strings.HasPrefix(downloadPath, "/tmp/foo/81405170-254015180-"))
	assert.True(t, strings.HasSuffix(downloadPath, ".mp4a"))

	videoUrl = "https://www.netflix.com"
	downloadPath = toDownloadPath(videoUrl, downloadDir, probeInfo{
		isAudio:    true,
		majorBrand: "mp42",
		isXHEAAC:   false,
	})
	assert.True(t, strings.HasPrefix(downloadPath, "/tmp/foo/DL-"))
	assert.True(t, strings.HasSuffix(downloadPath, ".mp4a"))
}

func TestToDownloadableURL(t *testing.T) {
	mediaURL := "https://ipv6-cndoca-avcd.nflxvideo.net/range/0-52342?t=12345"
	downloadableURL := toDownloadableURL(mediaURL)
	assert.Equal(t, "https://ipv6-cndoca-avcd.nflxvideo.net?t=12345", downloadableURL)
}
