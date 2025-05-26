package test

import "fmt"

var gA int = 100
var gB int = 200

func TestVar() {
	fmt.Println("TestVar Start")
	var a int
	fmt.Println("a=", a)
	var b = 100
	fmt.Println("b=", b)
	var c = 100
	fmt.Println("c=", c)
	d := 100
	fmt.Println("d=", d)
	fmt.Printf("type of a = %T\n", a)
	fmt.Printf("type of b = %T\n", b)
	fmt.Printf("type of c = %T\n", c)
	fmt.Printf("type of d = %T\n", d)

	var bb = "abcd"
	fmt.Printf("bb =%s , Type of bb=%T\n", bb, bb)

	g := 3.14
	fmt.Printf("g =%s , Type of g=%T\n", g, g)

	var xx, yy int = 100, 200
	fmt.Printf("xx =%d yy =%d \n", xx, yy)
	var kk, ll = 100, "abcd"
	fmt.Println("kk=", kk, "ll=", ll)

	var (
		vv int = 100
		jj     = true
	)
	fmt.Println("vv=", vv, ", jj=", jj)

	fmt.Println("gA = ", gA)
}
