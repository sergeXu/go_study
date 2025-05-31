package main

import "fmt"

type Hero struct {
	Name  string
	Ad    int
	Level int
}

// 方法是作用于特定类型的函数
func (this *Hero) show() {
	fmt.Println("Name =", this.Name)
	fmt.Println("Ad =", this.Ad)
	fmt.Println("Level =", this.Level)
}
func (this *Hero) getName() string {
	//fmt.Println("Name =" ,this.Name)
	return this.Name
}

func (this *Hero) setName(newName string) {
	this.Name = newName
}

// /////////////////////////
type Human struct {
	Name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()")
}

func (this *Human) Walk() {
	fmt.Println("Human.Walk()")
}

func (this *Human) Show() {
	fmt.Println("name =", this.Name, "sex =", this.sex)
}

type SuperMan struct {
	Human //SuperMan 继承了Human的方法
	level int
}

func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()")
}

func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()")
}

func (this *SuperMan) Show() {
	fmt.Println("SuperMan.show()")
	this.Human.Show()
	fmt.Println("Level =", this.level)
}

// ////////////////////////
type AnimalIf interface {
	Sleep()
	GetColor() string //获取动物的颜色
	GetType() string  //获取动物的类型
}

type Cat struct {
	color string
}

func (this *Cat) GetColor() string {
	return this.color
}
func (this *Cat) GetType() string {
	return "cat"
}
func (this *Cat) Sleep() {
	fmt.Println("Cat.Sleep()")
}

type Dog struct {
	color string
}

func (this *Dog) GetColor() string {
	return this.color
}
func (this *Dog) GetType() string {
	return "dog"
}
func (this *Dog) Sleep() {
	fmt.Println("Dog.Sleep()")
}

func showAnimal(animal AnimalIf) {
	animal.Sleep()
	fmt.Println("color=", animal.GetColor())
	fmt.Println("type=", animal.GetType())
}

// //////////////////
// 空接口
func myFunc(obj interface{}) {
	fmt.Println("MyFunc is called...")
	fmt.Println(obj)
	//interface{} 如何区分 原始类型
	//类型断言
	value, ok := obj.(string)
	if !ok {
		fmt.Println("转换失败")
	} else {
		fmt.Println("转换成功", value, "type =", "string")
	}
	value2, ok := obj.(Cat)
	if !ok {
		fmt.Println("转换失败")
	} else {
		fmt.Println("转换成功", value2, "type =", "Cat")
	}
}

// 等效
func myFunc2(obj any) {
	fmt.Println("MyFunc2 is called...")
}

func TestObj() {
	hero := Hero{
		Name:  "zhang3",
		Ad:    100,
		Level: 1,
	}
	hero.show()

	hero.setName("li4")
	hero.show()
	////// 继承逻辑使用

	human := Human{
		Name: "tom",
		sex:  "man",
	}
	human.Eat()
	human.Walk()

	//不同初始化方式
	super := SuperMan{Human{"jack", "man"}, 888}
	super.Show()

	var super2 SuperMan
	super.Human = Human{"jack", "man"}
	super2.level = 888
	super2.Show()

	var super3 SuperMan
	super3.Name = "jack"
	super3.sex = "man"
	super3.level = 888
	super3.Show()

	super.Eat()
	super.Walk()
	super.Fly()

	/////////////// 多态部分
	//定义接口类型的变量
	var animal AnimalIf

	//animal = &Cat{"green"}
	//animal.Sleep()
	//
	//animal = &Dog{"green"}
	//animal.Sleep()

	cat := Cat{"blue"}
	//接口是指针类型，需要处理
	animal = &cat
	showAnimal(animal)
	dog := &Dog{"yellow"}
	//接口是指针类型
	animal = dog
	showAnimal(animal)
	///
	// 空接口
	myFunc(cat)
	myFunc("tom")
	myFunc(10)
}

// //////////////////////////////
// main包下的main是可执行的
func main() {
	//TestObj()
	//TestReflect()
	testChannel()
}
