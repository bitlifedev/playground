package examples

import (
	"fmt"
)

func Run() string {
	fmt.Println(myfunction("Justin", "Mifkovich"))

	fullName, err := myfunction2("Ellon", "Musk")
	if err != nil {
		fmt.Println("Handle Error Case" + err.Error())
	}
	fmt.Println(fullName)

	myFunc := addOne()
	fmt.Println(myFunc()) // 2
	fmt.Println(myFunc()) // 3
	fmt.Println(myFunc()) // 4
	fmt.Println(myFunc()) // 5

	myVariadicFunction("hello", "world")

	return "Done\n"
}

func myfunction(firstName string, lastName string) string {
	fullname := firstName + " " + lastName
	return fullname
}

func myfunction2(firstName string, lastName string) (string, error) {
	return firstName + " " + lastName, nil
}

func addOne() func() int {
	var x int
	// we define and return an anonymous function which in turn returns an integer value
	return func() int {
		// this anonymous function has access to the x variable defined in the parent function
		x++
		return x + 1
	}
}

func myVariadicFunction(args ...string) {
	fmt.Println(args)
}
