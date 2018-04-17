package log

import (
	"sync"
	"fmt"
	"errors"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Level int8

const (
 	Debug Level = iota
 	Info
 	Warn
 	Error
 	Dpanic
 	Panic
 	Fatal
)

func getZapLevel(level Level) (zapcore.Level, error) {
	switch level {
		case Debug:
			return zapcore.DebugLevel, nil
		case Info:
			return zapcore.InfoLevel, nil
		case Warn:
			return zapcore.WarnLevel, nil
		case Error:
			return zapcore.ErrorLevel, nil
		case Dpanic:
			return zapcore.DPanicLevel, nil
		case Panic:
			return zapcore.PanicLevel, nil
		case Fatal:
			return zapcore.FatalLevel, nil
		default:
			// -2はzapのlevelに存在しない数値
			return -2, errors.New(fmt.Sprintf("Failed match level. config level=%s", level))
		}
}

type Logger struct {
	logger *zap.Logger
}

var singletonLogger *Logger
var once sync.Once

func New(level Level) {
	once.Do(func() {
		zapLevel, _ := getZapLevel(level)

		atomicLevel := zap.NewAtomicLevel()
		atomicLevel.SetLevel(zapLevel)
		zapConfig := zap.Config{
			Level:    atomicLevel,
			Encoding: "json",
			EncoderConfig: zapcore.EncoderConfig{
				TimeKey:        "Time",
				LevelKey:       "Level",
				NameKey:        "Name",
				CallerKey:      "Caller",
				MessageKey:     "Msg",
				StacktraceKey:  "St",
				EncodeLevel:    zapcore.CapitalLevelEncoder,
				EncodeTime:     zapcore.ISO8601TimeEncoder,
				EncodeDuration: zapcore.StringDurationEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
			},
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
		}

		logger, _ := zapConfig.Build()

		singletonLogger = &Logger{
			logger: logger,
		}
	})
}

func Fatalf(traceID string, format string, v ...interface{}) {
	singletonLogger.logger.Fatal(fmt.Sprintf(format, v...), zap.String("trace_id",traceID))
}


