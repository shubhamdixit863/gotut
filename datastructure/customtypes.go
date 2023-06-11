package datastructure

import "fmt"

type Log struct {
	Name string
	Date string
}

type Logs []Log

func (lg Logs) Formatstring() string {
	s := ""

	for _, va := range lg {

		s += fmt.Sprintf("Date %s,Name %s", va.Date, va.Name)
	}

	return s

}

type Money float64 // My money is my own custom type in go

func (mn Money) FormatMoney() string {
	return fmt.Sprintf("$%f", mn)

}

func CallMoney() {

	var money Money
	money = Money(9)
	//fmt.Println(money)
	fmt.Println(money.FormatMoney())

	slc := []Log{
		{
			Name: "Something",
			Date: "Something",
		},
	}

	lg := Logs(slc)
	fmt.Println(lg.Formatstring())

}
