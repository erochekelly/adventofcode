package main

import (
    "bufio"
    "fmt"
    "log"
    "math"
    "os"
    "sort"
    "strconv"
    "strings"
)

func count(slice []int, i int) int {
    count := 0
    for _, s := range slice {
        if int(s) == i {
            count++
        }
    }
    return count
}

func main() {
    var left []int
    var right []int
    distance := 0
    similarity := 0

    readFile, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
        log.Fatal(err)
    }

    fileScanner := bufio.NewScanner(readFile)
    for fileScanner.Scan() {
        id := strings.Fields(fileScanner.Text())
        l, _ := strconv.Atoi(id[0])
        r, _ := strconv.Atoi(id[1])
        left = append(left, l)
        right = append(right, r)
    }
    sort.Ints(left)
    sort.Ints(right)
    for i, v := range left {
        distance += int(math.Abs(float64(v) - float64(right[i])))
        similarity += count(right, v) * v
    }
    fmt.Println("Part 1: distance is", distance)
    fmt.Println("Part 2: similarity is", similarity)
}
