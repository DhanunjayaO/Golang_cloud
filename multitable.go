package main

import "fmt"

func main() {

	var num, multi int
	/*fmt.Print("Enter the number:")
	fmt.Scanln(&num)*/
	fmt.Print("Enter the multiplier: ")
	fmt.Scanln(&multi)
	num = 1

	for num <= 5 {

		//product := num * multi

		fmt.Println(num, "*", multi, "=", (num * multi))

		num++

	}

}
