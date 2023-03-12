package main

import "fmt"

func main() {
	var num, rev, temp int
	fmt.Print("Enter the number: ")
	fmt.Scanln(&num)
	rev = 0
	temp = num
	for num > 0 {
		digit := num % 10
		rev = rev*10 + digit
		num = num / 10
	}
	if temp == rev {
		fmt.Println("given number is palindrome")
	} else {
		fmt.Println("Given number is not palindrome")
	}

}
