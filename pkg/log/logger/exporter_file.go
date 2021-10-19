package logger

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
	"log"
	"path"
	"time"
)

type FileExporter struct {
	Core *File
}

// Encoder
/**
 * @Description:
 * @receiver e
 * @return zapcore.Encoder
 */
func (e *FileExporter) Encoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        Datetime.ToString(),
		LevelKey:       LevelKey.ToString(),
		NameKey:        "logger",
		CallerKey:      Caller.ToString(),
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     Message.ToString(),
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	} // zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.UTC().Format("2006-01-02T15:04:05.000Z"))
	}
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	var encoder zapcore.Encoder
	switch e.Core.Format {
	case "console":
		encoder = zapcore.NewConsoleEncoder(encoderConfig) // 普通模式
	case "json":
		encoder = zapcore.NewJSONEncoder(encoderConfig) // json格式
	default:
		encoder = zapcore.NewConsoleEncoder(encoderConfig) // 普通模式
	}
	return encoder
}

// Writer
/**
 * @Description:
 * @receiver e
 * @return zapcore.WriteSyncer
 */
func (e *FileExporter) Writer() zapcore.WriteSyncer {
	filename := path.Join(e.Core.FileDirectory, e.Core.FileName)

	hook, err := rotatelogs.New(
		filename+".%Y%m%d%H", // 没有使用go风格反人类的format格式
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*time.Duration(e.Core.MaxDays)), // 保存30天
		rotatelogs.WithRotationTime(time.Hour*24),                         //切割频率 24小时
	)
	if err != nil {
		log.Println("error init writer")
		panic(err)
	}
	return zapcore.AddSync(hook)
}

// Level
/**
 * @Description:
 * @receiver e
 * @return zapcore.Level
 */
func (e *FileExporter) Level() zapcore.Level {
	return e.Core.GetLevel().zapLevel()
}
