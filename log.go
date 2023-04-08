package core

import "context"

type ICtxLogger interface {
	DebugCtxf(ctx context.Context, format string, args ...interface{})
	InfoCtxf(ctx context.Context, format string, args ...interface{})
	WarnCtxf(ctx context.Context, format string, args ...interface{})
	ErrorCtxf(ctx context.Context, format string, args ...interface{})
}

type ILogger interface {
	DebugCtx(ctx context.Context, msg string, kv map[string]string)
	InfoCtx(ctx context.Context, msg string, kv map[string]string)
	WarnCtx(ctx context.Context, msg string, kv map[string]string)
	ErrorCtx(ctx context.Context, msg string, kv map[string]string)
}
