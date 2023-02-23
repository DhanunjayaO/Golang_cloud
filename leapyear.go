package main

import "fmt"

func main() {

	var num int
	fmt.Print("Enter the year: ")
	fmt.Scanln(&num)
	if num%400 == 0 || (num%4 == 0 && num%100 != 0) {
		fmt.Println(num, "is leap year")
	} else {
		fmt.Println(num, "is not a leap year")
	}
}
