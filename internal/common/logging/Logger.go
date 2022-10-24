package logging

import (
	"log"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	Logger      *zap.Logger
	LoggerSugar *zap.SugaredLogger
}

var loggerLock = &sync.Mutex{}
var loggerInstance *Logger
var once sync.Once

func NewLogger() *Logger {

	once.Do(func() {
		loggerConfig := zap.NewProductionConfig()

		pe := zap.NewProductionEncoderConfig()
		pe.EncodeTime = zapcore.ISO8601TimeEncoder

		loggerConfig.OutputPaths = []string{"stdout"}
		loggerConfig.Encoding = "console"
		loggerSimple, _ := loggerConfig.Build()

		loggerSimple = loggerSimple.WithOptions(zap.AddCallerSkip(1))

		loggerSugar := loggerSimple.Sugar()

		loggerInstance = &Logger{
			Logger:      loggerSimple,
			LoggerSugar: loggerSugar,
		}
	})

	return loggerInstance
}

func (m *Logger) SetupJsonLogger(d bool, f *os.File) *zap.SugaredLogger {

	pe := zap.NewProductionEncoderConfig()
	fileEncoder := zapcore.NewJSONEncoder(pe)
	pe.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(pe)

	level := zap.InfoLevel
	if d {
		level = zap.DebugLevel
	}

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(f), level),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level),
	)

	l := zap.New(core)

	return l.Sugar()
}

func (m *Logger) Debugf(template string, args ...interface{}) {
	if m.Logger == nil {
		log.Fatal("logger in debug is nil")
	}

	m.LoggerSugar.Debugf(template, args...)
}

func (m *Logger) Infof(template string, args ...interface{}) {
	if m.Logger == nil {
		log.Fatal("logger in info is nil")
	}

	m.LoggerSugar.Infof(template, args...)
}

func (m *Logger) Errorf(template string, args ...interface{}) {
	if m.Logger == nil {
		log.Fatal("logger in error is nil")
	}

	m.LoggerSugar.Errorf(template, args...)
}
