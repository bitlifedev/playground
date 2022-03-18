package pointers

import "fmt"

func addOne(a *int) {
	*a++
}

func Run() {
	fmt.Println("Go Pointers")
	a := new(int)
	*a = 42
	fmt.Println(*a)
	fmt.Println(&a)
	addOne(a)
	fmt.Println(*a)
	fmt.Println(&a)

}
