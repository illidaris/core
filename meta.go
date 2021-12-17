package core

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
