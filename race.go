package main

import (
	"fmt"
	"time"
)

// routine1 and routine2 races to add x by 1 and by 2
// when routine1 tries to print the value of x,
// there're two possibilities: 1 or 3, where x can already been increased by routine2 or not
// 10ms sleep is added to enlarge the race condition and make it easier to occur when run multiple times
// 100ms sleep in main() is used to wait both thread to finish working and print stuff
var x int

func routine1() {
	time.Sleep(10 * time.Millisecond)
	x = x+1
	fmt.Println(x)
}

func routine2() {
	time.Sleep(10 * time.Millisecond)
	x = x+2
}

func main() {
	for i := 0; i < 10; i++ {
		x = 0
		fmt.Printf("iteration %d : ", i)
		go routine1()
		go routine2()
		time.Sleep(100 * time.Millisecond)
	}
}
