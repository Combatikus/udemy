package main

import "fmt"

type person struct {
	First string
	Last  string
	Age   int
}

type doubleZero struct {
	person
	First         string
	LicenseToKill bool
}

func main() {
	bond := doubleZero{
		person: person{
			First: "James",
			Last:  "Bond",
			Age:   20,
		},
		First:         "Double Zero Seven",
		LicenseToKill: true,
	}

	mp := doubleZero{
		person: person{
			First: "Miss",
			Last:  "Money Penny",
			Age:   19,
		},
		First:         "If looks could kill",
		LicenseToKill: false,
	}

	fmt.Println(bond.First, bond.person.First)
	fmt.Println(mp.First, mp.person.First)
}
