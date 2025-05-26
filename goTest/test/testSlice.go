package test

import "fmt"

func testSlice() {
	fmt.Println("testSlice run")

	//固定长度数组
	var myArray [5]int
	//数组遍历
	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}
	fmt.Println("myArray: ", myArray)

	myArray2 := [3]int{1, 3, 4}
	for index, value := range myArray2 {
		fmt.Println("index: ", index, myArray2[index], value)
	}
	fmt.Println("myArray2 ", myArray2)
	//[3]int 类型只能给精确对应的入参函数传， 值类型传递
	fmt.Printf("type of myArray2 %T\n", myArray2)

	mySlice := []int{1, 2, 3}
	fmt.Println("mySlice ", mySlice)
	fmt.Printf("type of mySlice %T\n", mySlice)
	for _, value := range mySlice {
		println("value = ", value)
	}

	var mySlice2 []int
	if mySlice2 == nil {
		println("mySlice2 == nil")
	}
	//初始化容量空间
	mySlice2 = make([]int, 2)
	mySlice2[0] = 100
	fmt.Printf("len =%d,slice =%v \n", len(mySlice2), mySlice2)

	//var mySlice3 []int = make([]int, 5)
	//等效
	mySlice3 := make([]int, 5)
	fmt.Printf("len =%d,slice3 =%v \n", len(mySlice3), mySlice3)

	//cap容量为5，预留空间
	var numbers = make([]int, 3, 5)
	fmt.Printf("len =%d,cap = %d,slice3 =%v \n", len(numbers), cap(numbers), numbers)

	//末尾追加
	numbers = append(numbers, 111)
	numbers = append(numbers, 222)
	fmt.Printf("len =%d,cap = %d,slice3 =%v \n", len(numbers), cap(numbers), numbers)
	//满数组增加，cap自动扩容翻倍
	numbers = append(numbers, 333)
	fmt.Printf("len =%d,cap = %d,slice3 =%v \n", len(numbers), cap(numbers), numbers)

	s := []int{1, 2, 3}
	//数据截取，但底层仍用同一数组，会被一起修改
	s1 := s[0:2]
	fmt.Println("S1 = ", s1)

	s2 := make([]int, 3)
	//深拷贝函数
	copy(s2, s)
	fmt.Println("s2 = ", s2)

}
