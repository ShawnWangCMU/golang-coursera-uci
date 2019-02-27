package main

import "fmt"

func main() {
	var numFloat float32
	var numInt int

	fmt.Println("input a floating point number:")
	num, err := fmt.Scan(&numFloat)

	if err != nil {
		fmt.Println("please input a valid floating point number!")
		return
	}
	if num != 1 {
		fmt.Println("please input exactly one number!")
		return
	}

	numInt = int(numFloat)
	fmt.Println("truncated to integer equals ", numInt)
}