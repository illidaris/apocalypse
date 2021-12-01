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
	encoderConfig := configEncoder()
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
		filename+".%Y%m%d%H", // {filename}.{%Y%m%d%H}
		rotatelogs.WithMaxAge(time.Hour*24*time.Duration(config.MaxDays)), // days
		rotatelogs.WithRotationTime(time.Hour*24),                         // split
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
