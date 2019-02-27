package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	input, err := in.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return
	}

	input = strings.TrimSpace(strings.ToLower(input))

	if input[0] == 'i' && input[len(input)-1] == 'n' && strings.Contains(input, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
