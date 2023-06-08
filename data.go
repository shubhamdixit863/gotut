package main

import "fmt"

// functions

// definition ///declarion

func Add(age int, name string, data bool) { // parameters

	fmt.Println(age, name, data)

}

/// function that returns a value

func Subtract(a int, b int) int {

	//c:=a-b  // shorthand declartion

	var c int

	c = a - b

	return c
}

/// calculates sum of 100 natural numbers

func Loop() {

	var c int
	for i := 0; i <= 100; i++ {
		c = c + i

	}

	fmt.Println(c)

}

// you can return multiple values from a function in go

func foo(a int, b int) (int, float64, bool) {

	c := a + b      // int
	f := float64(c) // float64
	boo := true

	return c, f, boo

}

// Lazy way of defining parameters

func LazyParams(a, b, c int) (int, int, int) {
	return a, b, c
}
