package main

import "fmt"

/*
 * Let us assume the following formula for displacement s as a function of
 * time t, acceleration a, initial velocity v_o, and initial displacement s_o.
 * s = ½a*t^2 + v_o*t + s_o
 *
 * Write a program which first prompts the user to enter values for
 * acceleration, initial velocity, and initial displacement.
 * Then the program should prompt the user to enter a value for time
 * and the program should compute the displacement after the entered time.
 *
 * You will need to define and use a function called GenDisplaceFn()
 * which takes three float64 arguments, acceleration a, initial velocity vo, and initial displacement so.
 * GenDisplaceFn() should return a function which computes displacement as a function of time,
 * assuming the given values acceleration, initial velocity, and initial displacement.
 * The function returned by GenDisplaceFn() should take one float64 argument t, representing time,
 * and return one float64 argument which is the displacement travelled after time t.
 *
 * For example, let’s say that I want to assume the following values for
 * acceleration, initial velocity, and initial displacement: a = 10, vo = 2, so = 1.
 * I can use the following statement to call GenDisplaceFn() to generate a function fn
 * which will compute displacement as a function of time.
 * fn := GenDisplaceFn(10, 2, 1)
 * Then I can use the following statement to print the displacement after 3 seconds.
 * fmt.Println(fn(3))
 * And I can use the following statement to print the displacement after 5 seconds.
 * fmt.Println(fn(5))
 */

func GenDisplaceFn(accl, v0, s0 float64) func (float64) float64 {
	return func (t float64) float64 {
		return t*t*accl*0.5 + t*v0 + s0
	}
}
func main() {
	args := make([]float64, 4)

	for i := 0; i < 4; i++ {
		var numFloat float64
		switch i {
		case 0: fmt.Print("input acceleration: ")
		case 1: fmt.Print("input initial velocity: ")
		case 2: fmt.Print("input initial displacement: ")
		case 3: fmt.Print("input value of time: ")
		}

		_, err := fmt.Scan(&numFloat)
		if err != nil {
			fmt.Println("please input a valid floating point number!")
			return
		}
		args[i] = numFloat
	}

	fn := GenDisplaceFn(args[0], args[1], args[2])
	fmt.Println("displacement at t:", fn(args[3]))
}
