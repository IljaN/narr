package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/golang-queue/queue"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type DownloadQueue struct {
	statusMsgs chan DownloadStatus
	*queue.Queue
}

func NewDownloadQueue() *DownloadQueue {
	return &DownloadQueue{
		Queue:      queue.NewPool(8),
		statusMsgs: make(chan DownloadStatus, 50),
	}
}

type DownloadTask struct {
	SrcURL   string
	ToPath   string
	VideoUrl string
}

type DownloadStatus interface {
	TaskId() string
	Task() DownloadTask
}

func (q *DownloadQueue) OnStatusReceived(f func(DownloadStatus)) {
	go func() {
		for msg := range q.statusMsgs {
			f(msg)
		}
	}()
}

// QueueDownload downloads the audio file from the given url and saves it to the given path. As Netflix has no metadata for
// its media files, we need to probe the file format to determine if it's an audio file or a video file. This is done by
// reading the first 3000 bytes of the file and checking if it's an audio file. If it's a video file, we skip it.
func (q *DownloadQueue) QueueDownload(t DownloadTask) error {
	go func(t DownloadTask) {
		var taskId = newTaskId()
		q.statusMsgs <- Queuing{
			&taskInfo{taskId, t},
		}
		err := q.QueueTask(func(ctx context.Context) error {
			start := time.Now()
			var fromUrl = t.SrcURL
			var toPath = t.ToPath
			resp, err := http.Get(fromUrl)
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			// Probe file format
			header := make([]byte, 3000)
			isAudio, err := probeFileFormat(resp.Body, header)
			if err != nil {
				return fmt.Errorf("error probing file format of %s: %w", fromUrl, err)
			}

			if !isAudio {
				return nil
			}

			q.statusMsgs <- Begin{
				&taskInfo{taskId, t},
			}

			out, err := os.Create(toPath)

			if err != nil {
				return err
			}

			defer out.Close()

			_, err = out.Write(header)
			if err != nil {
				return err
			}

			n, err := io.Copy(out, resp.Body)
			if err != nil {
				return err
			}

			q.statusMsgs <- Finished{
				taskInfo:      &taskInfo{taskId, t},
				bytesReceived: n,
				duration:      time.Since(start),
			}

			return nil
		})
		if err != nil {
			return
		}

	}(t)

	return nil
}

type taskInfo struct {
	taskId string
	task   DownloadTask
}

func (d *taskInfo) TaskId() string {
	return d.taskId
}

func (d *taskInfo) Task() DownloadTask {
	return d.task
}

type Queuing struct {
	*taskInfo
}

type Begin struct {
	*taskInfo
}

type Finished struct {
	bytesReceived int64
	duration      time.Duration

	*taskInfo
}

func (f *Finished) BytesReceived() int64 {
	return f.bytesReceived
}

func (f *Finished) Duration() time.Duration {
	return f.duration
}

func newTaskId() string {
	bytes := make([]byte, 6)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(bytes)
}
