package main

type Saiyan struct {
	Name  string
	Power int
}

func main() {

	// basic structure initialization
	goku := Saiyan{
		Name:  "Goky",
		Power: 9000,
	}

	// other valid ways of set a structure data
	goku = Saiyan{}
	goku = Saiyan{Name: "Goku"}
	goku.Power = 9000
	goku = Saiyan{"Goku", 9000} // this one skip the field names and rely on the order of the field declarations
}
