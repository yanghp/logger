package main

import (
	"flag"
	"logger"
)

var (
	h      bool
	level  int
	format string
)

func main() {
	flag.BoolVar(&h, "h", false, "Print this help.")
	flag.IntVar(&level, "level", 0, "Log level")
	flag.StringVar(&format, "format", "console", "log output format .")
	flag.Parse()

	if h {
		flag.Usage()
		return
	}

	opts := &logger.Options{
		Level:            "info",
		Format:           format,
		EnableColor:      true,
		OutputPaths:      []string{"test.log", "stdout"},
		ErrorOutputPaths: []string{"error.log"},
	}

	l := logger.New(opts)
	defer l.Flush()

	l.Debug("This is a debug message")
	l.Info("This is a info message")
	l.Warn("This is a warn message")
	l.Error("This is a error message")

	l.Debugf("This is a %s message", "debug")
	l.Infof("This is a %s message", "info")
	l.Warnf("This is a %s message", "warn")
	l.Errorf("This is a %s message", "error")

	l.Debug("This is a debug message", logger.String("key", "value"))
	l.Info("This is a info message", logger.Int32("key2", 10))
	l.Warn("This is a warn message", logger.Bool("key3", false))
	l.Error("This is a error message", logger.Any("key4", "any"))

}
