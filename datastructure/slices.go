package datastructure

import "fmt"

func Modify(c []int) {
	c[0] = 100

}

func Modify2(fn [2]int) {
	fn[0] = 100

}

func SliceOperation() {

	//

	//var c []int //

	c := []int{1, 2} // c stores the address of first location or the zero index element
	Modify(c)
	d := [2]int{1, 2}

	Modify2(d)
	fmt.Println(d)
	fmt.Println(c) // will print ? [1,2] ,[100,2]
}

// Iterating over an slice
// go doesnt have concept of while loop
// it only has for loop

// variadic function

func Variadicexample(c ...int) {
	fmt.Println(c)
}

func ForLoop() {

	//var c [3]bool // it actually creates an array internally
	//c = append(c, 9)
	//c[0] = 1
	//c[1] = 2
	//c[2] = 4

	//c := make([]int, 9)
	//c[0] = 1
	var c []int // this has not created any memory space yet interNALLY INT the system
	//	c = append(c, 9) // actually adds the element inside the slice
	// slice creates memory on demand
	c = append(c, 90, 90, 12, 3, 4, 4, 4, 44)
	//d := []int{90, 89}
	//d[1] = 900
	//c[0] = 99
	//c[0] = 9
	//d := []int{}

	//.Println(d)

	for i := 0; i < len(c); i++ {
		fmt.Println(c[i])
	}

}

func RangeSlice() {

	var c []int // this has not created any memory space yet interNALLY INT the system
	// slice creates memory on demand
	c = append(c, 90, 90, 12, 3, 4, 4, 4, 44)

	for i, value := range c {
		fmt.Println(i, value)

	}

}

func SliceOperations() {

	//var c []int // this slice is nil no memory assignment yet
	//c = append(c, 9)
	//d := []int{}  //append will work
	//d[0] = 1
	d := make([]int, 9)
	d[7] = 99
	fmt.Println(d)

	//	fmt.Println(c)
	//fmt.Println(d == nil)

}

func CopySlices() {

	m := []int{2, 3}
	n := []int{4, 5}

	//copy(m, n)

	fmt.Println(m)

	m = append(m, n...)
	fmt.Println(m)

}
