package main

import "fmt"

func main() {

	var a, b, c int

	fmt.Print("\nEnter the Three Numbers to find Largest = ")
	fmt.Scanln(&a, &b, &c)

	if a > b && a > c {
		fmt.Println(a, " is Greater Than ", b, " and ", c)
	} else if b > a && b > c {
		fmt.Println(b, " is Greater Than ", a, " and ", c)
	} else if c > a && c > b {
		fmt.Println(c, " is Greater Than ", a, " and ", b)
	} else {
		fmt.Println("Either of them are Equal ")
	}
}
