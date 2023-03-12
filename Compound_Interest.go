package main

import (
	"fmt"
	"math"
)

func main() {
	var pamount, roi, time, tamount, interest float64
	fmt.Print("The principle amount: ")
	fmt.Scanln(&pamount)
	fmt.Print("The rate of interest: ")
	fmt.Scanln(&roi)
	fmt.Print("the time period: ")
	fmt.Scanln(&time)
	tamount = pamount * (math.Pow((1 + roi/100), time))
	interest = tamount - pamount
	fmt.Println("\nTotal amount: ", tamount)
	fmt.Println("\nInterest is: ", interest)
}
