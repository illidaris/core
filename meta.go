package core

import "context"

type MetaData string

const (
	TraceID      MetaData = "traceId"      // distributed trace, link system
	ParentID     MetaData = "parentId"     // parent id
	SessionID    MetaData = "sessionId"    // app trace, link function
	SessionBirth MetaData = "sessionBirth" // session birth
	SpanID       MetaData = "spanId"       // span id

	Action MetaData = "action"
	Step   MetaData = "step"
	XKey   MetaData = "x_key" // x_key keyword for business search

	Error    MetaData = "error"
	Duration MetaData = "duration"

	Category      MetaData = "category"
	Datetime      MetaData = "datetime"
	Caller        MetaData = "caller"
	LevelKey      MetaData = "level"
	Message       MetaData = "message"
	NameKey       MetaData = "logger"
	StacktraceKey MetaData = "stacktrace"
	LineEnding    MetaData = "\n"
	FunctionKey   MetaData = ""
)

func (key MetaData) String() string {
	return string(key)
}

func (key MetaData) Get(ctx context.Context) interface{} {
	return ctx.Value(key)
}

func (key MetaData) Set(ctx context.Context, v interface{}) context.Context {
	return context.WithValue(ctx, key, v)
}

func (key MetaData) GetString(ctx context.Context) string {
	v := key.Get(ctx)
	if str, ok := v.(string); ok {
		return str
	}
	return ""
}

func (key MetaData) SetString(ctx context.Context, v string) context.Context {
	return key.Set(ctx, v)
}
