package logger

type LoggerKey string

const (
	Datetime LoggerKey = "datetime"
	Caller   LoggerKey = "caller"
	Duration LoggerKey = "duration"
	LevelKey LoggerKey = "level"
	Message  LoggerKey = "message"
)

// ToString
/**
 * @Description:
 * @receiver key
 * @return string
 */
func (key LoggerKey) ToString() string {
	return string(key)
}
