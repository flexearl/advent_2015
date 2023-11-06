package Day_6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Part_1(data string) {
	grid := [1000][1000]bool{}

	f, err := os.Open("Day_6/Day_6.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		splitted := strings.Split(text, " ")
		startIndex := getIndex(splitted[len(splitted)-3])
		endIndex := getIndex(splitted[len(splitted)-1])
		if splitted[0] == "turn" {
			turnOn := false
			if splitted[1] == "on" {
				turnOn = true
			}
			for i := startIndex[0]; i <= endIndex[0]; i++ {
				for j := startIndex[1]; j <= endIndex[1]; j++ {
					grid[i][j] = turnOn
				}
			}
		} else {
			for i := startIndex[0]; i <= endIndex[0]; i++ {
				for j := startIndex[1]; j <= endIndex[1]; j++ {
					grid[i][j] = !grid[i][j]
				}
			}
		}
	}
	litCounter := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if grid[i][j] {
				litCounter++
			}
		}
	}
	fmt.Println(litCounter)
}

func Part_2(data string) {
	grid := [1000][1000]int{}

	f, err := os.Open("Day_6/Day_6.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		splitted := strings.Split(text, " ")
		startIndex := getIndex(splitted[len(splitted)-3])
		endIndex := getIndex(splitted[len(splitted)-1])
		if splitted[0] == "turn" {
			turnOn := 1
			if splitted[1] != "on" {
				turnOn *= -1
			}
			for i := startIndex[0]; i <= endIndex[0]; i++ {
				for j := startIndex[1]; j <= endIndex[1]; j++ {
					if grid[i][j] == 0 && turnOn < 0 {

					} else {
						grid[i][j] += turnOn
					}

				}
			}
		} else {
			for i := startIndex[0]; i <= endIndex[0]; i++ {
				for j := startIndex[1]; j <= endIndex[1]; j++ {
					grid[i][j] += 2
				}
			}
		}
	}
	litCounter := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			litCounter += grid[i][j]

		}
	}
	fmt.Println(litCounter)
}

func getNumber(numb string) int {
	numb = strings.Replace(numb, ",", "", 0)
	converted, err := strconv.Atoi(numb)
	if err != nil {
		log.Fatal(err)
	}
	return converted
}

func getIndex(numb string) [2]int {
	splitted := strings.Split(numb, ",")
	firstNumber, err := strconv.Atoi(splitted[0])
	secondNumber, err := strconv.Atoi(splitted[1])
	if err != nil {
		log.Fatal(err)
	}
	numbers := [2]int{firstNumber, secondNumber}
	return numbers
}
