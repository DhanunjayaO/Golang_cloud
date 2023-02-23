package main

import "fmt"

func main() {

	var num int
	fmt.Print("The number: ")
	fmt.Scanln(&num)
	fmt.Println("The even numbers upto", num, "is ")
	for i := 0; i <= num; i++ {
		if i%2 == 0 {
			fmt.Print(i)
		}
	}

}
