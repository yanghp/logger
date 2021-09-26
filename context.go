package logger

import (
	"context"
)

type key int

const logContextKey key= iota

func WithContext(ctx context.Context) context.Context{
	return std.WithContext(ctx)
}

