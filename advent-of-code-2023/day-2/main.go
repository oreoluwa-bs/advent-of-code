package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
	fmt.Println("-------------PART TWO---------------")
	partTwo(file)
}

func partOne(file *os.File) {
	maxVal := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	totalCount := 0

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text := scanner.Text()
		gameAndSet := strings.Split(text, ":")

		gameId, err := strconv.Atoi(strings.Split(gameAndSet[0], "Game ")[1])
		if err != nil {
			log.Fatal(err)
		}
		set := gameAndSet[1]

		set = strings.ReplaceAll(set, ";", ",")
		colorList := strings.Split(set, ",")

		isPossible := true

		for _, entry := range colorList {
			pairs := strings.Split(entry, ",")
			for _, pair := range pairs {
				pair = strings.TrimSpace(pair)
				parts := strings.Fields(pair)
				if len(parts) == 2 {
					color := parts[1]
					count, err := strconv.Atoi(parts[0])
					if err == nil {
						if count > maxVal[color] {
							isPossible = false
							break
						}

					}
				}
			}
		}

		if isPossible {
			fmt.Println(gameId)
			totalCount += gameId
		}

	}

	fmt.Println("Sum of ID's: ", totalCount)
}

func partTwo(file *os.File) {
	totalCount := 0

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text := scanner.Text()
		gameAndSet := strings.Split(text, ":")
		set := gameAndSet[1]

		set = strings.ReplaceAll(set, ";", ",")

		colorList := strings.Split(set, ",")

		colorCounts := make(map[string]int)

		for _, entry := range colorList {
			pairs := strings.Split(entry, ",")
			for _, pair := range pairs {
				pair = strings.TrimSpace(pair)
				parts := strings.Fields(pair)
				if len(parts) == 2 {
					color := parts[1]
					count, err := strconv.Atoi(parts[0])
					if err == nil {

						existingCount, found := colorCounts[color]

						if !found || count > existingCount {
							colorCounts[color] = count
						}

					}
				}
			}
		}

		power := 1
		for _, v := range colorCounts {
			power *= v
		}
		fmt.Println(power)
		totalCount += power

	}

	fmt.Println("Sum of powers: ", totalCount)
}
