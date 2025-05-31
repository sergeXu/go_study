package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func reflectNum(obj any) {
	fmt.Println("type =", reflect.TypeOf(obj))
	fmt.Println("value =", reflect.ValueOf(obj))
}

type User struct {
	Id   int
	Name string
	Age  int
}

func (this *User) call() {
	fmt.Println("call user ", *this)
}

func doFiledAndMethod(input any) {
	//获取input的type ,如果是指针类型，需要加Elem()
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType =", inputType)
	//获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue =", inputValue)
	//获取变量对应的Kind
	kd := inputValue.Kind()
	fmt.Println("kd =", kd)

	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("field[%d]:name=%s,type=%v,value=%v\n", i, field.Name, field.Type, value)
	}

}

///标签

type Resume struct {
	Name    string `info:"name" doc:"我的名字"`
	Sex     string `info:"sex"`
	Address string `info:"address" city:"杭州"`
}

func findTag(obj any) {
	//获取type elem()取指针指向的类型
	t := reflect.TypeOf(obj).Elem()
	for i := 0; i < t.NumField(); i++ {
		tagInfo := t.Field(i).Tag.Get("info")
		tagDoc := t.Field(i).Tag.Get("doc")
		tagCity := t.Field(i).Tag.Get("city")
		fmt.Println("tagInfo ", tagInfo)
		fmt.Println("tagDoc ", tagDoc)
		fmt.Println("tagCity ", tagCity)
	}
}

// Tag支持的类型
type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

func TestReflect() {
	var num_pi float64 = 3.1415926
	reflectNum(num_pi)

	user := User{1, "tom", 18}
	user.call()
	doFiledAndMethod(user)

	resume := Resume{
		Name:    "tom",
		Sex:     "man",
		Address: "beijing",
	}
	findTag(&resume)

	//json类型的应用
	movie := Movie{"喜剧之王", 1992, 80, []string{"周星驰", "张柏芝"}}
	str, err := json.Marshal(movie)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("json str = %s\n", str)

	myMovie := Movie{}
	err = json.Unmarshal(str, &myMovie)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("json movie = %v\n", myMovie)
}
