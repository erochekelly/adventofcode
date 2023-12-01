package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var nonNumericRegex = regexp.MustCompile(`[^0-9]+`)
var combined = map[string]string{
	"twone":     "21",
	"oneight":   "18",
	"threeight": "38",
	"fiveight":  "58",
	"eightwo":   "82",
	"eighthree": "83",
	"nineight":  "98",
	"sevenine":  "79",
}
var numbers = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func words(str string) string {
	for word, digit := range combined {
		if strings.Contains(str, word) {
			str = strings.ReplaceAll(str, word, digit)
		}
	}
	for word, digit := range numbers {
		if strings.Contains(str, word) {
			str = strings.ReplaceAll(str, word, digit)
		}
	}
	return str
}

func clearString(str string) string {
	return nonNumericRegex.ReplaceAllString(str, "")
}

func main() {
	var sum1 int
	var sum2 int
	readFile, err := os.Open("../data/day1/input.txt")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {

		s1 := fileScanner.Text()
		s2 := words(fileScanner.Text())

		s1 = clearString(s1)
		s2 = clearString(s2)

		c1 := fmt.Sprintf("%s%s", string(s1[0]), string(s1[len(s1)-1]))
		c2 := fmt.Sprintf("%s%s", string(s2[0]), string(s2[len(s2)-1]))

		n1, err := strconv.Atoi(c1)
		if err == nil {
			sum1 += n1
		}
		n2, err := strconv.Atoi(c2)
		if err == nil {
			sum2 += n2
		}
	}
	fmt.Println("Part 1", sum1)
	fmt.Println("Part 2", sum2)
}
