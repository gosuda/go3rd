package concurrency

import "context"

func First[T any](functions ...func() T) T {
	ch := make(chan T, len(functions))
	for _, f := range functions {
		go func(f func() T) {
			ch <- f()
		}(f)
	}
	return <-ch
}

func FirstCtx[T any](ctx context.Context, functions ...func(ctx context.Context) T) (T, error) {
	ch := make(chan T, len(functions))
	for _, f := range functions {
		go func(f func(ctx context.Context) T) {
			ch <- f(ctx)
		}(f)
	}

	select {
	case <-ctx.Done():
		var zero T
		return zero, ctx.Err()
	case res := <-ch:
		return res, nil
	}
}
