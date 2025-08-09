package log

import (
	"fmt"
	"time"
)

// CurrTime 返回当前时间
func CurrTime() string {
	currTime := time.Now().Format("2006-01-02 15:04:05")
	return currTime
}

// Default 补充开头
func Default(s, level string, values ...interface{}) string {
	msg := fmt.Sprintf("%s | %-10s | - %s", CurrTime(), level, PyFormat(s, values...))
	return msg
}

func Debug(s string, values ...interface{}) {
	Print(Default(s, "DEBUG", values...), "blue")
}

func Info(s string, values ...interface{}) {
	Print(Default(s, "INFO", values...), "")
}

func Warning(s string, values ...interface{}) {
	Print(Default(s, "WARNING", values...), "yellow")
}

func Error(s string, values ...interface{}) {
	Print(Default(s, "ERROR", values...), "red")
}

func Success(s string, values ...interface{}) {
	Print(Default(s, "SUCCESS", values...), "green")
}
