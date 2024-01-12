package question

import (
	"fmt"
	"math"
	"strconv"
)

func QuetionThree() {

	var number int
	fmt.Println("Question 3")
	fmt.Print("Number : ")
	fmt.Scanln(&number)

	if number == 0 {
		number = 1345679
	}

	pow := int(math.Pow10(len(strconv.Itoa(number)) - 1))

	for pow > 0 {
		digit := number / pow
		fmt.Println(digit * pow)
		number %= pow
		pow /= 10
	}
}