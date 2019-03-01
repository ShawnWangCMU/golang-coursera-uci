package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Write a program to sort an array of integers.
 * The program should partition the array into 4 parts, each of which is sorted by a different goroutine.
 * Each partition should be of approximately equal size.
 * Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.
 *
 * The program should prompt the user to input a series of integers.
 * Each goroutine which sorts ¼ of the array should print the subarray that it will sort.
 * When sorting is complete, the main goroutine should print the entire sorted list.
 */

var numSlices int

func sortSlice(subArray []int, c chan []int) {
	sort.Ints(subArray)
	fmt.Println(subArray)
	c <- subArray
}

func merge(subSlices [][]int) []int {
	sortedArray := make([]int, 0)
	index := make([]int, len(subSlices))
	for {
		nextIdx := 0
		nextVal := math.MaxInt32
		for i := 0; i < len(subSlices); i++ {
			if index[i] < len(subSlices[i]) && subSlices[i][index[i]] < nextVal {
				nextIdx = i
				nextVal = subSlices[i][index[i]]
			}
		}

		sortedArray = append(sortedArray, nextVal)
		index[nextIdx]++

		done := true
		for i := 0; i < len(subSlices); i++ {
			if index[i] < len(subSlices[i]) {
				done = false
				break
			}
		}
		if done {
			break
		}
	}

	return sortedArray
}

func main() {
	numSlices = 4

	in := bufio.NewReader(os.Stdin)
	fmt.Println("Input integers, separated by space: ")
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
	}

	c := make(chan []int, numSlices)
	subLen := int(len(numArray) / numSlices)
	for i := 0; i < numSlices; i++ {
		subArray := numArray[i*subLen:(i+1)*subLen]
		go sortSlice(subArray, c)
	}

	sortedSlices := make([][]int, numSlices)
	for i := 0; i < numSlices; i++ {
		sortedSlices[i] = <- c
	}

	sortedArray := merge(sortedSlices)
	fmt.Println("Sorted order:")
	fmt.Println(sortedArray)
}
