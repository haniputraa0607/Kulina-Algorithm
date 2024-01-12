package question

import "fmt"

func QuetionOne() {
	var index int
	fmt.Println("Question 1")
	fmt.Print("Number of socks: ")
	fmt.Scanln(&index)

	if index < 0 || index > 100 {
		fmt.Println("Index cant be lower than 0 and cant be greater than 100")
		return
	}

	var socks = make([]int, index)

	for i := 0; i < index; i++ {
		fmt.Printf("Enter color of sock %d: ", i+1)
		fmt.Scanln(&socks[i])

		if socks[i] < 0 || socks[i] > 100 {
			fmt.Println("Color of sock cant be lower than 0 and cant be greater than 100")
			return
		}
	}

	matching := matchingPairs(socks)
	fmt.Printf("There are %d matching socks", matching)
}

func matchingPairs(arr []int) int {
	countMap := make(map[int]int)
	totalDuplicates := 0

	for _, value := range arr {
		countMap[value]++

		if countMap[value] == 2 {
			totalDuplicates++
		}
	}

	return totalDuplicates
}