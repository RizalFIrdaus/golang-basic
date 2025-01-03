package main

import (
	"fmt"
)

func main() {
	const isPassedValue = 78

	absent := [20]int{
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	}

	uts := 76
	uas := 90

	finalScore := calculateFinalScore(absent, uts, uas)

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

func calculateFinalScore(absent [20]int, uts int, uas int) int {

	scoreAbsent := countAbsentScore(absent) * 5

	return (scoreAbsent + uts + uas) / 3
}

func countAbsentScore(absent [20]int) int {
	sum := 0

	for _, value := range absent {
		sum += value
	}
	return sum
}
