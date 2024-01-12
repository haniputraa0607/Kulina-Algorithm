package question

import (
	"fmt"
	"strings"
)

func QuetionTwo() {
	var index int
	fmt.Println("Question 2")
	fmt.Print("Number of steps Gary takes : ")
	fmt.Scanln(&index)

	if index < 2 || index > 10 {
		fmt.Println("Index cant be lower than 2 and cant be greater than 10")
		return
	}

	var paths = make([]string, index)
	defaultPath := []string{"D", "U"}

	for i := 0; i < index; i++ {
		var path string

		fmt.Printf("Path %d: ", i+1)
		fmt.Scanln(&path)
		paths[i] = strings.ToUpper(path)

		if !pathConstraints(paths[i], defaultPath) {
			fmt.Println("Path must be D or U")
			return
		}
	}

	valleys := countingValleys(paths)
	fmt.Printf("Number of valleys Gary walked through : %d", valleys)

}

func pathConstraints(inputString string, arrayValues []string) bool {
	for _, value := range arrayValues {
		if strings.Contains(inputString, value) {
			return true
		}
	}
	return false
}

func countingValleys(paths []string) int {
	height := 0
	count := 0

	for _, step := range paths {
		if step == "U" {
			height++
		} else if step == "D" {
			height--
			if height == -1 {
				count++
			}
		}
	}

	return count
}