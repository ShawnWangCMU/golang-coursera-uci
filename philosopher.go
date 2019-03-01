package main

import (
	"fmt"
	"sync"
	"time"
)

/*
 * Implement the dining philosopher’s problem with the following constraints/modifications.
 * 1. There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
 * 2. Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
 * 3. The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
 * 4. In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
 * 5. The host allows no more than 2 philosophers to eat concurrently.
 * 6. Each philosopher is numbered, 1 through 5.
 * 7. When a philosopher starts eating (after it has obtained necessary locks)
 *    it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
 * 8. When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>”
 *    on a line by itself, where <number> is the number of the philosopher.
 */

type Philo struct {
	id int
	chopsticksLeft, chopsticksRight sync.Mutex
}

func (philo Philo) Eat(permission chan int, finish chan int, numEats int) {
	for i := 0; i < numEats; i++ {
		// acquire permission from host
		<- permission

		philo.chopsticksLeft.Lock()
		philo.chopsticksRight.Lock()
		fmt.Printf("starting to eat %d\n", philo.id)
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("finishing to eat %d\n", philo.id)
		philo.chopsticksRight.Unlock()
		philo.chopsticksLeft.Unlock()

		// report finish eating to host
		finish <- philo.id
	}
}

func host(permission chan int, finish chan int, totalEats int, wg *sync.WaitGroup) {
	defer wg.Done()
	eatsDone := 0

	// init permission to have 2 eat slots
	permission <- 0
	permission <- 0
	for {
		select {
		case <- finish:
			permission <- 0
			eatsDone++
		}

		if eatsDone >= totalEats {
			break
		}
	}
}
func main() {
	// macro for total num of philos and times to eat
	var numPhilo int = 5
	var numEats int = 3

	// channels for communication
	permisson := make(chan int, 2)
	finish := make(chan int, 2)

	// init philos and chopsticks
	chopsticks := make([]sync.Mutex, numPhilo)
	philos := make([]Philo, 0)
	for i := 0; i < numPhilo; i++ {
		philos = append(philos, Philo{i, chopsticks[i], chopsticks[(i+1)%numPhilo]})
		go philos[i].Eat(permisson, finish, numEats)
	}

	// init and wait on host to finish
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go host(permisson, finish, numPhilo * numEats, wg)
	wg.Wait()
}
