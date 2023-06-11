package datastructure

import "fmt"

// we dont have inheritance  // we have composition
// we have abstraction
// we have polymorphism
// we have encapsulation

// structs are composite data type in golang

type Animal struct {
	Legs     int
	Category string
}

type School struct {
	students int // not exported and you cant access it
	Teachers int
}

//receiver method

func (an Animal) Walk(a string) { // value receiver method

	fmt.Println(an.Legs, an.Category, a)
	an.Legs = 9

}

type Human struct {
	Legs int
	Ears int
}

// pointer receiver method

func (hu *Human) ChangeThEar(ear int) {
	hu.Ears = ear

}

func StructMethods() {

	// how to create an object of struct
	/*
			dog := Animal{ // this is our dog object
				Legs:     4,
				Category: "Mamal",
			}

			//fmt.Println(dog)
			dog.Walk("hello")

			fmt.Println(dog)



		human := &Human{
			Legs: 2,
			Ears: 2,
		}
	*/

	human := new(Human)
	human.Legs = 90
	human.Ears = 8
	fmt.Println(human)

	fmt.Println(human)

	human.ChangeThEar(9)

	fmt.Println(human)

	// new keyword --->

	// cr
	Anonymoustruct()

}

func Anonymoustruct() {

	address := struct {
		City    string
		Pincode string
	}{
		// values of all the fields have that you have defined
		City:    "NewDelhi",
		Pincode: "1899910",
	}

	fmt.Println(address)

	// Anonymous structs

}
