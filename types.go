package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Field is an alias for the field structure in the underlying log frame.
type Field = zap.Field

const (
	KeyRequestID string = "requestID"
	KeyUserName  string = "username"
)

type Level = zapcore.Level

var (
	// debug 一般用户测试环境
	DebugLevel = zapcore.DebugLevel
	// info 一般的信息打印,默认日志打印级别
	InfoLevel = zapcore.InfoLevel
	// wran 告警信息，一般不需人为处理
	WarnLevel = zapcore.WarnLevel
	// 正常不会出现error日志，如果出现，证明有问题，需要观察并进行处理
	ErrorLevel = zapcore.ErrorLevel
	// 当go项目出现panic错误时，打印此日志，一般只有go项目才有
	PanicLevel = zapcore.PanicLevel
	// fatal日志，当项目重大问题时打印 fatal日志，并执行 os.Exit(1)
	FatalLevel = zapcore.FatalLevel
)

// 引用zap里的别名
var (
	Any         = zap.Any
	Array       = zap.Array
	Object      = zap.Object
	Binary      = zap.Binary
	Bool        = zap.Bool
	Bools       = zap.Bools
	ByteString  = zap.ByteString
	ByteStrings = zap.ByteStrings
	Complex64   = zap.Complex64
	Complex64s  = zap.Complex64s
	Complex128  = zap.Complex128
	Complex128s = zap.Complex128s
	Duration    = zap.Duration
	Durations   = zap.Durations
	Err         = zap.Error
	Errors      = zap.Errors
	Float32     = zap.Float32
	Float32s    = zap.Float32s
	Float64     = zap.Float64
	Float64s    = zap.Float64s
	Int         = zap.Int
	Ints        = zap.Ints
	Int8        = zap.Int8
	Int8s       = zap.Int8s
	Int16       = zap.Int16
	Int16s      = zap.Int16s
	Int32       = zap.Int32
	Int32s      = zap.Int32s
	Int64       = zap.Int64
	Int64s      = zap.Int64s
	Namespace   = zap.Namespace
	Reflect     = zap.Reflect
	Stack       = zap.Stack
	String      = zap.String
	Stringer    = zap.Stringer
	Strings     = zap.Strings
	Time        = zap.Time
	Times       = zap.Times
	Uint        = zap.Uint
	Uints       = zap.Uints
	Uint8       = zap.Uint8
	Uint8s      = zap.Uint8s
	Uint16      = zap.Uint16
	Uint16s     = zap.Uint16s
	Uint32      = zap.Uint32
	Uint32s     = zap.Uint32s
	Uint64      = zap.Uint64
	Uint64s     = zap.Uint64s
	Uintptr     = zap.Uintptr
	Uintptrs    = zap.Uintptrs
)
