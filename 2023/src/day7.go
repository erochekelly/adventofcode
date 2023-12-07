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

type handType int

const (
	high      handType = iota + 1 // High Card
	pair                          // One Pair
	twoPair                       // Two pair
	three                         // Three of a kind
	fullHouse                     // Full House
	four                          // 4 of a kind
	five                          // 5 of a kind
)

var jokers = true // Set to false for Part 1

type Hand struct {
	Cards string
	Bid   int
	Type  handType
}

type Hands []Hand

func (s Hands) Len() int      { return len(s) }
func (s Hands) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type Order struct{ Hands }

func (s Order) Less(i, j int) bool {
	if s.Hands[i].Type == s.Hands[j].Type {
		for c := 0; c < 5; c++ {
			vi := getValue(s.Hands[i].Cards[c : c+1])
			vj := getValue(s.Hands[j].Cards[c : c+1])
			if vi != vj {
				return vi < vj
			}
		}

	}
	return s.Hands[i].Type < s.Hands[j].Type

}

func getValue(s string) int {
	switch s {
	case "A":
		return 14
	case "K":
		return 13
	case "Q":
		return 12
	case "J":
		if jokers {
			return 1
		}
		return 11
	case "T":
		return 10
	}
	v, _ := strconv.Atoi(s)
	return v
}

func getType(s string) handType {
	cardCount := make(map[rune]int)
	var sum int

	for _, c := range s {
		n := strings.Count(s, string(c))
		if n == 5 {
			return five
		} else {
			cardCount[c] = n
			sum += n
		}
	}

	switch sum {
	case 17: // XXXXy : 4+4+4+4+1
		if jokers { // xxxxJ
			switch cardCount['J'] {
			case 1: // xxxJx
				return five
			case 4: // JJxJJ
				return five
			}
		}
		return four
	case 13: // XXXYY : 3+3+3+2+2
		if jokers && cardCount['J'] != 0 { // either JJJxx or xxxJJ
			return five
		}
		return fullHouse
	case 11: // XXXyz : 3+3+3+1+1
		if jokers {
			switch cardCount['J'] {
			case 1: // xxxJy
				return four
			case 2: // xxxJJ
				return five
			case 3: // JJJxy
				return four
			}
		}
		return three
	case 9: // XXyZZ : 2+2+1+2+2
		if jokers {
			switch cardCount['J'] {
			case 1: // xxJyy
				return fullHouse
			case 2: // JJyxx
				return four
			}
		}
		return twoPair
	case 7: // XXwyz : 2+2+1+1+1
		if jokers {
			switch cardCount['J'] {
			case 1: // xxJyz
				return three
			case 2: // JJxyz
				return three
			}
		}
		return pair
	}
	if jokers && cardCount['J'] != 0 { // Jwxyz
		return pair
	}
	return high
}

func main() {
	var hands []Hand
	var h Hand
	var winnings int

	readFile, err := os.Open("../data/day7/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		handAndBid := strings.Split(fileScanner.Text(), " ")
		h.Cards = handAndBid[0]
		h.Bid, _ = strconv.Atoi(handAndBid[1])
		h.Type = getType(h.Cards)
		hands = append(hands, h)

	}
	sort.Sort(Order{hands})
	for rank, h := range hands {
		winnings += h.Bid * (rank + 1)
	}
	if !jokers {
		fmt.Println("Part 1:", winnings)
	} else {
		fmt.Println("Part 2:", winnings)
	}
}
