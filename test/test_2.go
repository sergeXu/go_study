package test

import (
	"fmt"
	"math/big"
	"reflect"
	"unicode/utf8"
)

// 全局变量不做使用检查
var Globle_Char = "testChar"

func Test2() {
	fmt.Println("Test3 start")
	//变量赋值部分
	var length, size int
	var success bool
	var (
		lenth, siz int
		succ       bool
	)
	length = 20
	size = 30
	lenth, siz = 20, 30
	fmt.Println(lenth, siz, success, length, size, succ)
	lenth, siz = siz, lenth
	fmt.Println(lenth, siz, success, length, size, succ)

	a := 10
	//a := 30
	fmt.Println(a, Globle_Char)

	//常量部分
	const abc = 10
	const (
		n, k = 15, 10
		m    = 20
		q
	)
	fmt.Println(abc, m, n, k, q)

	const (
		_ = iota + 100
		Sunday
		Monday
		Tuesday
		Wednesday = 'y'
		Thursday  = iota
		Friday
		Saturday
	)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}

func Test3() {
	fmt.Println("Test3 start")
	//整形和浮点
	var a int
	var b uint
	var a1 int16
	var b1 uint16
	var c float32
	var d float64
	c = 1.0
	d = 1.21
	e := 1.4
	fmt.Println(a, b, a1, b1, c, d, e)
	fmt.Println("c的类型为：", reflect.TypeOf(c))
	fmt.Println("e的类型为：", reflect.TypeOf(e))
	//浮点数比较
	result := big.NewFloat(d).Cmp(big.NewFloat(e))
	if result < 0 {
		fmt.Println("d<e")
	} else if result > 0 {
		fmt.Println("d>e")
	} else {
		fmt.Println("d=e")
	}
	//字符类型
	var cc = '中'
	fmt.Println(cc)
	fmt.Printf("%c \n", cc)
	fmt.Println(string(cc))
	//字符串类型
	var s string
	s = "go 语言字符串"
	// %v 打印详细内容
	fmt.Printf("%v \n", s)
	fmt.Printf("s数据类型为 %s \n", reflect.TypeOf(s))
	//字节数组
	for i := 0; i < len(s); i++ {
		fmt.Printf("%X ", s[i])
	}
	//乱码打印
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
	fmt.Println("字符长度", utf8.RuneCountInString(s))
	for _, c := range s {
		fmt.Printf("%T %X %c     ", c, c, c)
	}

	//数组类型 todo
}
