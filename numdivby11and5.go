package main

import "fmt"

func main() {

	num := 1

	for num <= 100 {
		num++
		if (num%5) == 0 && (num%11) == 0 {
			fmt.Println(num, "divisible number")

		}
	}
}
