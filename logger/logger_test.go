package logger

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestGetLogLevel(t *testing.T) {
	config := []*LogModuleConfig{
		{
			ModuleName:   "debug",
			LogLevel:     DEBUG,
			FilePath:     "logs/debug.log",
			MaxAge:       365,
			RotationTime: 1,
			LogInConsole: false,
			ShowColor:    true,
		},
		{
			ModuleName:   "info",
			LogLevel:     INFO,
			FilePath:     "logs/info.log",
			MaxAge:       365,
			RotationTime: 1,
			LogInConsole: false,
			ShowColor:    true,
		},
		{
			ModuleName:   "warn",
			LogLevel:     WARN,
			FilePath:     "logs/warn.log",
			MaxAge:       365,
			RotationTime: 1,
			LogInConsole: false,
			ShowColor:    true,
		},
		{
			ModuleName:   "error",
			LogLevel:     ERROR,
			FilePath:     "logs/error.log",
			MaxAge:       365,
			RotationTime: 1,
			LogInConsole: false,
			ShowColor:    true,
		},
	}
	InitLogConfig(config)

	defaultLogConfig = &Config{}
	InitLogConfig(config)

}

func TestModuleNameNull(t *testing.T) {
	config := []*LogModuleConfig{
		{
			ModuleName:   "",
			LogLevel:     "null",
			FilePath:     "logs/null.log",
			MaxAge:       365,
			RotationTime: 1,
			LogInConsole: false,
			ShowColor:    true,
		},
	}
	InitLogConfig(config)
}

func TestLogger(t *testing.T) {
	// test init module log
	config := []*LogModuleConfig{
		{
			ModuleName:   "server",
			LogLevel:     INFO,
			FilePath:     "logs/server.log",
			MaxAge:       365,
			RotationTime: 1,
			LogInConsole: false,
			ShowColor:    true,
		},
		{
			ModuleName:   "default",
			LogLevel:     INFO,
			FilePath:     "logs/default.log",
			MaxAge:       365,
			RotationTime: 1,
			LogInConsole: false,
			ShowColor:    true,
		},
	}
	InitLogConfig(config)

	// test get logger
	var log *zap.SugaredLogger
	log = GetLogger(ModuleDefault)
	require.NotNil(t, log)
	log = GetLogger(ModuleServer)
	require.NotNil(t, log)
	log = GetLogger(ModuleCli)
	require.NotNil(t, log)
	log = GetLogger(ModuleHandler)
	require.NotNil(t, log)
	log = GetLogger(ModuleAdapter)
	require.NotNil(t, log)
}

func TestCustomLevelEncoder(t *testing.T) {
	je := zapcore.NewJSONEncoder(zapcore.EncoderConfig{})
	encoder, ok := je.(zapcore.PrimitiveArrayEncoder)
	require.Equal(t, ok, true)
	CustomLevelEncoder(zapcore.DebugLevel, encoder)
}

func TestCustomTimeEncoder(t *testing.T) {
	je := zapcore.NewJSONEncoder(zapcore.EncoderConfig{})
	encoder, ok := je.(zapcore.PrimitiveArrayEncoder)
	require.Equal(t, ok, true)
	CustomTimeEncoder(time.Now(), encoder)
}
