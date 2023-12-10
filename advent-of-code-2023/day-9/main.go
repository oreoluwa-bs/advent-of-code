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
	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		text := scanner.Text()

		reading := make([]int, 0)

		for _, v := range strings.Fields(text) {
			reading = append(reading, atoi(v))
		}

		readings := [][]int{reading}

		isReadingZero := false

		iter := 0
		// fmt.Println(readings)
		for !isReadingZero {
			curr := readings[iter]
			next := make([]int, 0)

			for i := 0; i < len(curr)-1; i++ {
				diff := curr[i+1] - curr[i]
				next = append(next, diff)
			}

			// fmt.Println(next)
			readings = append(readings, next)

			z := true
			for _, v := range next {
				if v != 0 {
					z = false
					break
				}
			}

			if len(next) == 0 || z {
				nextNum := 0
				for i := len(readings) - 2; i >= 0; i-- {
					nextNum += readings[i][len(readings[i])-1]
				}

				total += nextNum

				isReadingZero = true
				break
			}

			iter++
		}

		// fmt.Println(readings)

	}

	fmt.Println("Part One:", total)

}

func partTwo(file *os.File) {
	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		text := scanner.Text()

		reading := make([]int, 0)

		for _, v := range strings.Fields(text) {
			reading = append(reading, atoi(v))
		}

		readings := [][]int{reading}

		isReadingZero := false

		iter := 0
		// fmt.Println(readings)
		for !isReadingZero {
			curr := readings[iter]
			next := make([]int, 0)

			for i := 0; i < len(curr)-1; i++ {
				diff := curr[i+1] - curr[i]
				next = append(next, diff)
			}

			// fmt.Println(next)
			readings = append(readings, next)

			z := true
			for _, v := range next {
				if v != 0 {
					z = false
					break
				}
			}

			if len(next) == 0 || z {
				nextNum := 0
				for i := len(readings) - 2; i >= 0; i-- {
					nextNum = readings[i][0] - nextNum
				}

				total += nextNum

				isReadingZero = true
				break
			}

			iter++
		}

		// fmt.Println(readings)

	}

	fmt.Println("Part Two:", total)
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)

	if err != nil {
		log.Panic(err)
	}

	return n
}
