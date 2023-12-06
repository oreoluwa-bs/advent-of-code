package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
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

type Category struct {
	title  string
	values [][]int
}

func partOne(file *os.File) {
	scanner := bufio.NewScanner(file)
	startT := time.Now()

	seeds := make([]int, 0)
	categories := make([]Category, 0)

	for scanner.Scan() {
		text := scanner.Text()

		// Get seeds
		if strings.HasPrefix(text, "seeds:") {
			seedsList := strings.Split(text, ":")[1]
			temp_seeds := make([]int, 0)

			for _, v := range strings.Fields(seedsList) {
				temp_seeds = append(temp_seeds, atoi(v))
			}

			seeds = temp_seeds

		} else {
		}

		// Get category
		if strings.HasSuffix(text, "map:") {
			title := strings.Split(text, "map:")[0]

			categories = append(categories, Category{title: title})
		}

		if len(categories) > 0 {

			fields := strings.Fields(text)
			if len(fields) == 3 {
				numFields := make([]int, 0, 3)

				for _, v := range fields {
					numFields = append(numFields, atoi(v))
				}

				categories[len(categories)-1].values = append(categories[len(categories)-1].values, numFields)
			}
		}
	}

	fmt.Println("Seeds", seeds)

	seedsLoc := make([]int, 0)

	for _, seed := range seeds {

		// is within range of source and length
		nextV := seed

		for _, cat := range categories {

			for _, value := range cat.values {
				dest, source, length := value[0], value[1], value[2]

				if nextV >= source && nextV < source+length {
					// check diff between source and seeed  add to destination - get next cat
					diff := nextV - source
					nextV = dest + diff
					break
				}
			}
		}

		// fmt.Println(nextV)
		seedsLoc = append(seedsLoc, nextV)

	}

	min := seedsLoc[0]
	for _, loc := range seedsLoc {
		if loc < min {
			min = loc
		}
	}

	fmt.Println("Lowest location number: ", min)
	fmt.Println("End Time: ", time.Since(startT))
}

func partTwo(file *os.File) {
	scanner := bufio.NewScanner(file)
	startT := time.Now()

	seeds := make([]int, 0)
	categories := make([]Category, 0)

	for scanner.Scan() {
		text := scanner.Text()

		// Get seeds
		if strings.HasPrefix(text, "seeds:") {
			seedsList := strings.Split(text, ":")[1]
			temp_seeds := make([]int, 0)

			for _, v := range strings.Fields(seedsList) {
				temp_seeds = append(temp_seeds, atoi(v))
			}

			seeds = temp_seeds

		}

		// Get category
		if strings.HasSuffix(text, "map:") {
			title := strings.Split(text, "map:")[0]

			categories = append(categories, Category{title: title})
		}

		if len(categories) > 0 {

			fields := strings.Fields(text)
			if len(fields) == 3 {
				numFields := make([]int, 0, 3)

				for _, v := range fields {
					numFields = append(numFields, atoi(v))
				}

				categories[len(categories)-1].values = append(categories[len(categories)-1].values, numFields)
			}
		}
	}

	fmt.Println("Seeds", seeds)

	lows := math.MaxInt

	for i := 0; i < len(seeds); i += 2 {

		start, rnge := seeds[i], seeds[i+1]
		fmt.Println(start, rnge)

		for j := start; j < start+rnge-1; j++ {
			nextV := j

			for _, cat := range categories {

				for _, value := range cat.values {
					dest, source, length := value[0], value[1], value[2]

					if nextV >= source && nextV < source+length {
						// check diff between source and seeed  add to destination - get next cat
						diff := nextV - source
						nextV = dest + diff
						break
					}
				}
			}
			// seedsLoc = append(seedsLoc, nextV)
			if nextV <= lows {
				lows = nextV
			}
		}
	}

	fmt.Println("Lowest location number: ", lows)
	fmt.Println("End Time: ", time.Since(startT))
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Panic(err)
	}

	return n
}

// For each seed
//     - is within range of source and length
//         - if not - default to value
//         - if it is
//                 - check diff between source and seeed  add to destination - get next cat
