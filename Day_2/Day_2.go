package Day_2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Part_1(data string) {
	f, err := os.Open("Day_2/Day_2.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()
	scanner := bufio.NewScanner(f)
	total := 0
	for scanner.Scan() {
		subTotal := 0

		strings := strings.Split(scanner.Text(), "x")
		numbs := make([]int, len(strings))
		for i, s := range strings {
			numbs[i], _ = strconv.Atoi(s)
		}

		subTotal += numbs[0] * numbs[1] * 2
		subTotal += numbs[0] * numbs[2] * 2
		subTotal += numbs[1] * numbs[2] * 2

		smallest, secondSmallest := getSmallest(numbs)
		subTotal += smallest * secondSmallest
		total += subTotal
	}
	fmt.Println(total)
}

func Part_2(data string) {
	f, err := os.Open("Day_2/Day_2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	total := 0
	for scanner.Scan() {
		subTotal := 0

		strings := strings.Split(scanner.Text(), "x")
		numbs := make([]int, len(strings))
		for i, s := range strings {
			numbs[i], _ = strconv.Atoi(s)
		}
		smallest, secondSmallest := getSmallest(numbs)
		subTotal += (smallest * 2) + (secondSmallest * 2)
		subTotal += numbs[0] * numbs[1] * numbs[2]
		total += subTotal
	}
	fmt.Println(total)
}

func getSmallest(numbs []int) (int, int) {
	smallest := -1
	secondSmallest := -1

	for i := 0; i < len(numbs); i++ {
		if numbs[i] < smallest || smallest == -1 {
			secondSmallest = smallest
			smallest = numbs[i]

		} else if numbs[i] < secondSmallest || secondSmallest == -1 {
			secondSmallest = numbs[i]
		}
	}
	return smallest, secondSmallest
}
