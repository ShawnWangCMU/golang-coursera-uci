package main

import (
	"bufio"
	json2 "encoding/json"
	"fmt"
	"os"
	"strings"
)

/*
 * Write a program which prompts the user to first enter a name,
 * and then enter an address.
 * Your program should create a map and add the name and address to the map
 * using the keys “name” and “address”, respectively.
 * Your program should use Marshal() to create a JSON object from the map,
 * and then your program should print the JSON object.
 */

func main() {
	m := make(map[string]string)
	in := bufio.NewReader(os.Stdin)

	fmt.Print("Input your name: ")
	name, err := in.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	if err != nil {
		fmt.Println("Invalid input: ", err)
		return
	}

	fmt.Print("Input your address: ")
	address, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid input: ", err)
		return
	}

	m["name"] = strings.TrimSpace(name)
	m["address"] = strings.TrimSpace(address)

	json, err := json2.Marshal(m)
	if err != nil {
		fmt.Println("Error marshalling json: ", err)
		return
	}

	fmt.Println(string(json))
}
