package utils

import (
	"context"
	"time"

	"go.uber.org/zap"
)

type RetryConfig struct {
	MaxAttempts  int
	InitialDelay time.Duration
	MaxDelay     time.Duration
	Multiplier   float64
	Logger       *zap.Logger
}

func DefaultRetryConfig() RetryConfig {
	logger, _ := zap.NewProduction()
	return RetryConfig{
		MaxAttempts:  3,
		InitialDelay: 100 * time.Millisecond,
		MaxDelay:     5 * time.Second,
		Multiplier:   2.0,
		Logger:       logger,
	}
}

func Retry(ctx context.Context, config RetryConfig, operation func() error) error {
	var lastErr error
	delay := config.InitialDelay

	for attempt := 1; attempt <= config.MaxAttempts; attempt++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if err := operation(); err != nil {
			lastErr = err
			if attempt == config.MaxAttempts {
				break
			}

			config.Logger.Warn("operation failed, retrying",
				zap.Int("attempt", attempt),
				zap.Int("max_attempts", config.MaxAttempts),
				zap.Duration("delay", delay),
				zap.Error(err),
			)

			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(delay):
			}

			delay = time.Duration(float64(delay) * config.Multiplier)
			if delay > config.MaxDelay {
				delay = config.MaxDelay
			}
		} else {
			return nil
		}
	}

	return lastErr
}
