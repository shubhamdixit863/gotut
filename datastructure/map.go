package datastructure

import "fmt"

// If your function name starts with small character then you cant use it in other packages
// but you can use it in the same package

func PrintMap() {
	//var v map[string]int // its a nil map it just has the descriptor (it stores the memory addres) but we dnt have
	// underlying memory space created
	//v := map[string]int{} // it will create an empty map
	//v["name"] = 89
	//fmt.Println(v["name"])
	m := make(map[string]int) // it also creates an empty map
	m["name"] = 90
	fmt.Println(m)

}

func IterateOverMAp() {

	m := make(map[string]int)
	m["name"] = 9
	m["hello"] = 8

	for _, _ = range m {
		//fmt.Println(k, v)

	}

	//	fmt.Println(len(m))
	//j, ok := m["name"]
	//fmt.Println(j, ok)

	if _, ok := m["hello"]; ok {

		//	fmt.Println("YEs the value exists")

	}

	// you want to delete a particular key and value in golang --->
	delete(m, "hello")
	fmt.Println(m)
}

// key ,value --->

//insert or getting the value is of order O(1) //

// type assertion --
