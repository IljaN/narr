package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestProbeFileFormat_IsAudio(t *testing.T) {
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

			isAudio, _, err := probeFileFormat(data)
			assert.Nil(t, err)
			assert.Equal(t, tc.isAudio, isAudio)
		})
	}
}

func TestProbeFileFormat_Info(t *testing.T) {
	data, err := os.ReadFile("test/testdata/audio_iso6.mp4")
	if err != nil {
		t.Fatal(err)
	}

	_, info, err := probeFileFormat(data)

	assert.NoError(t, err)
	assert.Equal(t, true, info.isXHEAAC)
	assert.Equal(t, "iso6", info.majorBrand)
	assert.Equal(t, true, info.isAudio)

	data, err = os.ReadFile("test/testdata/audio_mp42.mp4")
	if err != nil {
		t.Fatal(err)
	}

	_, info, err = probeFileFormat(data)

	assert.NoError(t, err)
	assert.Equal(t, false, info.isXHEAAC)
	assert.Equal(t, "mp42", info.majorBrand)
	assert.Equal(t, true, info.isAudio)

	data, err = os.ReadFile("test/testdata/vid.mp4")
	if err != nil {
		t.Fatal(err)
	}

	_, info, err = probeFileFormat(data)

	assert.NoError(t, err)
	assert.Equal(t, false, info.isAudio)

}
