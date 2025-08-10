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

// MakeLog 制作日志
func MakeLog(level, s string, values ...any) string {
	//msg := fmt.Sprintf("%s | %-10s | - %s", CurrTime(), level, PyFormat(s, values...))
	msg := fmt.Sprintf("%s | %-10s | - %s", CurrTime(), level, fmt.Sprintf(s, values...))
	return msg
}

// Debug 调试日志
func Debug(s string, values ...any) {
	Print(MakeLog("INFO", s, values...), "blue")
}

// Info 一般日志
func Info(s string, values ...any) {
	Print(MakeLog("INFO", s, values...), "")
}

// Warning 警告日志
func Warning(s string, values ...any) {
	Print(MakeLog("WARNING", s, values...), "yellow")
}

// Error 错误日志
func Error(s string, values ...any) {
	Print(MakeLog("ERROR", s, values...), "red")
}

// Success 成功日志
func Success(s string, values ...any) {
	Print(MakeLog("SUCCESS", s, values...), "green")
}

// Red 红色打印
func Red(s string, values ...any) {
	Printf("red", s+"\n", values...)
}

// Yellow 黄色打印
func Yellow(s string, values ...any) {
	Printf("yellow", s+"\n", values...)
}

// Blue 蓝色打印
func Blue(s string, values ...any) {
	Printf("blue", s+"\n", values...)
}

// Green 绿色打印
func Green(s string, values ...any) {
	Printf("green", s+"\n", values...)
}
