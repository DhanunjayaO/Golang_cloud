package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter the digit: ")
	fmt.Scanln(&num)
	res := 1
	for num > 0 {
		rem := num % 10
		res = res * rem
		num = num / 10
	}
	fmt.Println("the product of result is: ", res)
}
