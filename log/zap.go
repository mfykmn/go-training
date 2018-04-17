package log

import (
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"fmt"
)

var singletonLogger *Logger

func GetLogger() *Logger{
	return singletonLogger
}

var once sync.Once

func New() {
	once.Do(func() {
		level := zap.NewAtomicLevel()
		level.SetLevel(zapcore.DebugLevel)
		zapConfig := zap.Config{
			Level:    level,
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

type Logger struct {
	logger *zap.Logger
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.logger.Fatal(fmt.Sprintf("Fatal %s", v), zap.String("test","a"))
}


