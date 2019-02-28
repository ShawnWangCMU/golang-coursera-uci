package animal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
 * Write a program which allows the user to create a set of animals and to get information about those animals.
 * Each animal has a name and can be either a cow, bird, or snake.
 * With each command, the user can either create a new animal of one of the three types,
 * or the user can request information about an animal that he/she has already created.
 * Each animal has a unique name, defined by the user.
 * Note that the user can define animals of a chosen type, but the types of animals are restricted to
 * either cow, bird, or snake. The following table contains the three types of animals and their associated data.
 *
 * Animal	Food eaten	Locomtion method	Spoken sound
 * cow		grass		walk				moo
 * bird		worms		fly					peep
 * snake	mice		slither				hsss
 *
 * Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
 * Your program should accept one command at a time from the user, print out a response,
 * and print out a new prompt on a new line. Your program should continue in this loop forever.
 * Every command from the user must be either a “newanimal” command or a “query” command.
 *
 * Each “newanimal” command must be a single line containing three strings.
 * The first string is “newanimal”. The second string is an arbitrary string which will be the name of the new animal.
 * The third string is the type of the new animal, either “cow”, “bird”, or “snake”.
 * Your program should process each newanimal command by
 * creating the new animal and printing “Created it!” on the screen.
 *
 * Each “query” command must be a single line containing 3 strings.
 * The first string is “query”. The second string is the name of the animal.
 * The third string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
 * Your program should process each query command by printing out the requested data.
 *
 * Define an interface type called Animal which describes the methods of an animal.
 * Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(),
 * which take no arguments and return no values. The Eat() method should print the animal’s food,
 * the Move() method should print the animal’s locomotion,
 * and the Speak() method should print the animal’s spoken sound.
 * Define three types Cow, Bird, and Snake. For each of these three types,
 * define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake all satisfy the Animal interface.
 * When the user creates an animal, create an object of the appropriate type.
 * Your program should call the appropriate method when the user issues a query command.
 */

type Animal interface {
    Eat()
	Move()
	Speak()
}

// Cow implements Animal
type Cow struct {
	name string
}

func (cow Cow) Eat() {
	fmt.Println("grass")
}

func (cow Cow) Move() {
	fmt.Println("walk")
}

func (cow Cow) Speak() {
	fmt.Println("moo")
}

// Bird implements Animal
type Bird struct {
	name string
}

func (bird Bird) Eat() {
	fmt.Println("worms")
}

func (bird Bird) Move() {
	fmt.Println("fly")
}

func (bird Bird) Speak() {
	fmt.Println("peep")
}

// Snake implements Animal
type Snake struct {
	name string
}

func (snake Snake) Eat() {
	fmt.Println("mice")
}

func (snake Snake) Move() {
	fmt.Println("slither")
}

func (snake Snake) Speak() {
	fmt.Println("hsss")
}

func usage() {
	fmt.Println("Usage: newanimal <animal_name> <animal_type>")
	fmt.Println("   or: query     <animal_name> <animal_info>")
}

func Run() {
	m := make(map[string]Animal)

	// prompt mode
	in := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")

		input, err := in.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		args := strings.Fields(input)
		if len(args) != 3 {
			usage()
			continue
		}

		option := strings.ToLower(args[0])
		animalName := strings.ToLower(args[1])
		if strings.Compare(option, "newanimal") == 0 {
			animalType := strings.ToLower(args[2])
			var a Animal
			switch animalType {
			case "cow": a = Cow{animalName}
			case "bird": a = Bird{animalName}
			case "snake": a = Snake{animalName}
			default:
				fmt.Println("Animal type unknown. Try cow, bird or snake")
				continue
			}
			m[animalName] = a
			fmt.Println("Created it!")
		} else if strings.Compare(option, "query") == 0 {
			animal, ok := m[animalName]
			if !ok {
				fmt.Println("Animal name unknown. Have you created it already?")
				continue
			}

			queryName := strings.ToLower(args[2])
			switch queryName {
			case "eat": animal.Eat()
			case "move": animal.Move()
			case "speak": animal.Speak()
			default: fmt.Println("Request info unknown. Try eat, move or speak")
			}
		} else {
			fmt.Println("unknown command")
			usage()
		}
	}
}
