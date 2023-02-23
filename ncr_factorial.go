// ncr = n!/r!(n-r)!

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
	t := 1
	fact2 := 1
	for t <= r {
		fact2 = fact2 * t
		t++
	}
	denom := n - r
	j := 1
	fact3 := 1
	for j <= denom {
		fact3 = fact3 * j
		j++
	}
	comb := fact1 / (fact2) * (fact3)
	fmt.Println("The ncr is: ", comb)
}
