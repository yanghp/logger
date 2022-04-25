package logger

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"logger/klog"
)

type Logger interface {
	DebugLogger
	InfoLogger
	WarnLogger
	ErrorLogger
	FatalLogger
	V(level int) InfoLogger
	Write(p []byte) (n int, err error)
	WithValues(kv ...interface{}) Logger
	WithName(name string) Logger
	WithContext(ctx context.Context) context.Context
	Flush()
}

// InfoLogger info
type InfoLogger interface {
	Info(msg string, fields ...Field)
	Infof(format string, v ...interface{})
	Infow(msg string, kv ...interface{})
	Enabled() bool
}

// DebugLogger debug
type DebugLogger interface {
	Debug(msg string, fields ...Field)
	Debugf(format string, v ...interface{})
	Debugw(msg string, kv ...interface{})
}

// WarnLogger warn
type WarnLogger interface {
	Warn(msg string, fields ...Field)
	Warnf(format string, v ...interface{})
	Warnw(msg string, kv ...interface{})
}

// ErrorLogger error
type ErrorLogger interface {
	Error(msg string, fields ...Field)
	Errorf(format string, v ...interface{})
	Errorw(msg string, kv ...interface{})
}

// PanicLogger panic
type PanicLogger interface {
	Panic(msg string, fields ...Field)
	Panicf(format string, v ...interface{})
	Panicw(msg string, kv ...interface{})
}

// FatalLogger fatal
type FatalLogger interface {
	Fatal(mgs string, fields ...Field)
	Fatalf(format string, v ...interface{})
	Fatalw(msg string, kv ...interface{})
}

var _ Logger = &zapLogger{}

// zapLogger is log.Logger that uses zap to log
type zapLogger struct {
	zapLogger *zap.Logger
	infoLogger
}

//noopInfoLogger is a logger.InfoLogger that's always disabled and does nothing.
type noopInfoLogger struct {
}

var disabledInfoLogger = &noopInfoLogger{}

func (n *noopInfoLogger) Info(msg string, fields ...Field) {}

func (n *noopInfoLogger) Infof(format string, v ...interface{}) {}

func (n *noopInfoLogger) Infow(msg string, kv ...interface{}) {}

func (n *noopInfoLogger) Enabled() bool {
	return false
}

func (l *zapLogger) Debug(msg string, fields ...Field) {
	l.zapLogger.Debug(msg, fields...)
}

func (l *zapLogger) Debugf(format string, v ...interface{}) {
	l.zapLogger.Sugar().Debugf(format, v...)
}

func (l *zapLogger) Debugw(msg string, kv ...interface{}) {
	l.zapLogger.Sugar().Debugw(msg, kv...)
}

func (l *zapLogger) Warn(msg string, fields ...Field) {
	l.zapLogger.Warn(msg, fields...)
}

func (l *zapLogger) Warnf(format string, v ...interface{}) {
	l.zapLogger.Sugar().Warnf(format, v...)
}

func (l *zapLogger) Warnw(msg string, kv ...interface{}) {
	l.zapLogger.Sugar().Warnw(msg, kv...)
}

func (l *zapLogger) Error(msg string, fields ...Field) {
	l.zapLogger.Error(msg, fields...)
}

func (l *zapLogger) Errorf(format string, v ...interface{}) {
	l.zapLogger.Sugar().Errorf(format, v...)
}

func (l *zapLogger) Errorw(msg string, kv ...interface{}) {
	l.zapLogger.Sugar().Errorw(msg, kv...)
}

func (l *zapLogger) Fatal(msg string, fields ...Field) {
	l.zapLogger.Fatal(msg, fields...)
}

func (l *zapLogger) Fatalf(format string, v ...interface{}) {
	l.zapLogger.Sugar().Fatalf(format, v...)
}

func (l *zapLogger) Fatalw(msg string, kv ...interface{}) {
	l.zapLogger.Sugar().Fatalw(msg, kv...)
}

func (l *zapLogger) V(level int) InfoLogger {
	if level < 0 || level > 1 {
		panic("log level error: valid log level is [0, 1]")
	}
	lvl := zapcore.Level(-1 * level)
	if l.zapLogger.Core().Enabled(lvl) {
		return &infoLogger{
			level: lvl,
			log:   l.zapLogger,
		}
	}
	return disabledInfoLogger
}

