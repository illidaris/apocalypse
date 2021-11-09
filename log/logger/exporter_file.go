package logger

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
	"log"
	"path"
	"time"
)

type FileExporter struct{}

// Encoder
/**
 * @Description:
 * @receiver e
 * @return zapcore.Encoder
 */
func (e *FileExporter) Encoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        Datetime.String(),
		LevelKey:       LevelKey.String(),
		NameKey:        "logger",
		CallerKey:      Caller.String(),
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     Message.String(),
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
	return fmtEncoder(config.Format, encoderConfig)
}

// Writer
/**
 * @Description:
 * @receiver e
 * @return zapcore.WriteSyncer
 */
func (e *FileExporter) Writer() zapcore.WriteSyncer {
	filename := path.Join(config.FileDirectory, config.FileName)

	hook, err := rotatelogs.New(
		filename+".%Y%m%d%H", // 没有使用go风格反人类的format格式
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*time.Duration(config.MaxDays)), // 保存30天
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
	return config.GetLevel().zapLevel()
}
