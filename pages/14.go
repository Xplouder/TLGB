package main

func main() {

	_, exists := power("gokuu") // ignore the first return var because we do not need it
	if exists == false {
		// handle this error case
		log("just a message about the error")
	} else {
		println("success")
	}
}

func log(message string) {
	println(message)
}

func add(a, b int) int { // shorter syntax for parameters with the same type
	return a + b
}

func power(name string) (int, bool) {
	return len(name), len(name)%2 == 0
}
