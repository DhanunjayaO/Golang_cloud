package main

import "fmt"

func main() {

	var num int
	fmt.Print("The number: ")
	fmt.Scanln(&num)
	if num%2 == 0 {
		fmt.Println("Given number is even", num)
	} else {
		fmt.Println("given number is Odd", num)
	}

}
