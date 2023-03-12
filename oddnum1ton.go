package main

import "fmt"

func main() 
	var num int
	fmt.Print("Enter the number: ")
	fmt.Scanln(&num)
	fmt.Println("The Odd numbers are")
	i := 1
	for i < num {
		if i%2 == 1 {
			fmt.Print(i)
		}
		i++
	}

}
