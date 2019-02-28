package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
 * Write a Bubble Sort program in Go.
 * The program should prompt the user to type in a sequence of up to 10 integers.
 * The program should print the integers out on one line, in sorted order, from least to greatest.
 * Use your favorite search tool to find a description of how the bubble sort algorithm works.
 *
 * As part of this program, you should write a function called BubbleSort()
 * which takes a slice of integers as an argument and returns nothing.
 * The BubbleSort() function should modify the slice so that the elements are in sorted order.
 *
 * A recurring operation in the bubble sort algorithm is the Swap operation
 * which swaps the position of two adjacent elements in the slice.
 * You should write a Swap() function which performs this operation.
 * Your Swap() function should take two arguments, a slice of integers and an index value i
 * which indicates a position in the slice. The Swap() function should return nothing,
 * but it should swap the contents of the slice in position i with the contents in position i+1.
 */

func Swap(numArray []int, i int) {
	temp := numArray[i]
	numArray[i] = numArray[i+1]
	numArray[i+1] = temp
}

func BubbleSort(numArray []int) {
	for i := 0; i < len(numArray) - 1; i++ {
		for j := 0; j < len(numArray) - i - 1; j++ {
			if numArray[j] > numArray[j+1] {
				Swap(numArray, j)
			}
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Println("Input upto 10 integers, separated by space: ")
	input, err := in.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	numArray := make([]int, 0)
	for _, numStr := range strings.Fields(input) {
		num, err := strconv.Atoi(numStr)
		if err == nil {
			numArray = append(numArray, num)
		}
		if len(numArray) >= 10 {
			break
		}
	}

	BubbleSort(numArray)
	fmt.Println("Sorted order:")
	fmt.Println(numArray)
}
