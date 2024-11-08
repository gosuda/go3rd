package concurrency

import (
	"context"
	"sync/atomic"
	"testing"
	"time"
)

func TestFirst(t *testing.T) {
	t.Run("string functions", func(t *testing.T) {
		f1 := func() string { time.Sleep(20 * time.Millisecond); return "first" }
		f2 := func() string { time.Sleep(10 * time.Millisecond); return "second" }
		f3 := func() string { time.Sleep(30 * time.Millisecond); return "third" }

		result := First(f1, f2, f3)
		if result == "" {
			t.Error("expected non-empty string result")
		}
	})

	t.Run("integer functions", func(t *testing.T) {
		f1 := func() int { time.Sleep(20 * time.Millisecond); return 1 }
		f2 := func() int { time.Sleep(10 * time.Millisecond); return 2 }
		f3 := func() int { time.Sleep(30 * time.Millisecond); return 3 }

		result := First(f1, f2, f3)
		if result == 0 {
			t.Error("expected non-zero integer result")
		}
	})
}

func TestFirstCtx(t *testing.T) {
	t.Run("successful completion", func(t *testing.T) {
		ctx := context.Background()
		f1 := func(ctx context.Context) string { time.Sleep(20 * time.Millisecond); return "first" }
		f2 := func(ctx context.Context) string { time.Sleep(10 * time.Millisecond); return "second" }

		result, err := FirstCtx(ctx, f1, f2)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if result == "" {
			t.Error("expected non-empty string result")
		}
	})

	t.Run("context cancellation", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		f1 := func(ctx context.Context) string {
			time.Sleep(100 * time.Millisecond)
			return "first"
		}
		f2 := func(ctx context.Context) string {
			time.Sleep(100 * time.Millisecond)
			return "second"
		}

		go func() {
			time.Sleep(10 * time.Millisecond)
			cancel()
		}()

		_, err := FirstCtx(ctx, f1, f2)
		if err == nil {
			t.Error("expected context cancellation error")
		}
	})

	t.Run("integer functions", func(t *testing.T) {
		ctx := context.Background()
		f1 := func(ctx context.Context) int { time.Sleep(20 * time.Millisecond); return 1 }
		f2 := func(ctx context.Context) int { time.Sleep(10 * time.Millisecond); return 2 }

		result, err := FirstCtx(ctx, f1, f2)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if result == 0 {
			t.Error("expected non-zero integer result")
		}
	})

	t.Run("cancellation after first successful completion", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		var val atomic.Int64

		f1 := func(ctx context.Context) int {
			after := time.After(100 * time.Millisecond)
			select {
			case <-after:
				val.Store(1)
				return 1
			case <-ctx.Done():
				return -1
			}
		}

		f2 := func(ctx context.Context) int {
			after := time.After(200 * time.Millisecond)
			select {
			case <-after:
				val.Store(2)
				return 2
			case <-ctx.Done():
				return -2
			}
		}

		result, err := FirstCtx(ctx, f1, f2)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		cancel()

		time.Sleep(200 * time.Millisecond)
		if result != 1 {
			t.Error("expected successful completion")
		}
		if val.Load() != 1 {
			t.Error("expected f1 to complete before cancellation")
		}
	})
}
