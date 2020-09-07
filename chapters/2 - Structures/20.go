package main

import "fmt"

type Saiyan struct {
	*Person // [*]
	Power   int
}

type Person struct {
	Name string
}

func main() {
	goku := &Saiyan{
		Person: &Person{"Goku"}, // [*]
		Power:  9001,
	}
	goku.Introduce()        // composed version
	goku.Person.Introduce() // non composed version

	fmt.Println("---------")

	goku = &Saiyan{
		Person: &Person{"Goku"},
	}
	fmt.Println(goku.Name) // we can implicitly access the fields and functions of the composed type because we didnt give an explicit structure field name [*]
	fmt.Println(goku.Person.Name)
}

func (p *Person) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

// a way to "hack" the non overloading Go capability - composed version is
func (s *Saiyan) Introduce() {
	fmt.Printf("Hi, I'm %s. Ya! \n", s.Name)
}
