package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strings"
)

type Point struct {
	X int
	Y int
}
type Galaxy []Point

func main() {
	var g Galaxy
	var y int
	var x int
	topLeft := Point{0, 0}

	readFile, err := os.Open("../data/day11/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		var p Point

		line := strings.Split(fileScanner.Text(), "")
		x = len(line)
		for i, r := range line {
			p.X = i
			p.Y = y
			if r == "#" {
				g = append(g, p)
			}
		}
		y++
	}
	bottomRight := Point{x, y}

	fmt.Println("Part 1: ", getDistance(expand(slices.Clone(g), topLeft, bottomRight, 2)))
	fmt.Println("Part 2: ", getDistance(expand(slices.Clone(g), topLeft, bottomRight, 1000000)))

}

func expand(g Galaxy, t, b Point, m int) Galaxy {
	n := m - 1
	var column, row bool
xAxis:
	for x := t.X; x < b.X; x++ {
		column = false
		for _, p := range g {
			if p.X == x {
				column = true
				continue xAxis

			}
		}

		if !column {
			for i, p := range g {
				if p.X > x {
					g[i].X += n
					b.X += n

				}
			}
			x += n

		}
	}
yAxis:
	for y := t.Y; y < b.Y; y++ {
		row = false
		for _, p := range g {
			if p.Y == y {
				row = true
				continue yAxis

			}
		}
		if !row {
			for i, p := range g {
				if p.Y > y {
					g[i].Y += n
					b.Y += n
				}

			}
			y += n
		}
	}
	return g
}

func getDistance(g Galaxy) int {
	var sum int
	for i, a := range g {
		for _, b := range g[i+1:] {
			x := math.Abs(float64(b.X - a.X))
			y := math.Abs(float64(b.Y - a.Y))
			sum += int(x + y)

		}

	}
	return sum
}
