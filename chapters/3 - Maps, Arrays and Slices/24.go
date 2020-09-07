package main

import "fmt"

func main() {

	scores := make([]int, 0, 5)
	c := cap(scores)
	fmt.Println("initial capacity:", c)

	for i := 0; i < 25; i++ {
		scores = append(scores, i)
		// 1) appends the new entry
		// 2) if the capacity have reached the limit, double it's size
		// 3) return a clone array

		// if our capacity has changed,
		//Go had to grow our array to accommodate the new data

		newCapacity := cap(scores) // this scores is a clone of the
		if newCapacity != c {
			c = newCapacity
			fmt.Println("capacity is now:", c)
		}
	}

	fmt.Print(scores)

}
