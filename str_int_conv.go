package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 123456789

	num1 := strconv.Itoa(num)

	fmt.Println(len(num1))

	num2 := fmt.Sprintf("%v", num)

	fmt.Println(len(num2))

}
