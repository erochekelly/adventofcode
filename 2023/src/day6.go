package main

import (
	"fmt"
)

var time = []int{50, 74, 86, 85}
var distance = []int{242, 1017, 1691, 1252}

var t2 = 50748685
var d2 = 242101716911252

// var t2 = 71530
// var d2 = 940200

func main() {
	var count []int

	for i, t := range time {
		for x := 1; x < t; x++ {
			if x*(t-x) > distance[i] {
				count = append(count, t-(x*2)+1)
				x = t
			}
		}

	}
	product := 1
	for _, c := range count {
		product *= c
	}

	fmt.Println("Part 1", product)
stop:
	for x := 1; x < t2; x++ {
		if x*(t2-x) > d2 {
			fmt.Println("Part2: ", t2-(x*2)+1)
			break stop
		}
	}
}
