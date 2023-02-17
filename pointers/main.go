package main

import "fmt"

func main() {
	/*
		var i int = 42
		j := &i
		fmt.Print(i, *j, "\n")
	*/

	name := "Abhishek"
	title := "Mr"
	display(&name, title)
	fmt.Print(title, "::", name, "\n")

}

func display(name *string, title string) {
	*name = "Abi"
	fmt.Print(title, ":", *name, "\n")
}
