package log

import "testing"

func TestLog(t *testing.T) {
	Debug("hello world")
	Info("hello world")
	Warning("hello world")
	Error("hello world")
	Success("hello world")
}
