package test

import "fmt"

func testMap() {
	fmt.Println("testMap Run")
	var stringMap map[string]string
	if stringMap == nil {
		fmt.Println("stringMap is nil")
	}
	//需要初始化 make
	stringMap = make(map[string]string, 5)
	stringMap["a"] = "ttt1"
	stringMap["b"] = "ttt2"
	stringMap["v"] = "ttt2"

	fmt.Println("stringMap", stringMap)
	fmt.Println("stringMap[a]", stringMap["a"])

	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "python"
	myMap2[3] = "c++"
	fmt.Println("myMap2", myMap2)

	//初始化 + 默认赋值
	myMap3 := map[string]string{
		"one":   "java",
		"two":   "c++",
		"three": "python",
	}
	fmt.Println("myMap3", myMap3)

	//添加
	myMap3["four"] = "php"

	//删除
	delete(myMap3, "two")

	//修 改
	myMap3["one"] = "java2"

	//遍历
	for key, value := range myMap3 {
		fmt.Println("key:", key, "value:", value)
	}

	printMyMap(myMap3)
}

func printMyMap(myMap map[string]string) {
	fmt.Println("printMyMap run")
	//引用传递（make()处理的）
	for key, value := range myMap {
		fmt.Println("key=", key, "value=", value)
	}
}
