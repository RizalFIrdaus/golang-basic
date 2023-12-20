package main

import "fmt"

type Customer struct {
	Name, Address, BirthPlace string
	Age                       int
}

func main() {

	rizal := Customer{
		Name:       "Rizal",
		Address:    "Jln Depok",
		BirthPlace: "Banyuwangi",
		Age:        24,
	}

	fmt.Println(rizal.Name)

	rizal.sayHello("Miftah")

}

func (customer Customer) sayHello(name  string)  {
	fmt.Println("Hello ",name,", my name is" , customer.Name)
}