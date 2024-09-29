package main

import (
	"github.com/stretchr/testify/assert"
	"os"
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

func TestProbeFileFormat(t *testing.T) {
	testCases := []struct {
		filePath string
		isAudio  bool
	}{
		{"test/testdata/audio_iso6.mp4", true},
		{"test/testdata/audio_mp42.mp4", true},
		{"test/testdata/vid.mp4", false},
	}

	for _, tc := range testCases {
		t.Run(tc.filePath, func(t *testing.T) {
			data, err := os.ReadFile(tc.filePath)
			if err != nil {
				t.Fatal(err)
			}

			isAudio, err := probeFileFormat(data)
			assert.Nil(t, err)
			assert.Equal(t, tc.isAudio, isAudio)
		})
	}
}
