package main

import "fmt"

func main() {
	fmt.Println("hello")

	var number1 uint = 22
	var number2 uint = 334
	var sum_of_1and2 = number1 + number2
	fmt.Println(sum_of_1and2)

	// you can declare vars like this too
	name := "chakravrthi"
	age := "33"
	//or
	var name2 = "hello"
	fmt.Println(name2 + "my name is " + name + " and my age is " + age + "")

	// u need type formating in here like c lang and using var for clarity
	var num1 uint = 22
	var num2 = 35.7

	sum := uint(num2) + num1

	fmt.Println(sum)
}
