package main

import (
	"fmt"
	"strconv"
)

type Saiyan struct {
	Name   string
	Power  int
	Father *Saiyan
}

// *Saiyan is the receiver of the Super method
func (s *Saiyan) Super() {
	s.Power += 10000
}

// Apparently structures do not have a constructor, so i guess this is valid (it is not, see the main function [*] )
func (s *Saiyan) constructor(name string, power int) *Saiyan {
	return &Saiyan{
		Name:  name,
		Power: power,
	}
}

func main() {
	goku := &Saiyan{"Goku", 9001, nil}
	goku.Super()
	fmt.Println(goku.Power) // will print 19001

	fmt.Println("---------")

	vegeta := NewSaiyan("Vegeta", 8000)
	fmt.Println(vegeta.Power)
	//vegeta := Saiyan.constructor("Vegeta", 8000) // [*] Not valid because any structure method invocation needs to be called from a pointer

	fmt.Println("---------")

	gohan := &Saiyan{
		Name:   "Gohan",
		Power:  1000,
		Father: goku,
	}

	fmt.Println("Gohan Power:", gohan.Power)
	fmt.Println("Goku power: " + strconv.Itoa(gohan.Father.Power))

}

func NewSaiyan(name string, power int) *Saiyan {
	//return new(Saiyan) // a syntax sugar for &Saiyan{}
	return &Saiyan{
		Name:  name,
		Power: power,
	}
}
