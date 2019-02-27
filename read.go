package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
 * Write a program which reads information from a file and represents it in a slice of structs.
 * Assume that there is a text file which contains a series of names.
 * Each line of the text file has a first name and a last name,
 * in that order, separated by a single space on the line.
 *
 * Your program will define a name struct which has two fields,
 * fname for the first name, and lname for the last name.
 * Each field will be a string of size 20 (characters).
 *
 * Your program should prompt the user for the name of the text file.
 * Your program will successively read each line of the text file
 * and create a struct which contains the first and last names found in the file.
 * Each struct created will be added to a slice, and after all lines have been read from the file,
 * your program will have a slice containing one struct for each line in the file.
 *
 * After reading all lines from the file, your program should iterate through your slice of structs
 * and print the first and last names found in each struct.
 */

type name struct {
	fname string
	lname string
}

func (n name) String() string {
	return fmt.Sprintf("%v %v", n.fname, n.lname)
}

func main() {
	var filename string
	fmt.Print("Input path to the text file: ")
	n, err := fmt.Scan(&filename)
	if err != nil || n > 1 {
		fmt.Println("invalid input")
		return
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	nameList := make([]name, 0, 5)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		nameList = append(nameList, name{fields[0], fields[1]})
	}

	for _, v := range nameList {
		fmt.Println(v)
	}
}
