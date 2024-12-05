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

	p1 := partOne(file)
	file.Seek(0, 0) // reset back to 0
	fmt.Println("-------------PART TWO---------------")
	partTwo(file, p1)
}

func partOne(file *os.File) int {
	scanner := bufio.NewScanner(file)

	order := make(map[string][2]int, 0)
	// mid := make([]int, 0)
	mul := 0
	isOrderSection := true
	// n := len(puzzle)

	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			isOrderSection = false
			continue
		}

		if isOrderSection {
			sp := strings.Split(text, "|")
			order[text] = [2]int{toInt(sp[0]), toInt(sp[1])}
			continue
		}

		list := make([]int, 0)
		t := strings.Split(text, ",")
		for _, v := range t {
			list = append(list, toInt(v))
		}

		valid := true
		for i := 0; i < len(list); i++ {
			val := true
			for j := i + 1; j < len(list); j++ {
				_, ok := order[fmt.Sprint(list[i], "|", list[j])]
				if !ok {
					val = false
					break
				}
			}

			if !val {
				valid = false
				break
			}
		}

		if valid {
			// mid = append(mid, list[len(list)/2])
			mul += list[len(list)/2]
		}

	}

	fmt.Println("Part one: ", mul)
	return mul
}

func partTwo(file *os.File, cr int) {
	scanner := bufio.NewScanner(file)

	order := make(map[string][2]int, 0)
	mul := 0
	isOrderSection := true

	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			isOrderSection = false
			continue
		}

		if isOrderSection {
			sp := strings.Split(text, "|")
			order[text] = [2]int{toInt(sp[0]), toInt(sp[1])}
			continue
		}

		list := make([]int, 0)
		t := strings.Split(text, ",")
		for _, v := range t {
			list = append(list, toInt(v))
		}

		valid := true
		for i := 0; i < len(list); i++ {
			val := true
			for j := i + 1; j < len(list); j++ {
				_, ok := order[fmt.Sprint(list[i], "|", list[j])]
				if !ok {
					list[i], list[j] = list[j], list[i]
				}
			}

			if !val {
				valid = false
				break
			}
		}

		if valid {
			mul += list[len(list)/2]
		}

	}
	fmt.Println("Part two: ", mul-cr)
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Panic(err)
	}
	return n
}
