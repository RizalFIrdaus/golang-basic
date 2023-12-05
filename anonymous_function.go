package main

import "fmt"

func main() {

	//Anonymous function
	//bl := func(username string) bool {
	//	lists := [...]string{
	//		"scammer",
	//		"tukang_cheat",
	//		"rizal300500",
	//	}
	//	found := false
	//
	//	for _, list := range lists {
	//		if list == username {
	//			found = true
	//			break
	//		}
	//	}
	//
	//	return found
	//}

	registerUser("scammer", func(username string) bool {
		lists := [...]string{
			"scammer",
			"tukang_cheat",
			"rizal300500",
		}
		found := false

		for _, list := range lists {
			if list == username {
				found = true
				break
			}
		}

		return found
	})
}

type BlackList func(username string) bool

func registerUser(username string, blacklist BlackList) {

	if blacklist(username) {
		fmt.Println("You cannot register, because your username is listed to blaclist")
	} else {
		fmt.Println("You succesfully create account")
	}
}

//func blackList(username string) bool {
//
//	lists := [...]string{
//		"scammer",
//		"tukang_cheat",
//		"rizal300500",
//	}
//	found := false
//
//	for _, list := range lists {
//		if list == username {
//			found = true
//			break
//		}
//	}
//
//	return found
//
//}
