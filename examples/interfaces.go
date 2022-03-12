package examples

import (
	"fmt"
	"go/types"
	"strconv"
)

type Test types.Chan

type Pet interface {
	GetName() string
	GetAge() int
}

type Cat struct {
	Name  string
	Age   int
	Noise string
}

type Dog struct {
	Name string
	Age  int
}

func (c *Cat) GetName() string {
	return "Cat's Name: " + c.Name
}
func (c *Cat) GetAge() int {
	return c.Age
}
func (c *Cat) PrintNoise() {
	fmt.Println(c.Noise)
}

func (d *Dog) GetName() string {
	return "Dogs Name: " + d.Name
}

func (d *Dog) GetAge() int {
	return d.Age
}

func PrintDetails(e Pet) {
	fmt.Println(e.GetName() + " " + strconv.Itoa(e.GetAge()))
}
