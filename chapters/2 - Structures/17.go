package main

import "fmt"

type Saiyan struct {
	Name  string
	Power int
}

func main() {
	structureByValue()
	structureByReference("SuperByReference")
	structureByReference("SuperByReference2")
}

func SuperByValue(s Saiyan) {
	s.Power += 10000
}

func SuperByReference(s *Saiyan) {
	s.Power += 10000
}

func SuperByReference2(s *Saiyan) {
	s = &Saiyan{"Gohan", 1000}
}

func structureByValue() {
	goku := Saiyan{"Goku", 9000}
	SuperByValue(goku)
	fmt.Println(goku.Power)
}

func structureByReference(funcName string) {
	goku := &Saiyan{"Goku", 9000}
	if funcName == "SuperByReference" {
		SuperByReference(goku)
	} else {
		SuperByReference2(goku)
	}

	fmt.Println(goku.Power)
}
