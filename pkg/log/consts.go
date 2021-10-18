package log

type LoggerKey string

const (
	Datetime LoggerKey = "datetime"
	Caller   LoggerKey = "caller"
	Duration LoggerKey = "duration"
	LevelKey LoggerKey = "level"
	Message  LoggerKey = "message"
)

func (key LoggerKey) ToString() string {
	return string(key)
}
