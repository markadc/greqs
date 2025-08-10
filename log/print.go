package log

import (
	"fmt"
	"strings"
)

// PyFormat 大括号 {} 表示占位
func PyFormat(template string, values ...interface{}) string {
	for _, value := range values {
		template = strings.Replace(template, "{}", fmt.Sprintf("%v", value), 1)
	}
	return template
}

var colors = map[string]string{
	"black":         "30",
	"red":           "31",
	"green":         "32",
	"yellow":        "33",
	"blue":          "34",
	"magenta":       "35",
	"cyan":          "36",
	"white":         "37",
	"gray":          "90",
	"light_red":     "91",
	"light_green":   "92",
	"light_yellow":  "93",
	"light_blue":    "94",
	"light_magenta": "95",
	"light_cyan":    "96",
	"light_white":   "97",
}

// Print 带颜色的打印
func Print(content any, color string) {
	colorCode, exists := colors[color]
	if !exists {
		fmt.Println(content)
	} else {
		fmt.Printf("\033[%sm%v\033[0m\n", colorCode, content)
	}
}

// Printf 带颜色的打印（格式化）
func Printf(color string, format string, values ...any) {
	colorCode, exists := colors[color]
	if !exists {
		fmt.Printf(format, values...)
	} else {
		formatted := fmt.Sprintf(format, values...)
		fmt.Printf("\033[%sm%s\033[0m", colorCode, formatted)
	}
}

// 打印者
type Printer struct{}

func NewPrinter() *Printer {
	return &Printer{}
}

func (p *Printer) Red(a any) {
	Print(a, "red")
}

func (p *Printer) Green(a any) {
	Print(a, "green")
}

func (p *Printer) Yellow(a any) {
	Print(a, "yellow")
}

func (p *Printer) Blue(a any) {
	Print(a, "blue")
}
