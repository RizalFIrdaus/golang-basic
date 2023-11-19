package main

import "fmt"

func main() {

	names := [...]string{
		"Muhammad",
		"Rizal",
		"Firdaus",
		"Suci",
		"Maydinna",
		"Nur",
	}

	slice1 := names[:3]
	slice2 := names[3:]
	slice3 := names[3:5]
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

	products := [...]string{"Indomie", "Goreng", "Special", "Rasa", "Ayam", "Geprek"}

	slice4 := products[1:3]
	slice5 := products[2:5]
	slice6 := products[4:6]

	fmt.Println(slice4)
	fmt.Println(slice5)
	fmt.Println(slice6)
	fmt.Println(len(slice6))
	fmt.Println(cap(slice5))

	slice7 := append(slice6, "Sedap")

	fmt.Println(slice7)

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "ahad"}
	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu"
	daySlice1[1] = "Ahad"
	fmt.Println(daySlice1)

	daySlice2 := append(daySlice1, "Liburan", "Weekdays")
	fmt.Println(daySlice2)
	fmt.Println(days)

	var newSlice []string = make([]string, 2, 5)

	newSlice[0] = "Google"
	newSlice[1] = "."
	newSlice1 := append(newSlice, "com")
	fmt.Println(newSlice1)

	copy(newSlice, daySlice2)

	iniArray := [...]byte{1, 2, 3, 4, 5}
	iniSlice := []byte{1, 2, 3, 4, 5}

	fmt.Println(daySlice2)

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
