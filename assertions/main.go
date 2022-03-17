package assertions

import (
	"fmt"
)

type Developer struct {
	Name string
	Age  int
}

func GetDeveloper(name interface{}, age interface{}) Developer {
	return Developer{
		Name: name.(string),
		Age:  age.(int),
	}
}

func Run() {
	fmt.Println("Hello World")

	var name interface{} = "Elliot"
	var age interface{} = 26

	dynamicDev := GetDeveloper(name, age)
	fmt.Println(dynamicDev.Name)
	fmt.Println(dynamicDev.Age)

	detType()
}

func detType() {
	var name interface{} = "Justin"

	value, ok := name.(string)
	fmt.Println(value)
	fmt.Println(ok)

	value2, ok := name.(int)
	fmt.Println(value2)
	fmt.Println(ok)
}
