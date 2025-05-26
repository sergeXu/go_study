package main

import (
	//包更多的类似于类的
	//别名导入
	f "fmt"

	//匿名导入，可不使用
	_ "math"

	//全部导入的方式
	. "goTest/test"
	//导入的是文件路径
	"goTest/util"
)

func main() {
	f.Println("Hello World")
	TestStart()
	//使用时调用的是package包名
	utils.TestUtil()
}
