package main

import "fmt"

func main() {
	power := 9000 // short variable declaration
	fmt.Printf("It's over %d\n", power)

	//power := 9001 // can not redeclare the variable (compiler error)
	power = 9001 // instead, to change the value just assign it "normally"
	fmt.Printf("It's over %d\n", power)

	name, power := "Goku", 9000 // redeclaring the power variable this way, when including a new variable does not produce an error
	fmt.Printf("%s's power is over %d \n ", name, power)

	//name_2, power_2 := "Goku" , 1000 // result in an compiler error when the variable is not used
	//fmt.Printf( "default power is %d \n " , power_2)
}
