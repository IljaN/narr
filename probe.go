package main

import (
	"bytes"
	"github.com/abema/go-mp4"
)

type probeInfo struct {
	isAudio    bool
	majorBrand string
	isXHEAAC   bool
}

// probeFileFormat reads the header of file and checks if it is an audio file and returns some info about it.
func probeFileFormat(header []byte) (isAudio bool, pInfo probeInfo, err error) {
	const xHE_AAC_AudOTI = 42 //  Audio Object Type Identifier (OTI) for xHE-AAC

	pi := probeInfo{}

	info, err := mp4.Probe(bytes.NewReader(header))
	if err != nil {
		return false, pi, err
	}

	majorBrand := string(info.MajorBrand[:])

	// On macOS xHE-AAC is delivered (sometimes?) in an iso6 container
	if (majorBrand == "mp42" || majorBrand == "iso6") && len(info.Tracks) == 1 && info.Tracks[0].Codec == mp4.CodecMP4A {
		pi.isAudio = true
		pi.isXHEAAC = info.Tracks[0].MP4A.AudOTI == xHE_AAC_AudOTI
		pi.majorBrand = majorBrand
	}

	return pi.isAudio, pi, nil
}
