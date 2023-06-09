package datastructure

import "fmt"

// Arrays and slices -->

// Arrays

func Arrays() {
	// first way
	var a [3]int // array of 3 length   //
	//var b[3]int{0,1,2}
	b := [3]int{0, 0, 0}
	fmt.Println(b)
	fmt.Println(a)

	/// slice

	var slc []int // slice  // reference type --->Address type

	fmt.Println(slc) // slice printing

}

func PassPointer(c *int) { // *int is just a representation of pointer
	// if i change the value of c here
	*c = 100 /// de referencing

}

func StringPointer(f **string) { // f is the pointer to  a variable which is storing string value
	//*f = "hii"
	**f = "hello people"
}

func Pointers() {

	h := "HEllo world"
	k := &h
	l := &k
	StringPointer(l)
	fmt.Println(h) // output ?hii

	//c := 1  // 1 would be stored somewhere in the memory
	//f := &c // f will be storing the memory address of c

	//fmt.Println(c) //1
	//PassPointer(f)
	//fmt.Println(c) //100

}
