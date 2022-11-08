package logging

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/maker-space-experimenta/printer-kiosk/internal/common/configuration"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	Logger      *zap.Logger
	LoggerSugar *zap.SugaredLogger
}

var lock = &sync.Mutex{}
var loggerInstance *Logger

func NewCustomLogger(config *configuration.Config) *Logger {

	var level zapcore.Level

	switch config.Logging.Level {
	case "fatal":
		level = zapcore.FatalLevel
	case "panic":
		level = zapcore.PanicLevel
	case "dpanic":
		level = zapcore.DPanicLevel
	case "error":
		level = zapcore.ErrorLevel
	case "warn":
		level = zapcore.WarnLevel
	case "info":
		level = zapcore.InfoLevel
	default:
		level = zapcore.DebugLevel
	}

	log.Printf("create new zap config, level %v", level)

	cfg := zap.Config{
		Encoding:    "console",
		Level:       zap.NewAtomicLevelAt(level),
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			EncodeTime:  zapcore.ISO8601TimeEncoder,
			MessageKey:  "message",
			TimeKey:     "timestamp",
			CallerKey:   "caller",
			LevelKey:    "level",
			EncodeLevel: zapcore.LowercaseColorLevelEncoder, // "lowercase",
		},
	}

	log.Printf("config created")

	loggerSimple, _ := cfg.Build(zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel), zap.AddCallerSkip(1))
	loggerSugar := loggerSimple.Sugar()

	if loggerSimple == nil {
		log.Printf("loggerSimple is nil")
	}
	if loggerSugar == nil {
		log.Printf("loggerSugar is nil")
	}

	return &Logger{
		Logger:      loggerSimple,
		LoggerSugar: loggerSugar,
	}
}

func NewProductionLogger() *Logger {
	log.Printf("create ProductionLogger")
	log.Printf("create logger config")
	loggerConfig := zap.NewProductionConfig()

	loggerConfig.OutputPaths = []string{"stdout"}
	loggerConfig.Encoding = "console"

	log.Printf("build logger from config")
	loggerSimple, err := loggerConfig.Build()

	if err != nil {
		log.Printf("cannot load config: %v", err)
		os.Exit(1)
	}

	loggerSimple = loggerSimple.WithOptions(zap.AddCallerSkip(1))
	loggerSugar := loggerSimple.Sugar()

	if loggerSimple == nil {
		log.Printf("loggerSimple is nil")
	}
	if loggerSugar == nil {
		log.Printf("loggerSugar is nil")
	}

	return &Logger{
		Logger:      loggerSimple,
		LoggerSugar: loggerSugar,
	}
}

func NewLogger() *Logger {

	if loggerInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if loggerInstance == nil {
			fmt.Println("Creating Logger instance now.")

			configService := configuration.NewConfigService()
			config, err := configService.GetConfig()
			if err != nil {

			}

			loggerInstance = NewCustomLogger(config)
			// loggerInstance = NewProductionLogger()
		} else {
			loggerInstance.Debugf("Logger instance already created.")
		}
	} else {
		loggerInstance.Debugf("Logger instance already created.")
	}

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
		log.Printf("FATAL: logger in error is nil")
		return
	}
	if m.LoggerSugar == nil {
		log.Fatal("FATAL: loggerSugar in debug is nil")
		return
	}

	m.LoggerSugar.Debugf(template, args...)
}

func (m *Logger) Infof(template string, args ...interface{}) {
	if m.Logger == nil {
		log.Printf("FATAL: logger in error is nil")
		return
	}
	if m.LoggerSugar == nil {
		log.Fatal("FATAL: loggerSugar in debug is nil")
		return
	}

	m.LoggerSugar.Infof(template, args...)
}

func (m *Logger) Errorf(template string, args ...interface{}) {
	if m.Logger == nil {
		log.Printf("FATAL: logger in error is nil")
		return
	}
	if m.LoggerSugar == nil {
		log.Fatal("FATAL: loggerSugar in debug is nil")
		return
	}

	m.LoggerSugar.Errorf(template, args...)
}

func (m *Logger) Fatalf(template string, args ...interface{}) {
	if m.Logger == nil {
		log.Printf("FATAL: logger in error is nil")
		return
	}
	if m.LoggerSugar == nil {
		log.Fatal("FATAL: loggerSugar in debug is nil")
		return
	}

	m.LoggerSugar.Fatalf(template, args...)
}
