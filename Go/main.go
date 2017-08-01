package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	dob   string
}

func (p person) fullname() string {
	return p.last + ", " + p.first
}

func main() {
	p1 := person{"James", "Bond", "05/04/1974"}
	p2 := person{"Joe", "Blow", "07/19/1990"}
	fmt.Println("Name:", p1.fullname(), "\tDOB:", p1.dob)
	fmt.Println("Name:", p2.fullname(), "\tDOB:", p2.dob)
}
