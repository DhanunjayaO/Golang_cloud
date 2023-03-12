package main

import "fmt"

func main() {
	var num, sum, temp int
	fmt.Print("Enter the number: ")
	fmt.Scanln(&num)
	sum = 0
	temp = num
	for num > 0 {
		rem := num % 10
		sum = sum + (rem * rem * rem)
		num = num / 10

	}
	if temp == sum {
		fmt.Println("given number armstrong")
	} else {
		fmt.Println("not armstrong")
	}

}
