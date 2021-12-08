package constant

import "testing"

const  (
	Monday = iota + 1
	Tuesday
	wednesday
	)

const (
	Readable = 1 << iota // 0001
	Writable
	Executable
)

func TestConstantTry(t *testing.T)  {
	t.Log(Monday, Tuesday)

}

func TestConstantTyy2(t *testing.T)  {
	a := 7 // 0111
	t.Log(a&Readable ==  Readable, a&Writable == Writable, a&Executable == Executable)
}