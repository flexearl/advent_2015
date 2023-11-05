package Day_3

import "fmt"

func Part_1(data string) {
	horizontalCounter := 0
	verticalCounter := 0
	indexes := make([][]int, 0)

	multiPresentTracker := 0
	for i := 0; i < len(data); i++ {
		newIndex := [2]int{horizontalCounter, verticalCounter}
		if itemInList(indexes, newIndex) {

		} else {
			multiPresentTracker++
			indexes = append(indexes, newIndex[:])
		}
		if data[i] == '^' {
			verticalCounter++
		} else if data[i] == '>' {
			horizontalCounter++
		} else if data[i] == 'v' {
			verticalCounter--
		} else {
			horizontalCounter--
		}
	}
	fmt.Println(multiPresentTracker)
}

func Part_2(data string) {

	horizontalCounter := 0
	verticalCounter := 0
	roboHorizontalCounter := 0
	roboVerticalCounter := 0
	indexes := make([][]int, 0)

	multiPresentTracker := 0
	for i := 0; i <= len(data); i++ {
		fmt.Println(i)
		newIndex := [2]int{horizontalCounter, verticalCounter}
		newRoboIndex := [2]int{roboHorizontalCounter, roboVerticalCounter}
		fmt.Println(newIndex, newRoboIndex)
		if (itemInList(indexes, newIndex) && i%2 == 1) || (itemInList(indexes, newRoboIndex) && i%2 == 0) {
			fmt.Println("Already in list")
		} else {
			multiPresentTracker++
			if i%2 == 1 {
				indexes = append(indexes, newIndex[:])
			} else {
				indexes = append(indexes, newRoboIndex[:])
			}

		}
		if i == len(data) {
			break
		}
		if i%2 == 0 {
			horizontalCounter, verticalCounter = increaseCounters(string(data[i]), horizontalCounter, verticalCounter)
		} else {
			roboHorizontalCounter, roboVerticalCounter = increaseCounters(string(data[i]), roboHorizontalCounter, roboVerticalCounter)
		}

	}
	fmt.Println(multiPresentTracker)
}

func increaseCounters(data string, horizontalCounter int, verticalCounter int) (int, int) {
	if data == "^" {
		verticalCounter++
	} else if data == ">" {
		horizontalCounter++
	} else if data == "v" {
		verticalCounter--
	} else {
		horizontalCounter--
	}
	return horizontalCounter, verticalCounter
}

func itemInList(list [][]int, items [2]int) bool {
	foundItem := false
	for i := 0; i < len(list); i++ {
		if list[i][0] == items[0] && list[i][1] == items[1] {
			foundItem = true
			break
		}
	}
	return foundItem
}
