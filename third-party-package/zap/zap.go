package zap

import (
	"errors"
	"fmt"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
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

		encoderConfig := zapcore.EncoderConfig{
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
		}

		highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.ErrorLevel
		})
		lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl < zapcore.ErrorLevel
		})

		consoleDebugging := zapcore.Lock(os.Stdout)
		consoleErrors := zapcore.Lock(os.Stderr)

		consoleEncoder := zapcore.NewJSONEncoder(encoderConfig)

		core := zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
			zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
		)

		singletonLogger = &Logger{
			logger: zap.New(core),
		}
	})
}

func Infof(traceID string, format string, v ...interface{}) {
	singletonLogger.logger.Info(fmt.Sprintf(format, v...), zap.String("trace_id", traceID))
}

func Errorf(traceID string, format string, v ...interface{}) {
	singletonLogger.logger.Error(fmt.Sprintf(format, v...), zap.String("trace_id", traceID))
}

func Fatalf(traceID string, format string, v ...interface{}) {
	singletonLogger.logger.Fatal(fmt.Sprintf(format, v...), zap.String("trace_id", traceID))
}
