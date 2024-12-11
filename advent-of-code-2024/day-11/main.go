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

	or := make([]int, 0)
	o := make([]int, 0)
	c := 25

	for scanner.Scan() {
		text := scanner.Text()

		for _, v := range strings.Fields(text) {
			or = append(or, toInt(v))
		}
	}

	for i := 0; i < c; i++ {
		for j := 0; j < len(or); j++ {
			va := validate(or[j])

			o = append(o, va...)
		}

		or = o
		o = make([]int, 0)
	}

	fmt.Println("Part one: ", len(or))
}

func partTwo(file *os.File) {
	scanner := bufio.NewScanner(file)
	or := make([]int, 0)
	c := 75
	sc := 0

	for scanner.Scan() {
		text := scanner.Text()

		for _, v := range strings.Fields(text) {
			or = append(or, toInt(v))
		}
	}

	ch := 2
	cun := chunkify(or, ch)
	for _, v := range cun {
		sc += process(c, v)
	}

	fmt.Println("Part two: ", sc)
}

func chunkify(array []int, size int) [][]int {
	if size <= 0 {
		return nil // Return nil if size is invalid
	}

	var chunks [][]int
	for i := 0; i < len(array); i += size {
		end := i + size
		if end > len(array) {
			end = len(array) // Adjust for the final chunk
		}
		chunks = append(chunks, array[i:end])
	}
	return chunks
}

func process(c int, or []int) int {
	r := or
	o := make([]int, 0)

	for i := 0; i < c; i++ {
		for j := 0; j < len(r); j++ {
			va := validate(r[j])

			o = append(o, va...)
		}

		r = o
		o = make([]int, 0)
	}

	return len(r)
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Panic(err)
	}
	return n
}

func validate(s int) []int {
	p := make([]int, 0)

	// rule 1 = if i = 0, i = 1
	if s == 0 {

		p = append(p, 1)

		return p
	}

	// rule 2
	str := fmt.Sprint(s)
	if len(str)%2 == 0 {
		// sp := strings.Split(str, "")
		st := str[0 : len(str)/2]
		et := str[len(str)/2:]

		p = append(p, []int{toInt(st), toInt(et)}...)
		return p
	}

	// rule 3

	// rule 4
	// p[0] = s * 2024
	p = append(p, s*2024)

	return p
}
