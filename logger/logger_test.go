package logger

import (
	"github.com/Kephas73/lib-kephas/env"
	"testing"
	"time"
)

func init() {
	env.SetupConfigEnv("../config.json")
}

func TestNewLogger(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			panic(err)
		}
	}()

	NewLogger(env.Environment.Logger.Path, env.Environment.Logger.Prefix, env.Environment.Debug)
	Info("Test:Test: %v", time.Now())
	Error("Test:Test: %v", time.Now())
}
