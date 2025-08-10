package log

import "testing"

func TestPyFormat(t *testing.T) {
	s1 := PyFormat("当前时间是：{}", CurrTime())
	Print(s1, "blue")

	name, code := "王天风", "毒蜂"
	s2 := PyFormat("姓名：{}，代号：{}。", name, code)
	Print(s2, "red")
}

func TestPrintf(t *testing.T) {
	Printf("green", "Hello Greqs\n")

	name, age := "Wauo", 22
	Printf("red", "My name is %s and age is %d\n", name, age)
	Printf("yellow", "My name is %s and age is %d\n", name, age)
	Printf("blue", "My name is %s and age is %d\n", name, age)
}
