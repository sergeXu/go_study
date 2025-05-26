package test

import "fmt"

const (
	BEIJING = iota * 2
	SHANGHAI
	HANGZHOU
)

const (
	//iota 只能在const中使用
	a, b = iota + 1, iota + 2
	c, d
	e, f
	g, h = iota * 2, iota * 3
	i, k
)

func TestConst() {
	fmt.Println("TestConst Run")
	const length int = 122
	fmt.Println("length =", length)

	fmt.Println("BEIJING=", BEIJING)
	fmt.Println("SHANGHAI=", SHANGHAI)
	fmt.Println("HANGZHOU=", HANGZHOU)

}
