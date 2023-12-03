package main

///
/// For part one there is a better way to do this
/// Hint its similar to how its done to part 2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	// "unicode"
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

	wholeText := make([]string, 0, 10)
	numberGroups := make([][][]int, 0, 10)

	i := 0
	for scanner.Scan() {
		text := scanner.Text()
		wholeText = append(wholeText, text)

		// fmt.Println(unicode.IsDigit([]rune("9")))

		currLine := make([][]int, 0, 10)
		currDigit := make([]int, 0, 10)
		for j := 0; j < len(text); j++ {
			if isDigit(string(text[j])) {
				currDigit = append(currDigit, j)
			} else {
				if len(currDigit) > 0 {
					currLine = append(currLine, currDigit)
					// numberGroups[i] = [][]int{currDigit}
					currDigit = make([]int, 0, 10)
				}
			}

			if j == len(text)-1 {
				currLine = append(currLine, currDigit)
				currDigit = make([]int, 0, 10)
			}

		}
		// fmt.Print("\n")
		numberGroups = append(numberGroups, currLine)
		i++
	}

	total := 0

	for i := 0; i < len(numberGroups); i++ {

		for j := 0; j < len(numberGroups[i]); j++ {

			for k := 0; k < len(numberGroups[i][j]); k++ {

				group := numberGroups[i][j]
				ind := group[k]

				num, err := strconv.Atoi(string(wholeText[i][group[0] : group[len(group)-1]+1]))
				if err != nil {
					continue
				}

				if (ind < len(wholeText[i])-1 && isValid(string(wholeText[i][ind+1]))) || //[i][j+1]
					(ind > 0 && (isValid(string(wholeText[i][ind-1])))) || // [i][j-1]

					(i > 0 && (isValid(string(wholeText[i-1][ind])))) || // [i-1][j]
					(i < len(numberGroups)-1 && (isValid(string(wholeText[i+1][ind])))) || // [i+1][j]

					(i < len(numberGroups)-1 && ind < len(wholeText[i])-1 && (isValid(string(wholeText[i+1][ind+1])))) || //[i+1][j+1]
					(i < len(numberGroups)-1 && ind > 0 && (isValid(string(wholeText[i+1][ind-1])))) || // [i+1][j-1]

					(i > 0 && ind < len(wholeText[i])-1 && (isValid(string(wholeText[i-1][ind+1])))) || //  [i-1][j+1]
					(i > 0 && ind > 0 && (isValid(string(wholeText[i-1][ind-1])))) { // [i-1][j-1]

					// fmt.Println(num)

					total += num
					break
				}

			}
		}

	}

	fmt.Println("Sum of part numbers: ", total)
	// fmt.Println(string(wholeText[0][0]))
}

func isDigit(v string) bool {

	return regexp.MustCompile(`\d`).MatchString(v)

}

func isValid(v string) bool {

	if v == "." || isDigit(v) {
		return false
	}

	return true
}

func partTwo(file *os.File) {
	scanner := bufio.NewScanner(file)

	wholeText := make([]string, 0, 10)
	gearRatios := make([]int, 0)

	starRe := regexp.MustCompile(`\*`)
	digitRe := regexp.MustCompile(`\d+`)

	for scanner.Scan() {
		text := scanner.Text()

		wholeText = append(wholeText, text)
	}

	padding := strings.Repeat(".", len(wholeText[0]))
	padded := append([]string{padding}, wholeText...)

	fmt.Println(padded)

	for i, text := range wholeText {

		gears := starRe.FindAllStringSubmatchIndex(text, -1)

		if len(gears) < 1 {
			continue
		}

		for _, gear := range gears {
			numbers := make([]int, 0)

			for row := i; row < i+3; row++ {

				for _, n := range digitRe.FindAllStringIndex(padded[row], -1) {
					lower, upper := n[0], n[1]
					if lower-1 <= gear[0] && gear[0] <= upper {
						numbers = append(numbers, atoi(padded[row][lower:upper]))
					}
				}
			}

			if len(numbers) == 2 {

				gearRatios = append(gearRatios, numbers[0]*numbers[1])
			}

		}

	}

	sum := 0
	for _, v := range gearRatios {
		sum += v
	}

	// fmt.Println(gearRatios)
	fmt.Println("Sum of gear ratios: ", sum)
}

func atoi(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}
