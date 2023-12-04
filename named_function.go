package main

import "fmt"

func main() {

	a, _, c := sayHelloGirl()
	fmt.Println(a)
	fmt.Println(c)
}

func sayHelloGirl() (firstName, middleName, lastName string) {

	firstName = "Muhammad"
	middleName = "Rizal"
	lastName = "Firdaus"

	return firstName, middleName, lastName
}
