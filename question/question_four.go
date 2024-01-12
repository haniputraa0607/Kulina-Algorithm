package question

import "fmt"

func QuetionFour() {

	var trips int
	fmt.Println("Question 4")
	fmt.Print("Trips : ")
	fmt.Scanln(&trips)

	totalLamps := countLamp(trips)

	fmt.Printf("After %d trips, the number of turned-on lamps is: %d\n", trips, totalLamps)
}

func countLamp(trips int) int {
	lamps := make([]bool, 100)

	for i := 1; i <= trips; i++ {
		toggleLamp(lamps, i)
	}

	count := 0
	for _, state := range lamps {
		if state {
			count++
		}
	}

	return count
}
func toggleLamp(lamps []bool, index int) {
	for i := index - 1; i < len(lamps); i += index {
		lamps[i] = !lamps[i]
	}
}
