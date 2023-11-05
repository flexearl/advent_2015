package Day_5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Part_1(data string) {
	f, err := os.Open("Day_5/Day_5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	niceTotal := 0

	for scanner.Scan() {
		text := scanner.Text()
		if hasThreeVowels(text) && twiceInRow(text) && notContain(text) {
			niceTotal++
		} else {
			fmt.Println(text)
		}
	}

	fmt.Println(niceTotal)
}

func Part_2(data string) {
	f, err := os.Open("Day_5/Day_5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	niceTotal := 0
	for scanner.Scan() {
		text := scanner.Text()
		if hasRepeatingMiddle(text) && hasRepeatingPair(text) {
			niceTotal++
		}
	}
	fmt.Println(niceTotal)
}

func hasRepeatingPair(text string) bool {
	hasRepeatingPair := false
	for i := 0; i < len(text)-1 && !hasRepeatingPair; i++ {
		pair := string(text[i]) + string(text[i+1])
		for j := i + 2; j < len(text)-1 && !hasRepeatingPair; j++ {
			fmt.Println(pair, (string(text[j]) + string(text[j+1])))
			hasRepeatingPair = pair == string(text[j])+string(text[j+1])
		}
	}
	return hasRepeatingPair

}

func hasRepeatingMiddle(text string) bool {
	hasRepeating := false
	for i := 0; i < len(text)-2 && !hasRepeating; i++ {
		hasRepeating = text[i] == text[i+2]
	}
	return hasRepeating
}

func hasThreeVowels(text string) bool {
	vowelCounter := 0
	vowelCounter += strings.Count(text, "a")
	vowelCounter += strings.Count(text, "e")
	vowelCounter += strings.Count(text, "i")
	vowelCounter += strings.Count(text, "o")
	vowelCounter += strings.Count(text, "u")

	return vowelCounter > 2
}

func twiceInRow(text string) bool {
	isInRow := false
	for i := 0; i < len(text)-1 && !isInRow; i++ {
		isInRow = text[i] == text[i+1]
	}
	return isInRow
}

func notContain(text string) bool {
	return !strings.Contains(text, "ab") && !strings.Contains(text, "cd") && !strings.Contains(text, "pq") && !strings.Contains(text, "xy")
}
