package main

import "fmt"

func main() {
	const isPassedValue = 78

	absent := [20]int{
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	}

	uas := 70
	uts := 80

	finalScore := calculateFinalScoreSwitch(absent, uts, uas)

	switch {
	case finalScore == 100:
		fmt.Println("Nilai kamu sempurna, nilai kamu :", finalScore)
	case finalScore >= 90 && finalScore < 100:
		fmt.Println("Nilai kamu hampir sempurna, nilai kamu :", finalScore)
	case finalScore >= isPassedValue && finalScore < 90:
		fmt.Println("Nilai kamu cukup baik, nilai kamu :", finalScore)
	default:
		fmt.Println("Kamu perlu ujian ulang agar kamu bisa lulus, nilai kamu :", finalScore)
	}

}

func calculateFinalScoreSwitch(absent [20]int, uts int, uas int) int {

	scoreAbsent := countAbsentScoreSwitch(absent) * 5

	return (scoreAbsent + uts + uas) / 3
}

func countAbsentScoreSwitch(absent [20]int) int {
	sum := 0

	for _, value := range absent {
		sum += value
	}
	return sum
}
