package logger

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	DebugLogger
	InfoLogger
	WarnLogger
	ErrorLogger
	FatalLogger
	V(level int) InfoLogger
	Write(p []byte)(n int,err error)
	WithValues(kv interface{})Logger
	WithName(name string)Logger
	WithContext(ctx context.Context)context.Context
	Flush()
}

// info
type InfoLogger interface {
	Info(msg string, fields ...Field)
	Infof(format string, v ...interface{})
	Infow(msg string, kv ...interface{})
	Enabled()bool
}

// debug
type DebugLogger interface {
	Debug(msg string, fields ...Field)
	Debugf(format string,v ...interface{})
	Debugw(msg string,kv ...interface{})
}

// warn
type WarnLogger interface {
	Warn(msg string ,fields ...Field)
	Warnf(format string,v ...interface{})
	Warnw(msg string,kv ...interface{})
}

// error
type ErrorLogger interface {
	Error(msg string,fields ...Field)
	Errorf(format string,v ...interface{})
	Errorw(msg string,kv ...interface{})
}

//panic
type PanicLogger interface {
	Panic(msg string,fields ...Field)
	Panicf(format string,v ...interface{})
	Panicw(msg string,kv ...interface{})
}

//Fatal
type FatalLogger interface {
	Fatal(mgs string,fields ...Field)
	Fatalf(format string, v ...interface{})
	Fatalw(msg string, kv ...interface{})
}

// zapLogger is log.Logger that uses zap to log
type zapLogger struct {
	zapLogger *zap.Logger
	InfoLogger
}

type infoLogger struct {
	level zapcore.Level
	log *zap.Logger
}

func (l *infoLogger) Enabled()bool{
	return true
}
func(l *infoLogger) Info(msg string, fields ...Field){
	if check := l.log.Check(l.level, msg); check !=nil{
		check.Write(fields...)
	}
}
func (l *infoLogger)Infow(msg string,keysAndValues ...interface{}){
	if check := l.log.Check(l.level, msg);check !=nil{
		check.Write(handleFields(l.log,keysAndValues)...)
	}
}

func handleFields(l *zap.Logger,args []interface{},additional ...zap.Field)[]zap.Field{
	if len(args) ==0{
		return additional
	}
	fields:= make([]zap.Field,0,len(args)/2+ len(additional))
	for i:=0; i< len(args);{
		if _,ok:= args[i].(zap.Field);ok{
			l.Debug("strongly-typed Zap Field passed to logr", zap.Any("zap field", args[i]))
			break
		}

		if i == len(args) -1{
			l.DPanic("odd number of arguments passed as key-value pairs for logging", zap.Any("ignored key", args[i]))
			break
		}
		key,val:= args[i],args[i+1]
		keyStr,isString:= key.(string)
		if !isString{
			l.DPanic("non-string key argument passed to logging, ignoring all later arguments",
				zap.Any("invalid key", key))
			break
		}
		fields = append(fields,zap.Any(keyStr,val))
		i+=2
	}
	return append(fields,additional...)
}

// todo yanghp