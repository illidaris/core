package core

import "context"

type ICtxLogger interface {
	DebugCtxf(ctx context.Context, format string, args ...interface{})
	InfoCtxf(ctx context.Context, format string, args ...interface{})
	WarnCtxf(ctx context.Context, format string, args ...interface{})
	ErrorCtxf(ctx context.Context, format string, args ...interface{})
}
