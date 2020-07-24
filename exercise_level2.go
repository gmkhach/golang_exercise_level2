package main

import "fmt"

func main() {
	/*
		Exercise 1
		Write a program that prints a number in decimal, binary, and hex
	*/
	a := 42
	fmt.Printf("decimal: %v\tbinary: %b\thexadecimal: %X\n", a, a, a)

	/*
		Exercise 2
		Using the following operators, write expressions and assign their values to variables:
			==
			<=
			>=
			!=
			<
			>
		Now print each of the variables.
	*/
	b := 0.5 == 1/2
	c := 1 <= 5
	d := 5 >= 5
	e := 0.5 != 1/2
	f := -5 < 1
	g := -5 > 1

	fmt.Println(b, c, d, e, f, g)

	/*
		Exercise 3
	  	Create TYPED and UNTYPED constants. Print the values of the constants.
	*/
	const (
		typed   int    = 42
		untyped string = "The meaning of life"
	)

	fmt.Println(typed, untyped)

	/*
		Exercise 4
		Write a program that:
			1. assigns an int to a variable
			2. prints that int in decimal, binary, and hex
			3. shifts the bits of that int over 1 position to the left, and assigns that to a variable
			4. prints that variable in decimal, binary, and hex
	*/
	h := 64
	fmt.Printf("decimal: %v\tbinary: %b\t\thexadecimal: %X\n", h, h, h)
	i := h << 1
	fmt.Printf("decimal: %v\tbinary: %b\thexadecimal: %X\n", i, i, i)

	/*
	   	Exercise 5
		Create a variable of type string using a raw string literal. Print it.
	*/
	j := `"Do, or do not! There is no try." 
			- Yoda`
	fmt.Println(j)

	/*
	   	Exercise 6
		Using iota, create 4 constants for the NEXT 4 years. Print the constant values.
	*/
	const (
		y2021 = 2021 + iota
		y2022 = 2021 + iota
		y2023 = 2021 + iota
		y2024 = 2021 + iota
	)
	fmt.Println(y2021, y2022, y2023, y2024)
}
