package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log logger
)

type fdLogger interface{
	Printf(format string, v ...interface{})
//	Info(msg string, tags ...zap.Field)
//	Error(msg string, err error, tags ...zap.Field)
}

type logger struct {
	log *zap.Logger
}

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "msg",
			LevelKey:       "level",
			TimeKey:        "time",
			NameKey:        "",
			CallerKey:      "",
			StacktraceKey:  "",
			LineEnding:     "",
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: nil,
			EncodeCaller:   zapcore.ShortCallerEncoder,
			EncodeName:     nil,
		}}

	var err error
	if log.log, err = logConfig.Build(); err != nil {
		panic(err)
	}

}

func GetLogger() fdLogger {
	return log
}

func (l logger) Printf(format string, v ...interface{}) {
	if len(v) == 0 {
		Info(format)
	} else {
		Info(fmt.Sprintf(format, v))
	}
}

func Info(msg string, tags ...zap.Field) {
	log.log.Info(msg, tags...)
	log.log.Sync()
}

func Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.log.Error(msg, tags...)
	log.log.Sync()
}
