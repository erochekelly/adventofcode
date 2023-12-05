package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	var sum int
	var prize int
	var winners int
	var cardCount int
	cards := make([]int, 256)
	readFile, err := os.Open("../data/day4/input.txt")
	// readFile, err := os.Open("test")
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {

		game := strings.Split(fileScanner.Text(), ":")
		g, _ := strconv.Atoi(strings.Fields(game[0])[1])
		numbers := strings.Split(game[1], "|")
		winning := strings.Fields(numbers[0])
		mine := strings.Fields(numbers[1])
		prize = 0
		winners = 0
		cards[g]++
		for _, w := range winning {
			if slices.Contains(mine, w) {
				winners++
				if prize == 0 {
					prize = 1
				} else {
					prize *= 2
				}
			}
		}
		sum += prize
		for j := 0; j < cards[g]; j++ {
			for i := 1; i <= winners; i++ {
				cards[g+i]++
			}
		}
	}
	for _, i := range cards {
		cardCount += i
	}
	fmt.Println("Part 1", sum)
	fmt.Println("Part 2", cardCount)
}
