package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	re := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)

	bytes, _ := os.ReadFile("./Input/3.txt")
	text := string(bytes)

	fmt.Println(text)

	var new_str string
	start := 0
	index := 0
	total := 0

	for {
		fmt.Printf("[%d:%d]\n", start, index)
		if index+4 >= len(text) {
			break
		}

		index = strings.Index(text[index:], "don't()")
		if index == -1 {
			break
		}
		new_str += text[start:index]
		fmt.Println(new_str)
		start = strings.Index(text[index:], "do()")
		if start == -1 {
			break
		}
		index += 4
	}

	times := re.FindAll([]byte(new_str), -1)
	for _, time_byte := range times {
		time := string(time_byte)
		str := strings.Replace(time, "mul(", "", -1)
		str = strings.Replace(time, ")", "", -1)
		split := strings.Split(str, ",")
		left, _ := strconv.Atoi(split[0])
		right, _ := strconv.Atoi(split[1])
		total += (left * right)
	}

	fmt.Println(total)
}
