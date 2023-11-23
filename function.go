package main

import "fmt"

func main() {

	sayHello()
	helloBro("Rizal", "Firdaus")
	fmt.Println(getHello("Rizal"))
}

func sayHello() {
	fmt.Println("Hello World")

}

func helloBro(firstName string, lastName string) {
	fmt.Println("Hello brok,", firstName, lastName)
}

func getHello(name string) string {
	return "Hello " + name
}
