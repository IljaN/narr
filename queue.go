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

func NewDownloadQueue() *DownloadQueue {
	return &DownloadQueue{
		Queue:      queue.NewPool(8),
		statusMsgs: make(chan DownloadStatus, 50),
	}
}

type DownloadQueue struct {
	statusMsgs chan DownloadStatus
	*queue.Queue
}

type DownloadTask struct {
	SrcURL   string
	ToPath   string
	VideoUrl string
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
func (q *DownloadQueue) QueueDownload(t *DownloadTask) error {
	go func(t *DownloadTask) {
		var taskId = newTaskId()
		q.statusMsgs <- DownloadStatusQueued{
			taskId: taskId,
			task:   t,
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
				return fmt.Errorf("error probing file format: %w", err)
			}

			if !isAudio {
				return nil
			}

			q.statusMsgs <- DownloadStatusStarted{
				taskId: taskId,
				task:   t,
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

			q.statusMsgs <- DownloadStatusFinished{
				taskId:        taskId,
				task:          t,
				BytesReceived: n,
				Duration:      time.Since(start),
			}

			return nil
		})
		if err != nil {
			return
		}

	}(t)

	return nil
}

type DownloadStatus interface {
	TaskId() string
	Task() *DownloadTask
}

type DownloadStatusQueued struct {
	taskId string
	task   *DownloadTask
}

func (s DownloadStatusQueued) TaskId() string {
	return s.taskId
}

func (s DownloadStatusQueued) Task() *DownloadTask {
	return s.task
}

type DownloadStatusStarted struct {
	taskId string
	task   *DownloadTask
}

func (s DownloadStatusStarted) TaskId() string {
	return s.taskId
}

func (s DownloadStatusStarted) Task() *DownloadTask {
	return s.task
}

type DownloadStatusFinished struct {
	taskId        string
	BytesReceived int64
	Duration      time.Duration
	task          *DownloadTask
}

func (s DownloadStatusFinished) TaskId() string {
	return s.taskId
}

func (s DownloadStatusFinished) Task() *DownloadTask {
	return s.task
}

func newTaskId() string {
	bytes := make([]byte, 6)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(bytes)
}
