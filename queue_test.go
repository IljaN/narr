package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func FileServerHandler(directory string) http.Handler {
	// Use http.FileServer to serve the directory
	fs := http.FileServer(http.Dir(directory))
	return http.StripPrefix("/files/", fs)
}

func TestDownloadQueue_QueueDownload_Formats(t *testing.T) {
	// Create a test server with the file server handler
	ts := httptest.NewServer(FileServerHandler("./test/testdata"))
	q := NewDownloadQueue()

	t.Run("ISO6Download", func(t *testing.T) {
		toPath := t.TempDir() + "/audio_iso6.mp4"
		task := DownloadTask{
			SrcURL:   ts.URL + "/files/audio_iso6.mp4",
			ToPath:   toPath,
			VideoUrl: "http://example.com/video.mp4",
		}
		err := q.QueueDownload(task)
		assert.NoError(t, err)

		// Check every 50 ms for maximum of 5 seconds if the download has been queued, begun and finished, if not fail the test
		for i := 0; i < 100; i++ {
			time.Sleep(50 * time.Millisecond)
			if _, err := os.Stat(toPath); errors.Is(err, os.ErrNotExist) {
				continue
			}
			break

		}

		assert.FileExists(t, toPath)
		assertIdenticalFiles(t, "./test/testdata/audio_iso6.mp4", toPath)
	})

	t.Run("MP42Download", func(t *testing.T) {
		toPath := t.TempDir() + "/audio_mp42.mp4"
		task := DownloadTask{
			SrcURL:   ts.URL + "/files/audio_mp42.mp4",
			ToPath:   toPath,
			VideoUrl: "http://example.com/video.mp4",
		}
		err := q.QueueDownload(task)
		assert.NoError(t, err)

		for i := 0; i < 100; i++ {
			time.Sleep(50 * time.Millisecond)
			if _, err := os.Stat(toPath); errors.Is(err, os.ErrNotExist) {
				continue
			}
			break
		}

		assert.FileExists(t, toPath)
		assertIdenticalFiles(t, "./test/testdata/audio_mp42.mp4", toPath)
	})

	t.Run("InvalidFile", func(t *testing.T) {
		toPath := t.TempDir() + "/vid.mp4"
		task := DownloadTask{
			SrcURL:   ts.URL + "/files/vid.mp4",
			ToPath:   toPath,
			VideoUrl: "http://example.com/video.mp4",
		}
		err := q.QueueDownload(task)
		assert.NoError(t, err)

		for i := 0; i < 5; i++ {
			time.Sleep(1 * time.Second)
			assert.NoFileExists(t, toPath)
		}
	})

	t.Cleanup(func() {
		q.Release()
		ts.Close()
	})
}

func TestDownloadQueue_Statuses(t *testing.T) {
	// Create a test server with the file server handler
	testServer := httptest.NewServer(FileServerHandler("./test/testdata"))
	defer testServer.Close()

	q := NewDownloadQueue()
	defer q.Release()

	hasQueued, hasBegun, hasFinished := false, false, false
	_, _, _ = hasQueued, hasBegun, hasFinished

	// Listen for download status updates
	q.OnStatusReceived(func(status DownloadStatus) {
		switch status.(type) {
		case Queuing:
			hasQueued = true
		case Begin:
			hasBegun = true
		case Finished:
			hasFinished = true
		}
	})

	task := DownloadTask{
		SrcURL:   testServer.URL + "/files/audio_iso6.mp4",
		ToPath:   t.TempDir() + "/audio_iso6.mp4",
		VideoUrl: "http://example.com/video.mp4",
	}

	err := q.QueueDownload(task)
	assert.NoError(t, err)

	// Check every 50 ms for maximum of 5 seconds if the download has been queued, begun and finished, if not fail the test
	for i := 0; i < 100; i++ {
		if hasQueued && hasBegun && hasFinished {
			break
		}
		time.Sleep(50 * time.Millisecond)
	}

	assert.True(t, hasQueued)
	assert.True(t, hasBegun)
	assert.True(t, hasFinished)
}

// assertIdenticalFiles checks if two files are identical.
func assertIdenticalFiles(t *testing.T, filePath1, filePath2 string) bool {
	// Open the first file
	file1, err := os.Open(filePath1)
	if err != nil {
		return assert.Fail(t, "failed to open file1: %v", err)
	}
	defer file1.Close()

	// Open the second file
	file2, err := os.Open(filePath2)
	if err != nil {
		return assert.Fail(t, "failed to open file1: %v", err)

	}
	defer file2.Close()

	// Get the file info for both files
	fileInfo1, err := file1.Stat()
	if err != nil {
		return assert.Fail(t, "failed to stat file1: %v", err)
	}
	fileInfo2, err := file2.Stat()
	if err != nil {
		return assert.Fail(t, "failed to stat file2: %v", err)
	}

	// Compare sizes
	if fileInfo1.Size() != fileInfo2.Size() {
		return assert.Fail(t, "file sizes are different: %d != %d", fileInfo1.Size(), fileInfo2.Size())
	}

	// Compare content using MD5 hash
	hash1, err := calculateMD5(file1)
	if err != nil {
		return assert.Fail(t, "failed to calculate hash for file1: %v", err)
	}
	hash2, err := calculateMD5(file2)
	if err != nil {
		return assert.Fail(t, "failed to calculate hash for file2: %v", err)
	}

	// Compare the hashes
	return hash1 == hash2
}

// calculateMD5 calculates the MD5 hash of a file
func calculateMD5(file *os.File) (string, error) {
	// Reset the file pointer to the beginning
	_, err := file.Seek(0, io.SeekStart)
	if err != nil {
		return "", fmt.Errorf("failed to reset file pointer: %v", err)
	}

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", fmt.Errorf("failed to hash file: %v", err)
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
