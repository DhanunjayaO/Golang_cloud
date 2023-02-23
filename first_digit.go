package main

import "fmt"

func main() {
	var num int
	fmt.Print("The number is: ")
	fmt.Scanln(&num)
	temp := num
	for temp >= 10 {
		temp = temp / 10

	}
	fmt.Println("The first digit of number is: ", temp)
}
