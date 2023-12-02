package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	part_one()
	part_two()
}

func part_one() {
	var line string
	var matches [][]string
	var red_limit, green_limit, blue_limit, id, n, sum int
	var possible bool

	sum = 0
	red_limit = 12
	green_limit = 13
	blue_limit = 14

	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	regex := regexp.MustCompile(`(\d+) red|(\d+) green|(\d+) blue`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		split := strings.Split(line, ":")
		game := split[0]
		id, _ = strconv.Atoi(strings.Split(game, " ")[1])
		cubes := split[1]

		possible = true
		matches = regex.FindAllStringSubmatch(cubes, -1)

		for _, match := range matches {
			if strings.Contains(match[0], "red") {
				n, _ = strconv.Atoi(match[1])
				if n > red_limit {
					possible = false
					break
				}
			} else if strings.Contains(match[0], "green") {
				n, _ = strconv.Atoi(match[2])
				if n > green_limit {
					possible = false
					break
				}
			} else {
				n, _ = strconv.Atoi(match[3])
				if n > blue_limit {
					possible = false
					break
				}

			}
		}
		if possible {
			sum += id
		}
	}
	fmt.Println(sum)
}

func part_two() {
	var line string
	var matches [][]string
	var red_min, green_min, blue_min, n, sum int

	sum = 0
	red_min = 0
	green_min = 0
	blue_min = 0

	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	regex := regexp.MustCompile(`(\d+) red|(\d+) green|(\d+) blue`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		split := strings.Split(line, ":")
		cubes := split[1]

		matches = regex.FindAllStringSubmatch(cubes, -1)

		for _, match := range matches {
			if strings.Contains(match[0], "red") {
				n, _ = strconv.Atoi(match[1])
				if n > red_min {
					red_min = n
				}
			} else if strings.Contains(match[0], "green") {
				n, _ = strconv.Atoi(match[2])
				if n > green_min {
					green_min = n
				}
			} else {
				n, _ = strconv.Atoi(match[3])
				if n > blue_min {
					blue_min = n
				}

			}
		}
		power := red_min * green_min * blue_min
		red_min = 0
		green_min = 0
		blue_min = 0
		sum += power
	}
	fmt.Println(sum)
}
