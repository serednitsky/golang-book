package main

import "fmt"

func main() {
	//for i := 1; i <= 10; i++ {
	//	if i%2 == 0 {
	//		fmt.Println(i, "even")
	//	} else {
	//		fmt.Println(i, "odd")
	//	}
	//}

	// Task1

	for i := 1; i <= 10; i++ {
		if i > 10 {
			fmt.Println("Big")
		} else {
			fmt.Println("Small")
		}
	}

	// Task2

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i, " divided by 3")
		}
	}

	// Task3

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, ": FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println(i, ": Fizz")
		} else if i%5 == 0 {
			fmt.Println(i, ": Buzz")
		}
	}
}
