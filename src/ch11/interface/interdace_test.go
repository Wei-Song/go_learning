package _interface

import "testing"

type Programmer interface {
	WrightHelloWorld() string
}

type GoProgrammer struct{

}

func (g *GoProgrammer) WrightHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WrightHelloWorld())
}