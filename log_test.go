package logger

import "testing"

func Test_WithValues(t *testing.T) {

	log := New(NewOptions())
	log.WithName("test module")
	log.Debug("debug msg")
	log.Info("info msg")
	log.Warn("warn msg")
	log.Error("error msg")
	log.Fatal("fatal msg")
}
