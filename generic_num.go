package main

import "fmt"

func main() {

	var num, sum, rem int
	fmt.Print("Enter the number: ")
	fmt.Scanln(&num)
	for num >= 10 {
		for sum = 0; num > 0; num = num / 10 {
			rem = num % 10
			sum = sum + rem
		}
		if sum >= 10 {
			num = sum
		} else {
			fmt.Println("The genreric root of number is: ", sum)
		}
	}
}
