package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

	universe := make([]string, 0)

	// count := 0
	for scanner.Scan() {
		text := scanner.Text()

		// Replace # with number
		// FOR TEST PURPOSES
		// result := ""
		// for _, char := range text {
		// 	if char == '#' {
		// 		count++
		// 		result += fmt.Sprintf("%d", (count))
		// 		// 1 + 1 * 1
		// 	} else {
		// 		result += string(char)
		// 	}
		// }
		universe = append(universe, text)
	}

	// expandedUniverse := expandUniverse(universe)

	// for _, v := range expandedUniverse {
	// 	fmt.Println(v)
	// }

	galaxies := getGalaxies(universe, 2)

	total := 0

	// fmt.Println(galaxies)

	for i := 0; i < len(galaxies)-1; i++ {

		for j := i + 1; j < len(galaxies); j++ {
			dist := manhattanDist(galaxies[i][0], galaxies[i][1], galaxies[j][0], galaxies[j][1])

			// fmt.Println(i+1, "distance to", j+1, "is", dist)

			total += int(dist)

		}
	}

	fmt.Println("Part one: ", total)
}

func partTwo(file *os.File) {
	scanner := bufio.NewScanner(file)

	universe := make([]string, 0)

	// count := 0
	for scanner.Scan() {
		text := scanner.Text()
		universe = append(universe, text)
	}

	galaxies := getGalaxies(universe, 1000000)

	total := 0

	// fmt.Println(galaxies)

	for i := 0; i < len(galaxies)-1; i++ {

		for j := i + 1; j < len(galaxies); j++ {
			dist := manhattanDist(galaxies[i][0], galaxies[i][1], galaxies[j][0], galaxies[j][1])

			// fmt.Println(i+1, "distance to", j+1, "is", dist)

			total += int(dist)

		}
	}

	fmt.Println("Part two: ", total)
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)

	if err != nil {
		log.Panic(err)
	}

	return n
}

func expandUniverse(universe []string) []string {
	expandedUniverseL := make([]string, 0)

	for _, v := range universe {
		if len(v) == strings.Count(v, ".") {
			expandedUniverseL = append(expandedUniverseL, v)
		}
		expandedUniverseL = append(expandedUniverseL, v)
	}

	expandedUniverse := make([]string, len(expandedUniverseL))

	// fmt.Println(expandedUniverseL)
	// for _, v := range expandedUniverseL {
	// 	fmt.Println(v)
	// }

	for i := 0; i < len(expandedUniverseL[0]); i++ {
		shouldExpand := true

		for j := 0; j < len(expandedUniverseL); j++ {
			// fmt.Println(string(expandedUniverseL[j][i]))
			if expandedUniverseL[j][i] != '.' {
				shouldExpand = false
			}
			expandedUniverse[j] += string(expandedUniverseL[j][i])

		}

		if shouldExpand {
			for j := 0; j < len(expandedUniverseL); j++ {
				expandedUniverse[j] += "."
			}

		}

	}

	return expandedUniverse

}

func getGalaxies(universe []string, space int) [][]int {

	galaxies := make([][]int, 0)

	expandY := 0
	for i := 0; i < len(universe[0]); i++ {
		expandX := 0
		shouldExpandY := true

		for j, v := range universe {
			if len(v) == strings.Count(v, ".") {

				// -1 to replace instead of just add
				expandX += (space - 1)
			}

			if v[i] != '.' {
				shouldExpandY = false
				galaxies = append(galaxies, []int{j + expandX, i + expandY})
				// expandX = 0
			}
		}
		if shouldExpandY {
			expandY += space - 1
		}

	}

	return galaxies

}

func euclideanDist(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2))
}

func manhattanDist(x1, y1, x2, y2 int) float64 {
	return math.Abs(float64(x2-x1)) + math.Abs(float64(y2-y1))
}
