package sorting

import (
	"fmt"
	"sort"
)

type Programmer struct {
	Age int
}

type byAge []Programmer

func (p byAge) Len() int {
	return len(p)
}

func (p byAge) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p byAge) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}

func Run() {

	fmt.Println("Go Sorting")

	// our unsorted int array
	//unsorted := []int{1, 3, 2, 6, 3, 4}
	//fmt.Println(unsorted)
	//sort.Ints(unsorted)
	//fmt.Println(unsorted)

	programmers := []Programmer{
		{Age: 30},
		{Age: 20},
		{Age: 50},
		{Age: 1000},
	}

	sort.Sort(byAge(programmers))
	fmt.Println(programmers)
	fmt.Println(sort.IsSorted(byAge(programmers)))
	sort.Sort(sort.Reverse(byAge(programmers)))
	fmt.Println(programmers)
	fmt.Println(sort.IsSorted(byAge(programmers)))

}
