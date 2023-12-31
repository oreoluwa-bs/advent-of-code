package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	if len(os.Args[1:]) < 1 {
		log.Fatal("Please input file path")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	partOne(file)
	file.Seek(0, 0) // reset back to 0
	// fmt.Println("-------------PART TWO---------------")
	partTwo(file)
}

func partOne(file *os.File) {
	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		text := scanner.Text()

		// split by  : to get cards
		// split by | to get card groups
		cards := strings.Split(strings.Split(text, ":")[1], "|")
		winning, mycards := strings.Fields(cards[0]), strings.Fields(cards[1])

		point := 0

		for _, winNum := range winning {
			if slices.Contains(mycards, winNum) {
				if point == 0 {
					point = 1
				} else {

					point *= 2
				}
			}
		}

		// fmt.Println("Points - ", strings.Split(text, ":")[0], " : ", point)
		total += point
	}

	fmt.Println("Sum of points is: ", total)
}

func partTwo(file *os.File) {
	scanner := bufio.NewScanner(file)

	total := 0
	cardList := make([]string, 0)
	multiplicity := make([]int, 0)

	for scanner.Scan() {
		text := scanner.Text()

		cardList = append(cardList, strings.Split(text, ":")[1])
		multiplicity = append(multiplicity, 1)
	}
	// fmt.Println(multiplicity)

	for i, card := range cardList {
		winCount := calcCardsWon(card)
		if winCount > 0 {
			for j := 1; j <= winCount; j++ {
				multiplicity[i+j] += multiplicity[i]
			}
		}
	}

	for _, count := range multiplicity {
		total += count
	}

	fmt.Println("Sum of points is: ", total)
}

func calcCardsWon(card string) int {
	// split by  : to get cards
	// split by | to get card groups
	cards := strings.Split(card, "|")
	winning, mycards := strings.Fields(cards[0]), strings.Fields(cards[1])

	count := 0
	for _, winNum := range winning {
		if slices.Contains(mycards, winNum) {
			count++
		}
	}

	return count
}
