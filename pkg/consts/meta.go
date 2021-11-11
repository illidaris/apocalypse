package consts

type MetaData string

const (
	TraceID      MetaData = "traceId"
	SessionID    MetaData = "sessionId"
	SessionBirth MetaData = "sessionBirth"

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
