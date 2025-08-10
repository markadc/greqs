package log

import "testing"

func TestLog(t *testing.T) {
	name := "Greqs"
	age := 22

	Info("Hello World")
	Debug("hello %v", name)
	Info("type is %T", name)
	Warning("values is %q", name)
	Error("in %p", &name)
	Success("Name is %s Age is %d", name, age)
}
