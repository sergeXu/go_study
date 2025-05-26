package test

import "fmt"

// 数据类型声明,int别名处理
type myInt int

type Book struct {
	title  string
	author string
}

func testStruct() {
	var a myInt
	a = 14
	fmt.Println("a=", a)
	fmt.Printf("type of a = %T \n", a)

	var book1 Book
	book1.title = "Go Programming Language"
	book1.author = "James Bond"
	changeBook(book1)
	fmt.Printf("book1=%v\n", book1)

	changeBook2(&book1)
	fmt.Printf("book1=%v\n", book1)
}

func changeBook(book Book) {
	book.title = "Test Change"
}

func changeBook2(book *Book) {
	//指针传递
	book.title = "Test Change"
	//等效
	(*book).title = "TTT"
}
