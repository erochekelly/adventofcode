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

func getInts(s string) []int {
	split := strings.Split(s, " ")
	numbers := make([]int, len(split))
	for i, n := range split {
		numbers[i], _ = strconv.Atoi(n)
	}
	return numbers
}

func getNext(n []int) int {
	diffs := make([]int, len(n)-1)
	if slices.Max(n) == 0 && slices.Min(n) == 0 {
		return 0
	}
	for i := 0; i < len(diffs); i++ {
		diffs[i] = n[i+1] - n[i]
	}
	return n[len(n)-1] + getNext(diffs)

}

func getPrev(n []int) int {
	diffs := make([]int, len(n)-1)
	if slices.Max(n) == 0 && slices.Min(n) == 0 {
		return 0
	}
	for i := 0; i < len(diffs); i++ {
		diffs[i] = n[i+1] - n[i]
	}
	return n[0] - getPrev(diffs)

}

func main() {
	var history []int
	var next []int
	var prev []int
	var sum int
	readFile, err := os.Open("../data/day9/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	// start := time.Now()

	for fileScanner.Scan() {
		history = getInts(fileScanner.Text())
		next = append(next, getNext(history))
		prev = append(prev, getPrev(history))

	}

	for _, n := range next {
		sum += n
	}
	fmt.Println("Part1: ", sum)
	sum = 0
	for _, n := range prev {
		sum += n
	}
	fmt.Println("Part2: ", sum)

}
