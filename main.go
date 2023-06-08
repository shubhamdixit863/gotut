package main

import "fmt"

func main() {

	// variable declaration
	// there are three ways to declarae a variable
	// var way
	// shorthand declaration
	// const declaration

	// var way

	//var data int

	// short hand way

	//u := 89  // automatic type inference in go actually works

	//const name = 90   // you have to give the value while creating a variable its necessary

	//name=99

	//fmt.Println(reflect.TypeOf(u))

	// different data types in go --->
	// built in types in go --->
	// int ,float ,int64 ,int32 ,float64 ,float32 -->
	// string ---->text variables
	// boolean

	// 900000000000

	// ---zero value concept

	var c string
	fmt.Println(c) // empty string getting printed

	//name:="shubham"  // string here

	// starting point of all the code in your app

	Add(9, "john", false) // function call
	var result int
	f := Subtract(9, 7)
	fmt.Println(f)

	result = Subtract(90, 70)
	fmt.Println(result)
	fmt.Println(foo(9, 8))

}
