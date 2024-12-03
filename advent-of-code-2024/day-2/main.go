package day2

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

	safeC := 0
	for scanner.Scan() {
		text := scanner.Text()
		f := strings.Fields(text)
		vals := []int{}

		for i := range f {
			vals = append(vals, toInt(f[i]))
		}

		if isSafe(vals) {
			safeC++
		}
	}

	fmt.Println("Part one: ", safeC)
}

func partTwo(file *os.File) {
	scanner := bufio.NewScanner(file)

	safeC := 0
	for scanner.Scan() {
		text := scanner.Text()
		f := strings.Fields(text)

		vals := []int{}

		for i := range f {
			vals = append(vals, toInt(f[i]))
		}

		if isSafe(vals) {
			safeC++
			continue
		}

		for i := 0; i < len(vals); i++ {
			temp := remove(vals, i)
			if isSafe(temp) {
				safeC++
				break
			}
		}
	}

	fmt.Println("Part two: ", safeC)
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Panic(err)
	}
	return n
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func isMonotonous(inp []int) bool {
	increasing, decreasing := true, true

	for i := 1; i < len(inp); i++ {
		if inp[i] < inp[i-1] {
			increasing = false
		}
		if inp[i] > inp[i-1] {
			decreasing = false
		}

		if inp[i] == inp[i-1] {
			return false
		}

	}
	return increasing || decreasing
}

func remove(slice []int, index int) []int {
	temp := append([]int{}, slice[:index]...)
	temp = append(temp, slice[index+1:]...)
	return temp
}

func isSafe(vals []int) bool {
	if !isMonotonous(vals) {
		return false
	}

	for i := 1; i < len(vals); i++ {
		d := abs(vals[i] - vals[i-1])
		if d > 3 || d < 1 {
			return false
		}
	}
	return true
}
