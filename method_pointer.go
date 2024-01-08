package main

import "fmt"

func main() {
	rizal := Man{"Rizal"}
	rizal.Married()
	fmt.Println(rizal.Name)

}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

type Man struct {
	Name string
}