package main

import "fmt"

func main() {
	var n, r int
	fmt.Print("Enter the number: ")
	fmt.Scanln(&n)
	fmt.Print("Enter the number: ")
	fmt.Scanln(&r)
	i := 1
	fact1 := 1
	for i <= n {
		fact1 = fact1 * i
		i++
	}
	denom := n - r
	j := 1
	fact2 := 1
	for j <= denom {
		fact2 = fact2 * j
		j++
	}
	per := fact1 / fact2
	fmt.Println("The npr is: ", per)
}
