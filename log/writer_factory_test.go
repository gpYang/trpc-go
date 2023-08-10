package log_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"trpc.group/trpc-go/trpc-go/log"
)

func TestWriterFactory(t *testing.T) {
	f1 := &log.ConsoleWriterFactory{}
	assert.Equal(t, "log", f1.Type())

	// empty decoder
	err := f1.Setup("default", nil)
	assert.NotNil(t, err)

	f2 := &log.FileWriterFactory{}
	assert.Equal(t, "log", f2.Type())
	// empty decoder
	err = f2.Setup("default", nil)
	assert.NotNil(t, err)

	f3 := &log.ConsoleWriterFactory{}
	assert.Equal(t, "log", f3.Type())
	err = f3.Setup("default", &fakeDecoder{})
	assert.NotNil(t, err)

	f4 := &log.FileWriterFactory{}
	assert.Equal(t, "log", f4.Type())
	err = f4.Setup("default", &fakeDecoder{})
	assert.NotNil(t, err)
}

func TestFileWriterFactory_Setup(t *testing.T) {
	var fileCfg = []log.OutputConfig{
		{
			Writer: "file",
			WriteConfig: log.WriteConfig{
				Filename:   "trpc_time.log",
				MaxAge:     7,
				MaxBackups: 10,
				MaxSize:    100,
				TimeUnit:   log.Day,
				LogPath:    "log",
			},
		},
	}
	logger := log.NewZapLog(fileCfg)
	assert.NotNil(t, logger)
}