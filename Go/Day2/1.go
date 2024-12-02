package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.Open("./Input/2.txt")

	scanner := bufio.NewScanner(data)
	total := 0

	for scanner.Scan() {
		arr := line_to_arr(scanner.Text())
		fmt.Println(arr)
		result := eval_line(arr)
		fmt.Println(result)
		if result {
			total += 1
		}
	}

	fmt.Println(total)
}

func eval_line(arr []int) bool {
	var eval func(prev int, curr int) bool
	var prev int
	for i, v := range arr {
		if i == 0 {
			prev = v
			continue
		}
		if eval == nil {
			fmt.Println("setting eval func")
			eval = set_eval(prev, v)
		}
		if int(math.Abs(float64(prev-v))) > 3 || int(math.Abs(float64(prev-v))) < 1 {
			return false
		}
		if !eval(prev, v) {
			return false
		}
		prev = v
	}

	return true
}

func set_eval(prev int, curr int) func(int, int) bool {
	if prev < curr {
		return func(prev int, curr int) bool {
			fmt.Printf("Checking increasing %d - %d\n", prev, curr)
			return prev < curr
		}
	} else {
		return func(prev, curr int) bool {
			fmt.Printf("Checking decreasing %d - %d\n", prev, curr)

			return prev > curr
		}
	}
}

func line_to_arr(line string) []int {
	arr_str := strings.Split(line, " ")
	arr := make([]int, len(arr_str))
	for i, v := range arr_str {
		num, _ := strconv.Atoi(v)
		arr[i] = num
	}

	return arr
}