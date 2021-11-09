package logger

type LoggerKey string

const (
	Datetime LoggerKey = "datetime"
	Caller   LoggerKey = "caller"
	Duration LoggerKey = "duration"
	LevelKey LoggerKey = "level"
	Message  LoggerKey = "message"
)

// String
/**
 * @Description:
 * @receiver key
 * @return string
 */
func (key LoggerKey) String() string {
	return string(key)
}
