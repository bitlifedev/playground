package examples

import (
	"fmt"
	"runtime"
)

func Main() {
	fmt.Println("Welcome to the Playground running " + runtime.Version())
	fmt.Println(Convert(uint8(15)))
	DaysOfTheWeek()
	StrutExample()
	fmt.Println(Run())

	//Work with employees
	employee := Employee{}
	employee.UpdateName("Justin")
	employee.PrintName()

	cat := &Cat{Name: "Bit", Age: 6, Noise: "Meow"}
	dog := &Dog{Name: "Lulu"}
	PrintDetails(cat)
	cat.PrintNoise()
	PrintDetails(dog)

	//Terminal Input
	//examples.GetHello()

	fmt.Println("Running console commands")
	Exec()

	fmt.Println("Running Json")
	RunJson()
	fmt.Println("Running Xml")
	RunXML()

	RunInit()
}
