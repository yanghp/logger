package logger

import "testing"

func Test_WithValues(t *testing.T) {

	l := std.WithValues("k", "v")
	l.Info("foo")
	l.Info("bar")
	l.Flush()
}

func Test_WithName(t *testing.T) {
	l := std.WithName("test")
	l.Infow("hello world", "foo", "bar")
}
