package test

import (
	"fmt"
	"strconv"
)

func TestOther() {
	//int转string
	fmt.Println(strconv.Itoa(112))
	//输出 p，Go 语言中，string(n)（n 为整数）会将 n 解释为 Unicode 码点（即 rune 类型），并转换为对应的字符串
	fmt.Println(string(112))
}
