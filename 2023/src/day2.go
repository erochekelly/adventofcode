package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var max = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func goodDraw(draw string) bool {
	balls := strings.Split(draw, ",")
	var count int
	var color string
	for _, b := range balls {
		fmt.Sscanf(b, "%d %s", &count, &color)
		if max[color] < count {
			return false
		}
	}
	return true
}

func ballCount(draw string) map[string]int {
	balls := strings.Split(draw, ",")
	var countColors = map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	var count int
	var color string
	for _, b := range balls {
		fmt.Sscanf(b, "%d %s", &count, &color)
		countColors[color] = count

	}
	return countColors
}

func main() {
	var game int
	var sum int
	var powerSum int
	var good bool

	readFile, err := os.Open("../data/day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		good = true
		var min = map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		input := strings.Split(fileScanner.Text(), ":")
		fmt.Sscanf(input[0], "Game %d", &game)
		draws := strings.Split(input[1], ";")

		for _, d := range draws {
			if !goodDraw(d) {
				good = false
			}
			countDraw := ballCount(d)
			for c, _ := range countDraw {
				if countDraw[c] > min[c] {
					min[c] = countDraw[c]
				}
			}
		}
		if good {
			sum += game
		}
		powerSum += min["red"] * min["green"] * min["blue"]
	}
	fmt.Println("Part 1", sum)

	fmt.Println("Part 2", powerSum)
}
