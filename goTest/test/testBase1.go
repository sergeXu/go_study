package test

import (
	"fmt"
	"time"
)

func goFunc(i int) {
	fmt.Println("go routine", i, " ...")
}

func TestStart() {
	fmt.Println("TestStart Run")
	//TestGofunc()
	//TestVar()
	//TestConst()
	//testFunc()
	testSlice()
	testMap()
	testStruct()
}

func TestGofunc() {
	for i := 0; i < 10000; i++ {
		go goFunc(i)
	}
	time.Sleep(time.Second)
}
