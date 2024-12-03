package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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
	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		text := scanner.Text()

		// Compile the regex pattern to match "mul(" followed by digits and commas, ending with ")"
		re := regexp.MustCompile(`mul\((\d+(?:,\d+)*)\)`)
		matches := re.FindAllStringSubmatch(text, -1)
		for _, match := range matches {
			if len(match) > 1 {
				// fmt.Println("Extracted numbers:", match[1])
				n := strings.Split(match[1], ",")
				sum += toInt(n[0]) * toInt(n[1])
			}
		}
	}

	fmt.Println("Part one: ", sum)
}

func partTwo(file *os.File) {
	scanner := bufio.NewScanner(file)

	sum := 0
	t := ""

	for scanner.Scan() {
		text := scanner.Text()

		t += text

	}

	// Compile the regex pattern to match "mul(" followed by digits and commas, ending with ")"
	re := regexp.MustCompile(`mul\((\d+(?:,\d+)*)\)`)
	matches := re.FindAllStringIndex(t, -1)

	for _, match := range matches {
		lDont := strings.LastIndex(t[0:match[1]], "don't()")
		lDo := strings.LastIndex(t[0:match[1]], "do()")

		if lDont > lDo {
			continue
		}

		re := regexp.MustCompile(`\d+`)

		m := re.FindAllString(t[match[0]:match[1]], -1)
		sum += toInt(m[0]) * toInt(m[1])
	}

	fmt.Println("Part two: ", sum)
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Panic(err)
	}
	return n
}
