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

type Robot struct {
	sp [2]int
	p  [2]int
	v  [2]int
}

const grows = 101
const gcols = 103

// const grows = 7
// const gcols = 11
const cx, cy = grows / 2, gcols / 2

const secs = 100

func partOne(file *os.File) {
	scanner := bufio.NewScanner(file)

	robots := make([]Robot, 0)

	for scanner.Scan() {
		text := scanner.Text()
		rI := strings.Fields(text)

		r := Robot{}
		r.p = extractVectors(rI[0])
		r.sp = r.p
		r.v = extractVectors(rI[1])

		robots = append(robots, r)
	}

	for i := 0; i < secs; i++ {
		// Update position
		for j := 0; j < len(robots); j++ {
			rob := &robots[j]
			rob.p[0] += rob.v[0]
			rob.p[1] += rob.v[1]

			//  loop grid
			if rob.p[0] >= grows {
				rob.p[0] %= grows
			}
			if rob.p[1] >= gcols {
				rob.p[1] %= gcols
			}

			if rob.p[0] < 0 {
				rob.p[0] += grows
			}
			if rob.p[1] < 0 {
				rob.p[1] += gcols
			}
		}
	}

	quad := [4]int{}
	for i := 0; i < len(robots); i++ {
		rob := &robots[i]
		x, y := rob.p[0], rob.p[1]

		if x == cx || y == cy {
			continue
		}

		if x < cx && y < cy {
			quad[0]++
		} else if x < cx && y >= cy {
			quad[1]++
		} else if x >= cx && y < cy {
			quad[2]++
		} else if x >= cx && y >= cy {
			quad[3]++
		}
	}

	sf := 1
	for _, v := range quad {
		if v == 0 {
			continue
		}

		sf *= v
	}

	fmt.Println("Part one: ", sf)
}

func partTwo(file *os.File) {
	scanner := bufio.NewScanner(file)

	robots := make([]Robot, 0)

	for scanner.Scan() {
		text := scanner.Text()
		rI := strings.Fields(text)

		r := Robot{}
		r.p = extractVectors(rI[0])
		r.sp = r.p
		r.v = extractVectors(rI[1])

		robots = append(robots, r)
	}

	s := 0
	var points [][]int

	for _, v := range robots {
		points = append(points, v.p[:])
	}

	for !isChristmasTree(points, grows, gcols) {
		// Update position
		for j := 0; j < len(robots); j++ {
			rob := &robots[j]
			rob.p[0] += rob.v[0]
			rob.p[1] += rob.v[1]

			//  loop grid
			if rob.p[0] >= grows {
				rob.p[0] %= grows
			}
			if rob.p[1] >= gcols {
				rob.p[1] %= gcols
			}

			if rob.p[0] < 0 {
				rob.p[0] += grows
			}
			if rob.p[1] < 0 {
				rob.p[1] += gcols
			}
		}

		points = [][]int{}
		for _, v := range robots {
			points = append(points, v.p[:])
		}
		s++
	}

	fmt.Println("Part two: ", s)
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Panic(err)
	}
	return n
}

func extractVectors(str string) [2]int {
	vec := make([]int, 2)

	p := strings.Split(str, "=")[1]
	px := strings.Split(p, ",")

	for i, v := range px {
		vec[i] = toInt(v)
	}

	return [2]int(vec)
}

func isChristmasTree(points [][]int, rows int, cols int) bool {
	grid := make(map[[2]int]bool)
	for _, point := range points {
		grid[[2]int{point[0], point[1]}] = true
	}

	// Find the apex (single point in the top row)
	apex := -1
	center := -1
	for r := 0; r < rows; r++ {
		count := 0
		for c := 0; c < cols; c++ {
			if grid[[2]int{r, c}] {
				count++
				center = c
			}
		}
		if count == 1 {
			apex = r
			break
		}
	}
	if apex == -1 {
		return false // No apex found
	}

	// Check triangle symmetry
	for r := apex; r < rows; r++ {
		width := r - apex
		start := int(math.Max(float64(0), float64(center-width)))
		end := int(math.Min(float64(cols-1), float64(center+width)))
		for c := start; c <= end; c++ {
			if !grid[[2]int{r, c}] {
				return false
			}
		}
	}

	// Check trunk (assuming it's at the center directly below the broadest part of the triangle)
	trunkStart := rows / 2 // This is an arbitrary start point; may need to adjust as per definition
	for r := trunkStart; r < rows; r++ {
		if !grid[[2]int{r, center}] {
			return false
		}
	}

	return true
}
