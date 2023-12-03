package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	part_one()
	part_two()
}

func part_one() {
	var line string
	var symb_search_min, symb_search_max int
	var data []string
	var symb_matches []string
	var sum int

	sum = 0

	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	num_regex := regexp.MustCompile(`\d+`)
	symbol_regex := regexp.MustCompile(`[^A-Za-z0-9.]`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		data = append(data, line)
	}

	for line_idx, line := range data {
		num_matches := num_regex.FindAllStringIndex(line, -1)
		for _, num_pos := range num_matches {
			if num_pos[0] == 0 {
				symb_search_min = 0
			} else {
				symb_search_min = num_pos[0] - 1
			}

			if num_pos[1]-1 == len(line)-1 {
				symb_search_max = len(line) - 1
			} else {
				symb_search_max = num_pos[1] - 1 + 1
			}

			//Szukanie
			//prev line
			if line_idx > 0 {
				symb_matches = symbol_regex.FindAllString(data[line_idx-1][symb_search_min:symb_search_max+1], -1)
				if len(symb_matches) > 0 {
					//found in prev
					sum += find_number(line, num_pos)
					continue
				}
			}
			//same line
			symb_matches = symbol_regex.FindAllString(line[symb_search_min:symb_search_max+1], -1)
			if len(symb_matches) > 0 {
				//found in same
				sum += find_number(line, num_pos)
				continue
			}

			//next line
			if line_idx < len(data)-1 {
				symb_matches = symbol_regex.FindAllString(data[line_idx+1][symb_search_min:symb_search_max+1], -1)
				if len(symb_matches) > 0 {
					//found in next
					sum += find_number(line, num_pos)
					continue
				}
			}

		}

	}
	fmt.Println("SUM ", sum)
}

func find_number(line string, pos []int) int {
	str_cut := line[pos[0]:pos[1]]
	num, _ := strconv.Atoi(str_cut)
	return num
}

func part_two() {
	var line string
	var num_search_min, num_search_max int
	var data []string
	var num_matches [][]int
	var all_num_matches []int
	var sum int
	var gear_pos int
	var gear_num_n int

	sum = 0

	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	num_regex := regexp.MustCompile(`\d+`)
	gear_regex := regexp.MustCompile(`\*`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		data = append(data, line)
	}

	for line_idx, line := range data {
		gear_matches := gear_regex.FindAllStringIndex(line, -1)
		for _, gear_match_pos := range gear_matches {
			all_num_matches = []int{}
			gear_num_n = 0
			gear_pos = gear_match_pos[0]
			if gear_pos == 0 {
				num_search_min = 0
			} else {
				num_search_min = gear_pos - 1
			}

			if gear_pos+1 == len(line)-1 {
				num_search_max = len(line) - 1
			} else {
				num_search_max = gear_pos + 1
			}

			//Szukanie
			//prev line
			if line_idx > 0 {
				num_matches = num_regex.FindAllStringIndex(data[line_idx-1][num_search_min:num_search_max+1], -1)
				//all_num_matches = append(all_num_matches, num_matches...)
				for _, num_pos := range num_matches {
					all_num_matches = append(all_num_matches, find_gear_ratio(num_regex, data[line_idx-1], gear_pos, num_pos))
				}

				gear_num_n += len(num_matches)
			}
			//same line
			num_matches = num_regex.FindAllStringIndex(line[num_search_min:num_search_max+1], -1)
			//all_num_matches = append(all_num_matches, num_matches...)
			gear_num_n += len(num_matches)
			for _, num_pos := range num_matches {
				all_num_matches = append(all_num_matches, find_gear_ratio(num_regex, line, gear_pos, num_pos))
			}

			//next line
			if line_idx < len(data)-1 {
				num_matches = num_regex.FindAllStringIndex(data[line_idx+1][num_search_min:num_search_max+1], -1)
				for _, num_pos := range num_matches {
					all_num_matches = append(all_num_matches, find_gear_ratio(num_regex, data[line_idx+1], gear_pos, num_pos))
				}
				gear_num_n += len(num_matches)
			}
			if gear_num_n == 2 {
				gear_ratio := all_num_matches[0] * all_num_matches[1]

				sum += gear_ratio
			}

		}

	}
	fmt.Println("SUM ", sum)
}

func find_gear_ratio(num_regex *regexp.Regexp, line string, gear_pos int, num_pos []int) int {
	num_matches := num_regex.FindAllStringIndex(line, -1)
	search_idx := gear_pos - 1 + num_pos[0]
	for _, num_pos_line := range num_matches {
		if (search_idx >= num_pos_line[0]) && (search_idx <= num_pos_line[1]) {
			num, _ := strconv.Atoi(line[num_pos_line[0]:num_pos_line[1]])
			return num
		}
	}
	return 0
}
