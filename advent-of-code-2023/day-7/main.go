package main

import (
	"bufio"
	"fmt"
	"log"
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

type Hand struct {
	key   string
	name  string
	cards string
	bid   int
	point int
	// rank int
}

var handMap = map[string]Hand{
	"hc": {
		key:   "hc",
		name:  "High card",
		point: 0,
	},
	"1p": {
		key:   "1p",
		name:  "One pair",
		point: 1,
	},
	"2p": {
		key:   "2p",
		name:  "Two pair",
		point: 2,
	},
	"3k": {
		key:   "3k",
		name:  "Three of a kind",
		point: 3,
	},
	"fh": {
		key:   "fh",
		name:  "Full house",
		point: 4,
	},
	"4k": {
		key:   "4k",
		name:  "Four of a kind",
		point: 5,
	},
	"5k": {
		key:   "5k",
		name:  "Five of a kind",
		point: 6,
	},
}

var pointsMap = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

func partOne(file *os.File) {
	scanner := bufio.NewScanner(file)

	hands := make([]Hand, 0)

	for scanner.Scan() {
		text := scanner.Text()

		fields := strings.Fields(text)
		cards, bid := fields[0], atoi(fields[1])

		currHand := handMap["hc"]
		countMap := make(map[string]int)

		for _, l := range cards {
			countMap[string(l)] = strings.Count(cards, string(l))
		}

		counts := make([]int, 0)

		for _, count := range countMap {
			counts = append(counts, count)
		}

		sort.Sort(sort.Reverse(sort.IntSlice(counts)))

		fmt.Println(counts)

		if counts[0] == 5 {
			currHand = handMap["5k"]
		}
		if counts[0] == 4 {
			currHand = handMap["4k"]
		}
		if counts[0] == 2 {
			currHand = handMap["1p"]
		}
		if counts[0] == 3 && counts[1] == 2 {
			currHand = handMap["fh"]
		}
		if counts[0] == 3 && counts[1] == 1 {
			currHand = handMap["3k"]
		}
		if counts[0] == 2 && counts[1] == 2 {
			currHand = handMap["2p"]
		}
		if counts[0] == 2 && counts[1] == 1 {
			currHand = handMap["1p"]
		}

		currHand.bid = bid
		currHand.cards = cards

		hands = append(hands, currHand)
	}

	// Rank

	sort.Slice(hands, func(i, j int) bool {

		if hands[i].point == hands[j].point {
			v := true
			for p := range hands[i].cards {

				pointI, pointJ := pointsMap[string(hands[i].cards[p])], pointsMap[string(hands[j].cards[p])]

				if pointI == pointJ {
					continue
				}

				v = pointI < pointJ

				break
			}

			return v
		}

		return hands[i].point < hands[j].point
	})
	// fmt.Println(hands)

	totalWinnings := 0

	for i, v := range hands {
		totalWinnings += v.bid * (i + 1)
	}

	fmt.Println("Total Winnings: ", totalWinnings)

}

func partTwo(file *os.File) {
	scanner := bufio.NewScanner(file)

	hands := make([]Hand, 0)
	pointsMap["J"] = 1

	for scanner.Scan() {
		text := scanner.Text()

		fields := strings.Fields(text)
		cards, bid := fields[0], atoi(fields[1])

		currHand := handMap["hc"]
		countMap := make(map[string]int)

		for _, l := range cards {
			countMap[string(l)] = strings.Count(cards, string(l))
		}

		countJ := countMap["J"]
		if 0 < countMap["J"] && countMap["J"] < 5 {
			delete(countMap, "J")

			max := ""

			for k, v := range countMap {
				if countMap[max] < v {
					max = k
				}
			}

			countMap[max] += countJ
		}

		counts := make([]int, 0)
		for _, count := range countMap {
			counts = append(counts, count)
		}

		sort.Sort(sort.Reverse(sort.IntSlice(counts)))

		if counts[0] == 5 {
			currHand = handMap["5k"]
		}
		if counts[0] == 4 {
			currHand = handMap["4k"]
		}
		if counts[0] == 2 {
			currHand = handMap["1p"]
		}
		if counts[0] == 3 && counts[1] == 2 {
			currHand = handMap["fh"]
		}
		if counts[0] == 3 && counts[1] == 1 {
			currHand = handMap["3k"]
		}
		if counts[0] == 2 && counts[1] == 2 {
			currHand = handMap["2p"]
		}
		if counts[0] == 2 && counts[1] == 1 {
			currHand = handMap["1p"]
		}

		currHand.bid = bid
		currHand.cards = cards

		hands = append(hands, currHand)
	}

	// Rank

	sort.Slice(hands, func(i, j int) bool {

		if hands[i].point == hands[j].point {
			v := true
			for p := range hands[i].cards {

				pointI, pointJ := pointsMap[string(hands[i].cards[p])], pointsMap[string(hands[j].cards[p])]

				if pointI == pointJ {
					continue
				}

				v = pointI < pointJ

				break
			}

			return v
		}

		return hands[i].point < hands[j].point
	})
	// fmt.Println(hands)

	totalWinnings := 0

	for i, v := range hands {
		totalWinnings += v.bid * (i + 1)
	}

	fmt.Println("Total Winnings: ", totalWinnings)
}

func atoi(s string) int {

	n, err := strconv.Atoi(s)

	if err != nil {
		log.Panic(err)
	}

	return n
}

// Five of a kind
// Four of a kind
