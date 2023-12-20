package main

import "fmt"

func logging() {
	fmt.Println("Selesai logging")
	message := recover()
	fmt.Println("Terjadi Error : ", message)
}

func runApplication(error bool){
	defer logging()

	if error {
		panic("Something error with database")
	}

	fmt.Println("App successfull running")
}

func main (){
	runApplication(true)
}