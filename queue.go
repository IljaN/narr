package main

import (
	"context"
	"github.com/golang-queue/queue"
)

func NewDownloadQueue() *DownloadQueue {
	return &DownloadQueue{
		Queue: queue.NewPool(8),
	}
}

type DownloadQueue struct {
	*queue.Queue
}

type DownloadTask struct {
	SrcURL   string
	ToPath   string
	VideoUrl string
}

func (q *DownloadQueue) QueueDownload(t DownloadTask) error {
	go func(t DownloadTask) {
		err := q.QueueTask(func(ctx context.Context) error {
			return download(t.SrcURL, t.ToPath, t.VideoUrl)
		})
		if err != nil {
			return
		}

	}(t)

	return nil
}
