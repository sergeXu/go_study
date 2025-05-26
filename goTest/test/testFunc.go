package test

import "fmt"

func testFunc() {
	fmt.Println("testFunc run")
	c := fool("tt", 555)
	fmt.Println(c)

	ret1, ret2 := foo2("dasd", 3123)
	fmt.Println("foo2 = ", ret1, ret2)

	ret1, ret2 = foo3("dasd", 3123)
	fmt.Println("foo3 = ", ret1, ret2)

	ChangeRun()
	testDefer()
}

func fool(a string, b int) int {
	fmt.Println(a, b)
	c := 100
	return c
}

func foo2(a string, b int) (int, int) {
	fmt.Println(a, b)
	return 666, 777
}

func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println(a, b)
	r1, r2 = 12, 13
	return r1, r2
}

func changeValue(p *int) {
	*p = 10
}
func ChangeRun() {
	var a int = 1
	changeValue(&a)
	fmt.Println("a= ", a)
}

func testDefer() {
	//类似析构函数  finally 方法
	defer fmt.Println("testDefer end1")
	defer fmt.Println("testDefer end2")
	fmt.Println("testDefer run1")
	fmt.Println("testDefer run2")
}
