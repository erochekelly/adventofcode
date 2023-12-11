package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type Node struct {
	left  string
	right string
}

type Nodes map[string]*Node

var nonNodeRegex = regexp.MustCompile(`[^A-Z0-9 ]+`)

func clearString(str string) string {
	return nonNodeRegex.ReplaceAllString(str, "")
}

func findZZZ(n Nodes, e string, rl []string, index int, moves int) int {
	moves++
	if index > len(rl)-1 {
		index = 0
	}
	if rl[index] == "L" {
		if n[e].left == "ZZZ" {
			return moves
		} else {
			e = n[e].left
			index++
		}
	} else {
		if n[e].right == "ZZZ" {
			return moves
		} else {
			e = n[e].right
			index++
		}
	}
	return findZZZ(n, e, rl, index, moves)
}

func findZ(n Nodes, e string, rl []string, index int, moves int) int {
	moves++
	if index > len(rl)-1 {
		index = 0
	}
	if rl[index] == "L" {
		if n[e].left[2] == 'Z' {
			return moves
		} else {
			e = n[e].left
			index++
		}
	} else {
		if n[e].right[2] == 'Z' {
			return moves
		} else {
			e = n[e].right
			index++
		}
	}
	return findZ(n, e, rl, index, moves)
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(vs []int) int {
	if len(vs) == 0 {
		return 1
	}
	lcm := vs[0]
	for _, v := range vs[1:] {
		d := GCD(lcm, v)
		lcm = lcm / d * v
	}
	return lcm
}

func main() {
	var rl []string
	var all []string
	var n, l, r string
	var movesAll []int
	nodes := make(Nodes)
	readFile, err := os.Open("../data/day8/input.txt")
	// readFile, err := os.Open("day8.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	// start := time.Now()

	for fileScanner.Scan() {
		if fileScanner.Text() == "" {
			continue
		} else if !strings.Contains(fileScanner.Text(), "=") {
			rl = strings.Split(fileScanner.Text(), "")
		} else {

			line := clearString(fileScanner.Text())
			fmt.Sscanf(line, "%s %s %s", &n, &l, &r)
			nodes[n] = &Node{l, r}
		}
	}
	for k, _ := range nodes {
		if k[2] == 'A' {
			all = append(all, k)
		}
	}
	moves := findZZZ(nodes, "AAA", rl, 0, 0)

	fmt.Println("Part1: ", moves)

	for _, a := range all {
		movesAll = append(movesAll, findZ(nodes, a, rl, 0, 0))
	}

	fmt.Println("Part2: ", LCM(movesAll))

}
