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

	text, _ := os.ReadFile("./Input/3.txt")

	times := re.FindAll(text, -1)
	total := 0

	for _, byte_arr := range times {
		str := string(byte_arr)

		str = strings.Replace(str, "mul(", "", -1)
		str = strings.Replace(str, ")", "", -1)
		split := strings.Split(str, ",")
		left, _ := strconv.Atoi(split[0])
		right, _ := strconv.Atoi(split[1])
		total += (left * right)
	}

	fmt.Println(total)
}
