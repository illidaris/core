package core

import "context"

type MetaData string

const (
	TraceID      MetaData = "traceId"      // distributed trace, link system
	SessionID    MetaData = "sessionId"    // app trace, link function
	SessionBirth MetaData = "sessionBirth" // session birth

	Action MetaData = "action"
	Step   MetaData = "step"

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
