package my_string

import (
	"testing"
	"unsafe"
)

func TestString(t *testing.T) {

	var s string

	//初始化为默认值""
	t.Log(s)

	s = "hello"
	t.Log(len(s))

	//string是不可变的byte slice
	//s[1] = '3'

	//可以存储任何二进制数据
	s = "\xE4\xB8\xA5"
	t.Log(s)

	s = "中"
	//是byte数
	t.Log(len(s))

	c := []rune(s)
	t.Log("rune size:", unsafe.Sizeof(c[0]))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _,c := range s {
		t.Logf("%[1]c %[1]x",c)
	}
}
