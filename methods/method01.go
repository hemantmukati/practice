package main

import "fmt"

func main() {

	r := register{
		name:     "hemant mukati",
		subject:  "cse",
		mobile:   123,
		password: "h@123",
	}

	r.registerUser()
	fmt.Println("details of user:", r)
}

type register struct {
	name     string
	subject  string
	mobile   int
	password string
}

func (x register) registerUser() {

	fmt.Println("user name is :", x.name)
	fmt.Println("user mobile no. is :", x.mobile)
}
