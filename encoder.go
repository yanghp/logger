package logger

import (
	"go.uber.org/zap/zapcore"
	"time"
)

func timeEncoder(t time.Time,enc zapcore.PrimitiveArrayEncoder){
	enc.AppendString(t.Format("2016-01-02 15:04:05.000"))
}

func milliSecondDurationEncoder(d time.Duration,enc zapcore.PrimitiveArrayEncoder){
	enc.AppendFloat64(float64(d)/float64(time.Millisecond))
}