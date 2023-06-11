package datastructure

import "fmt"

// composition is a has a relationship

type Address struct {
	Pincode     string
	City        string
	Country     string
	Coordinates            // embedding
	people      Population // named embedding

}

func (ad Address) PrintAddress() {

	fmt.Println("Address")

}

type Population struct {
	PeopleNumber int
}

func (pu Population) PopulationPrint() {
	fmt.Println("Population printed")
}

type Coordinates struct {
	Latitude  int
	Longitude int
}

func (c Coordinates) Printcoordinates() {
	fmt.Println(c)

}

func CompositionInvoke() {

	ad := Address{
		Pincode:     "20200",
		City:        "29900",
		Country:     "India",
		Coordinates: Coordinates{8, 9},
	}

	ad.Printcoordinates()
	ad.PrintAddress()
	ad.people.PopulationPrint()

}
