package main

import "fmt"

func main() {
	const x string = "Hello World"
	//x = "Some other string"
	fmt.Println(x)

	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)

	//Task: fahrenheit to celsius
	fmt.Print("Enter a fahrenheit: ")
	var fahrenheit float64
	fmt.Scanf("%f", &fahrenheit)

	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Println("Degree celsius = ", celsius)

	//Task: feet in meters
	fmt.Print("Enter a feet: ")
	var feet float64
	fmt.Scanf("%f", &feet)

	const metre = 0.3048

	fmt.Println(feet, "Feet = ", feet*metre, " meters")
}
