package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
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

	s, d, sum := make([]int, 0), make([]int, 0), 0

	for scanner.Scan() {
		text := scanner.Text()
		f := strings.Fields(text)
		s = append(s, toNumber(f[0]))
		d = append(d, toNumber(f[1]))
	}

	sort.Ints(s)
	sort.Ints(d)

	for i := 0; i < len(s); i++ {
		sum += int(math.Abs(float64(d[i] - s[i])))
	}

	fmt.Println("Part one: ", sum)
}

func partTwo(file *os.File) {
	scanner := bufio.NewScanner(file)

	c, l, sum := make(map[int]int), make([]int, 0), 0

	for scanner.Scan() {
		text := scanner.Text()
		f := strings.Fields(text)
		c[toNumber(f[0])] = 0
		l = append(l, toNumber(f[0]))
	}

	file.Seek(0, 0)

	scanner = bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		f := strings.Fields(text)

		_, ok := c[toNumber(f[1])]
		if !ok {
			continue
		}
		c[toNumber(f[1])]++

	}

	for _, v := range l {
		sum += v * c[v]
	}

	fmt.Println("Part two: ", sum)
}

func toNumber(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Panic(err)
	}
	return n
}
