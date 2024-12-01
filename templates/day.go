package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	var lT string

	for scanner.Scan() {
		text := scanner.Text()
		lT = text
	}

	fmt.Println("Part one: ", lT)
}

func partTwo(file *os.File) {
	scanner := bufio.NewScanner(file)

	var lT string

	for scanner.Scan() {
		text := scanner.Text()
		lT = text
	}

	fmt.Println("Part two: ", lT)
}
