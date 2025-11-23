package async

import (
	"context"
	"sync"

	"go.uber.org/zap"
)

type Task func(ctx context.Context) error

type WorkerPool struct {
	workers   int
	taskQueue chan Task
	wg        sync.WaitGroup
	ctx       context.Context
	cancel    context.CancelFunc
	logger    *zap.Logger
}

func NewWorkerPool(workers int) *WorkerPool {
	ctx, cancel := context.WithCancel(context.Background())
	logger, _ := zap.NewProduction()

	return &WorkerPool{
		workers:   workers,
		taskQueue: make(chan Task, 100), // buffered channel
		ctx:       ctx,
		cancel:    cancel,
		logger:    logger,
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workers; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()

	wp.logger.Info("worker started", zap.Int("worker_id", id))

	for {
		select {
		case task := <-wp.taskQueue:
			wp.logger.Info("worker executing task", zap.Int("worker_id", id))
			if err := task(wp.ctx); err != nil {
				wp.logger.Error("task execution failed",
					zap.Int("worker_id", id),
					zap.Error(err))
			}
		case <-wp.ctx.Done():
			wp.logger.Info("worker shutting down", zap.Int("worker_id", id))
			return
		}
	}
}

func (wp *WorkerPool) Submit(task Task) {
	select {
	case wp.taskQueue <- task:
		wp.logger.Info("task submitted to queue")
	default:
		wp.logger.Warn("task queue is full, dropping task")
	}
}

func (wp *WorkerPool) Stop() {
	wp.cancel()
	close(wp.taskQueue)
	wp.wg.Wait()
	wp.logger.Info("worker pool stopped")
}
