package main

import "fmt"

func main() {

	// Arrays ----------------------------------------------------------------------------------------------------------

	var scores [30]int
	scores[0] = 339

	scores2 := [4]int{9001, 9333, 212, 33}

	for index, value := range scores {
		fmt.Println(index, value)
	}

	for index, value := range scores2 {
		fmt.Println(index, value)
	}

	// error test
	//scores[100] = 123

	// Slices ----------------------------------------------------------------------------------------------------------

	//slice := []int{1, 4, 293, 4, 9}
	//slices2 := make([]int, 10)
	//scores3 := make([]int, 0, 10)      // creates a slice with a length of 0 but with a capacity of 10

	slices := make([]int, 0, 10)
	//slices[7] = 9033                   // panic: runtime error: index out of range [7] with length 0
	fmt.Println(slices)

	slices2 := make([]int, 0, 10)
	slices2 = append(slices2, 5)
	fmt.Println(slices2) // prints [5]

	slices3 := make([]int, 0, 10)
	slices3 = slices3[0:8]
	slices3[7] = 9033
	fmt.Println(slices3)

}
