package main

import (
	"fmt"
)

func main() {

	isPassedValue := 78

	absent := []int{
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	}

	scoreAbsent := countAbsentScore(absent) * 5

	uts := 76
	uas := 90

	finalScore := (scoreAbsent + uts + uas) / 3

	if finalScore == 100 {
		fmt.Println("Nilai kamu sempurna, nilai kamu :", finalScore)
	} else if finalScore >= 90 && finalScore < 100 {
		fmt.Println("Nilai kamu hampir sempurna, nilai kamu :", finalScore)
	} else if finalScore >= isPassedValue && finalScore < 90 {
		fmt.Println("Nilai kamu cukup baik, nilai kamu :", finalScore)
	} else {
		fmt.Println("Kamu perlu ujian ulang agar kamu bisa lulus, nilai kamu :", finalScore)
	}

}

func countAbsentScore(absent []int) int {
	sum := 0

	for _, value := range absent {
		sum += value
	}
	return sum
}
