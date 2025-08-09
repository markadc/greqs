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

// Print 带颜色的输出
func Print(content any, color string) {
	colorCode, exists := colors[color]
	if !exists {
		fmt.Println(content)
	} else {
		fmt.Printf("\033[%sm%v\033[0m\n", colorCode, content)
	}
}