func (l *zapLogger) Write(p []byte) (n int, err error) {
	l.zapLogger.Info(string(p))
	return len(p), nil
}

func (l *zapLogger) WithValues(kv ...interface{}) Logger {
	newLogger := l.zapLogger.With(handleFields(l.zapLogger, kv)...)
	return NewLogger(newLogger)
}

// NewLogger creates a new logr.Logger using the given Zap Logger to log.
func NewLogger(l *zap.Logger) Logger {
	return &zapLogger{
		zapLogger: l,
		infoLogger: infoLogger{
			log:   l,
			level: zap.InfoLevel,
		},
	}
}

func (l *zapLogger) WithName(name string) Logger {
	newLogger := l.zapLogger.Named(name)
	return NewLogger(newLogger)
}

func (l *zapLogger) Flush() {
	_ = l.zapLogger.Sync()
}

type infoLogger struct {
	level zapcore.Level
	log   *zap.Logger
}

func New(opts *Options) *zapLogger {
	if opts == nil {
		opts = NewOptions()
	}
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(opts.Level)); err != nil {
		zapLevel = zapcore.InfoLevel
	}
	encodeLevel := zapcore.CapitalLevelEncoder
	if opts.Format == consoleFormat && opts.EnableColor {
		encodeLevel = zapcore.CapitalColorLevelEncoder
	}

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "timestamp",
		NameKey:        "logger",
		CallerKey:      "call",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    encodeLevel,
		EncodeTime:     timeEncoder,
		EncodeDuration: milliSecondDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	loggerConfig := &zap.Config{
		Level:             zap.NewAtomicLevelAt(zapLevel),
		Development:       opts.Development,
		DisableCaller:     opts.DisableCaller,
		DisableStacktrace: opts.DisableStacktrace,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         opts.Format,
		EncoderConfig:    encoderConfig,
		OutputPaths:      opts.OutputPaths,
		ErrorOutputPaths: opts.ErrorOutputPaths,
	}
	var err error
	l, err := loggerConfig.Build(zap.AddStacktrace(zapcore.PanicLevel), zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}
	logger := &zapLogger{
		zapLogger: l.Named(opts.Name),
		infoLogger: infoLogger{
			level: zap.InfoLevel,
			log:   l,
		},
	}
	klog.InitLogger(l)
	zap.RedirectStdLog(l)
	return logger
}

func (l *infoLogger) Enabled() bool {
	return true
}
func (l *infoLogger) Info(msg string, fields ...Field) {
	if check := l.log.Check(l.level, msg); check != nil {
		check.Write(fields...)
	}
}
func (l *infoLogger) Infow(msg string, keysAndValues ...interface{}) {
	if check := l.log.Check(l.level, msg); check != nil {
		check.Write(handleFields(l.log, keysAndValues)...)
	}
}

func (l *infoLogger) Infof(format string, args ...interface{}) {
	if checkedEntry := l.log.Check(l.level, fmt.Sprintf(format, args...)); checkedEntry != nil {
		checkedEntry.Write()
	}
}

func (l *zapLogger) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, logContextKey, l)
}

func handleFields(l *zap.Logger, args []interface{}, additional ...zap.Field) []zap.Field {
	if len(args) == 0 {
		return additional
	}
	fields := make([]zap.Field, 0, len(args)/2+len(additional))
	for i := 0; i < len(args); {
		if _, ok := args[i].(zap.Field); ok {
			l.Debug("strongly-typed Zap Field passed to logr", zap.Any("zap field", args[i]))
			break
		}

		if i == len(args)-1 {
			l.DPanic("odd number of arguments passed as key-value pairs for logging", zap.Any("ignored key", args[i]))
			break
		}
		key, val := args[i], args[i+1]
		keyStr, isString := key.(string)
		if !isString {
			l.DPanic("non-string key argument passed to logging, ignoring all later arguments",
				zap.Any("invalid key", key))
			break
		}
		fields = append(fields, zap.Any(keyStr, val))
		i += 2
	}
	return append(fields, additional...)
}
