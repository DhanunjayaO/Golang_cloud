package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter the number: ")
	fmt.Scanln(&num)
	for i := 20; i >= num; i-- {
		fmt.Println(i)
	}
}
