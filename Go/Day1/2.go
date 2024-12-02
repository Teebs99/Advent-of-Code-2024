package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("./Go/Day1/1-input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	left := make([]int, 1000)
	count := make(map[int]int)
	for scanner.Scan() {
		line := scanner.Text()
		split_str := strings.Split(line, "   ")
		left_num, _ := strconv.Atoi(split_str[0])
		left = append(left, left_num)
		Count(count, split_str[1])
	}

	total := 0
	for _, v := range left {
		multiple := 0
		if val, ok := count[v]; ok {
			multiple = val
		}
		total += (v * multiple)
	}
	fmt.Println(total)
}

func Count(count map[int]int, num_str string) {
	num, _ := strconv.Atoi(num_str)
	count[num] += 1
}
