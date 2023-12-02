package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var digitMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	var totalCount int

	if len(os.Args[1:]) < 1 {
		log.Fatal("Please input file path")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text := scanner.Text()

		nums := extractWord(text)

		num := nums[0] + nums[len(nums)-1]
		// fmt.Println(text, num)

		v, err := strconv.Atoi(num)
		if err != nil {
			continue
		}

		totalCount += v
	}

	fmt.Println("Sum of the calibrations: ", totalCount)
}

func extractWord(input string) []string {
	var digitKeys []string
	for key := range digitMap {
		digitKeys = append(digitKeys, key)
	}

	shouldRepeat := true

	words := []string{}

	currWord := ""

	for shouldRepeat {
		for i, char := range input {

			if unicode.IsDigit(char) {
				words = append(words, string(char))
				currWord = string(char)
			} else {
				currWord += string(char)

				w, ok := PM(digitKeys, currWord)
				// fmt.Println(currWord, w, ok)
				if ok {
					words = append(words, strconv.Itoa(digitMap[w]))
					currWord = string(char)
				}
			}

			if i == len(input)-1 {
				shouldRepeat = false
			}
		}
	}

	return words

}

func PM(arr []string, target string) (string, bool) {
	for _, element := range arr {
		if strings.HasSuffix(target, element) {
			return element, true
		}
	}
	return target, false
}
