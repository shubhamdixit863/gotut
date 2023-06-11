package datastructure

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Salary string `json:"salary"`
}

func UnMarshallingJSon() {

	h := []byte(`{
   "id": 23,
   "name": "Jimmie Gillespie",
   "salary": "$20488"
 }`)

	pe := &Person{}

	err := json.Unmarshal(h, pe)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(pe.Name)

}

func Marshallingjson() {

	// We are gonna be converting the struct to the json data type

	pe := Person{
		Id:     1,
		Name:   "Shubham",
		Salary: "2000000",
	}
	// we will convert to json

	marshal, err := json.Marshal(&pe)
	if err != nil {
		return
	}

	fmt.Println(string(marshal))

	// Reference types - error ,pointer ,slices ,maps interface
	// value types --struct ,arrays ,primitive data types all are value types

}

func Comparison() {

	//h := [2]int{1, 20}
	//f := [2]int{1, 2}
	//fmt.Println(h == f)

	p1 := Person{
		Id:     0,
		Name:   "",
		Salary: "",
	}

	p2 := Person{
		Id:     0,
		Name:   "",
		Salary: "",
	}
	//fmt.Println(p1, p2)

	fmt.Println(p1 == p2)
}

/**
{
   "id": 23,
   "name": "Jimmie Gillespie",
   "salary": "$20488"
 },
*/
