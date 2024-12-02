package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./Go/Day1/1-input.txt")
	if err != nil {
		panic(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	left, right := make([]int, 1000), make([]int, 1000)

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "   ")
		left_num, _ := strconv.Atoi(split[0])
		right_num, _ := strconv.Atoi(split[1])
		left = append(left, left_num)
		right = append(right, right_num)
	}

	sort.Sort(sort.IntSlice(left))
	sort.Sort(sort.IntSlice(right))

	total := 0
	for i := range len(left) {
		dif := math.Abs(float64(left[i] - right[i]))
		total += int(dif)
	}

	fmt.Println(total)
}
