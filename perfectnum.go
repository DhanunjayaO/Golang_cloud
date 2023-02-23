package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter the perfect number: ")
	fmt.Scanln(&num)
	i := 1
	sum := 0
	for i < num {
		if num%i == 0 {
			sum = sum + i
		}
		i++
	}
	if sum == num {
		fmt.Println(num, "Perfect number")
	} else {
		fmt.Println(num, "not perfect number")
	}

}
