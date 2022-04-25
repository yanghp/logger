package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Options_Validate(t *testing.T) {
	opts := &Options{
		Level:            "debug",
		Format:           "json",
		EnableColor:      true,
		DisableCaller:    false,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		Development:      true,
		Name:             "test",
	}

	errs := opts.Validate()
	assert.Nil(t, errs)
}
