package day_1

import (
	"fmt"
)

func Day_1_part_1(data string) {

	floorCounter := 0
	for i := 0; i < len(data); i++ {
		if data[i] == '(' {
			floorCounter += 1
		} else {
			floorCounter -= 1
		}
	}
	fmt.Println(floorCounter)
}

func Day_1_part_2(data string) {
	floorCounter := 0
	charPosition := 0
	for i := 0; i < len(data); i++ {
		if data[i] == '(' {
			floorCounter += 1
		} else {
			floorCounter -= 1
		}
		if floorCounter == -1 {
			charPosition = i
			break
		}
	}
	fmt.Println(charPosition + 1)
}
