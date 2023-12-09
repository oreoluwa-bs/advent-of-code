package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

type Node struct {
	value string
	left  string
	right string
}

func partOne(file *os.File) {
	scanner := bufio.NewScanner(file)

	directions := make([]string, 0)
	paths := make(map[string]Node)

	currLine := 0
	for scanner.Scan() {
		text := scanner.Text()

		if currLine == 0 {
			directions = append(directions, strings.Split(text, "")...)
		} else {
			// AAA = (BBB, CCC)

			if strings.Contains(text, "=") {
				vals := strings.Split(text, "=")

				val, dir := strings.TrimSpace(vals[0]), strings.TrimSpace(vals[1])

				dirs := strings.Split(dir[1:len(vals[1])-2], ", ")

				paths[val] = Node{
					value: val,
					left:  dirs[0],
					right: dirs[1],
				}
			}

		}

		currLine++
	}

	// fmt.Println(directions)
	// fmt.Println(paths)

	// startNode := "AAA"
	endNode := "ZZZ"

	currNode := "AAA"
	iteration := 0

	if _, ok := paths[currNode]; !ok {
		fmt.Println("Node not found")
		return
	}

	for currNode != endNode {

		direction := directions[iteration%len(directions)]

		if direction == "R" {
			currNode = paths[currNode].right
		}
		if direction == "L" {
			currNode = paths[currNode].left
		}

		iteration++
	}

	fmt.Println("Part One:", iteration)

}

func partTwo(file *os.File) {
	scanner := bufio.NewScanner(file)

	directions := make([]string, 0)
	paths := make(map[string]Node)
	startNodes := make([]string, 0)

	currLine := 0
	for scanner.Scan() {
		text := scanner.Text()

		if currLine == 0 {
			directions = append(directions, strings.Split(text, "")...)
		} else {
			// AAA = (BBB, CCC)

			if strings.Contains(text, "=") {
				vals := strings.Split(text, "=")

				val, dir := strings.TrimSpace(vals[0]), strings.TrimSpace(vals[1])

				dirs := strings.Split(dir[1:len(vals[1])-2], ", ")

				// fmt.Println(val, strings.HasSuffix(val, "A"))
				if strings.HasSuffix(val, "A") {
					startNodes = append(startNodes, val)
				}

				paths[val] = Node{
					value: val,
					left:  dirs[0],
					right: dirs[1],
				}
			}

		}

		currLine++
	}

	// currNodes := startNodes
	// iteration := 0
	// endNodes := make([]string, 0,len(startNodes))
	results := make([]int, 0)

	// isAtEnd := false

	if len(startNodes) < 1 {
		fmt.Println("Nothing to check!")
		return
	}
	// fmt.Println(len(currNodes))

	for _, n := range startNodes {
		iteration := 0
		currNode := n

		for !strings.HasSuffix(currNode, "Z") {

			direction := directions[iteration%len(directions)]

			if direction == "R" {
				currNode = paths[currNode].right
			}
			if direction == "L" {
				currNode = paths[currNode].left
			}

			iteration++
		}

		results = append(results, iteration)
	}

	// This approach takes tooo long to compute
	// for !isAtEnd {

	// 	direction := directions[iteration%len(directions)]

	// 	for i, currNode := range currNodes {
	// 		if direction == "R" {
	// 			currNodes[i] = paths[currNode].right
	// 		}
	// 		if direction == "L" {
	// 			currNodes[i] = paths[currNode].left

	// 		}
	// 	}

	// 	// fmt.Println(currNodes)

	// 	shouldEnd := false
	// 	for _, v := range currNodes {
	// 		if !strings.HasSuffix(v, "Z") {
	// 			shouldEnd = false
	// 			break
	// 		}

	// 		shouldEnd = true
	// 	}

	// 	if shouldEnd {
	// 		isAtEnd = true
	// 	}

	// 	iteration++
	// 	fmt.Println(iteration)
	// }

	fmt.Println("Part two: ", lcm(results...))

}
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
func lcm(integers ...int) int {
	if len(integers) == 0 {
		return 0
	}

	result := integers[0]
	for _, i := range integers[1:] {
		result = (result * i) / gcd(result, i)
	}
	return result
}
