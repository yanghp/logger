package logger

import (
	"context"
)

type key int

const logContextKey key = iota

func FromContext(ctx context.Context) Logger {
	if ctx != nil {
		log := ctx.Value(logContextKey)
		if log != nil {
			return log.(Logger)
		}
	}
	return nil
}
