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

type RaceLimits struct {
	time     int
	distance int
}

func partOne(file *os.File) {
	scanner := bufio.NewScanner(file)

	races := make([]RaceLimits, 0)

	for scanner.Scan() {
		text := scanner.Text()

		if strings.HasPrefix(text, "Time:") {
			v := strings.Split(text, "Time:")[1]
			times := strings.Fields(v)

			tempTime := make([]RaceLimits, 0, len(times))

			for _, t := range times {
				tempTime = append(tempTime, RaceLimits{
					time: atoi(t),
				})
			}

			races = tempTime
		}

		if strings.HasPrefix(text, "Distance:") {
			v := strings.Split(text, "Distance:")[1]
			distances := strings.Fields(v)

			for i, t := range distances {
				races[i].distance = atoi(t)
			}

		}
	}

	multiple := 1

	for _, r := range races {

		upper, lower := 0, 0

		for i := 0; i < r.time; i++ {
			if (i * (r.time - i)) > r.distance {
				lower = i
			}

			j := r.time - i
			if (j * i) > r.distance {
				upper = j
			}

			if upper != 0 && lower != 0 {
				break
			}
		}

		// fmt.Println((upper - lower) + 1)

		multiple *= (upper - lower) + 1

	}

	fmt.Println("Multiple: ", multiple)
}
func partTwo(file *os.File) {
	scanner := bufio.NewScanner(file)

	race := RaceLimits{}

	for scanner.Scan() {
		text := scanner.Text()

		if strings.HasPrefix(text, "Time:") {
			v := strings.Split(text, "Time:")[1]
			times := strings.Join(strings.Fields(v), "")

			race.time = atoi(times)

		}

		if strings.HasPrefix(text, "Distance:") {
			v := strings.Split(text, "Distance:")[1]
			distance := strings.Join(strings.Fields(v), "")

			race.distance = atoi(distance)

		}
	}

	multiple := 1

	upper, lower := 0, 0

	for i := 0; i < race.time; i++ {
		if (i * (race.time - i)) > race.distance {
			lower = i
		}

		j := race.time - i
		if (j * i) > race.distance {
			upper = j
		}

		if upper != 0 && lower != 0 {
			break
		}
	}

	// fmt.Println((upper - lower) + 1)
	multiple *= (upper - lower) + 1

	fmt.Println("Multiple: ", multiple)
}

func atoi(s string) int {

	n, err := strconv.Atoi(s)

	if err != nil {
		log.Panic(err)
	}

	return n
}

//
// between i - (D-i)
//  0 * 7 = 0
//  1 *  6 =  1
//  2 * 5 =  10
//  3 * 4 =  12
//  4 * 3 =  12
//  5 * 2 =  10
//  6 * 1 =  6
//  7 * 0 =  6
