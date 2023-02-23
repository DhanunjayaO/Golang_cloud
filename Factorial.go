package main

import "fmt"

func main() {
	var num int
	fmt.Print("The factorial number is: ")
	fmt.Scanln(&num)
	temp := 1
	for i := 1; i <= num; i++ {
		temp = temp * i
	}
	fmt.Println(temp)

}
