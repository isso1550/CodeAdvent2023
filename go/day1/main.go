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
	//part_two_test()
}

func part_one() {
	var line string
	var matches []string
	var num int
	var sum int = 0

	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line = scanner.Text()
		regex := regexp.MustCompile(`\d`)
		matches = regex.FindAllString(line, -1)

		if len(matches) > 1 {
			num, _ = strconv.Atoi(matches[0] + matches[len(matches)-1])
		} else if len(matches) == 1 {
			num, _ = strconv.Atoi(matches[0] + matches[0])
		} else {
			os.Exit(1)
		}
		sum += num
	}

	fmt.Println(sum)
}
func part_two_test() {
	/*Funkcja testująca alternatywne podejście działające dla innego zestawu danych*/
	var line string
	var matches []string
	var num int
	var sum int = 0
	var sum2 int = 0

	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var first_digit string
	var second_digit string
	digits_str := [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	regex_str := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine)`)
	regex := regexp.MustCompile(`\d`)

	for scanner.Scan() {
		line = scanner.Text()
		line_orig := line
		line = regex_str.ReplaceAllStringFunc(line, replace_string)

		matches = regex.FindAllString(line, -1)

		if len(matches) > 1 {
			num, _ = strconv.Atoi(matches[0] + matches[len(matches)-1])
		} else if len(matches) == 1 {
			num, _ = strconv.Atoi(matches[0] + matches[0])
		} else {
			os.Exit(1)
		}
		sum += num

		num1 := num
		first_digit = find_first_digit(line_orig, digits_str[:], regex)
		second_digit = find_last_digit(line_orig, digits_str[:], regex)

		if second_digit != "" {
			num, _ = strconv.Atoi(first_digit + second_digit)
		} else {
			num, _ = strconv.Atoi(first_digit + first_digit)
		}
		sum2 += num
		if num != num1 {
			fmt.Println(line_orig, line, num1, num, sum, sum2)
		}
	}

	fmt.Println(sum)
}

func replace_string(str string) string {
	digits_str := [10]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for digit, digit_str := range digits_str {
		if digit_str == str {
			return strconv.Itoa(digit + 1)
		}
	}
	fmt.Println(str)
	return ""
}

func part_two() {
	var line string
	var first_digit string
	var second_digit string
	var num int
	var sum int = 0

	digits_str := [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line = scanner.Text()
		regex := regexp.MustCompile(`\d`)
		first_digit = find_first_digit(line, digits_str[:], regex)
		second_digit = find_last_digit(line, digits_str[:], regex)

		if second_digit != "" {
			num, _ = strconv.Atoi(first_digit + second_digit)
		} else {
			num, _ = strconv.Atoi(first_digit + first_digit)
		}
		sum += num
	}

	fmt.Println(sum)
}

func find_first_digit(line string, digits_str []string, regex *regexp.Regexp) string {
	var first_digit string
	var substr string = ""
	for _, char := range line {
		substr += string(char)
		digit_matches := regex.FindAllString(substr, -1)
		if len(digit_matches) > 0 {
			first_digit = digit_matches[0]
			return first_digit
		} else {
			for digit, digit_str := range digits_str {
				digit_str_regex := regexp.MustCompile(digit_str)
				matches := digit_str_regex.FindAllString(substr, -1)
				if len(matches) > 0 {
					first_digit = strconv.Itoa(digit)
					return first_digit
				}
			}
		}

	}
	return ""
}

func find_last_digit(line string, digits_str []string, regex *regexp.Regexp) string {
	var first_digit string
	var substr string = ""
	for i := len(line) - 1; i >= 0; i-- {
		char := line[i]
		substr = string(char) + substr
		digit_matches := regex.FindAllString(substr, -1)
		if len(digit_matches) > 0 {
			first_digit = digit_matches[0]
			return first_digit
		} else {
			for digit, digit_str := range digits_str {
				digit_str_regex := regexp.MustCompile(digit_str)
				matches := digit_str_regex.FindAllString(substr, -1)
				if len(matches) > 0 {
					first_digit = strconv.Itoa(digit)
					return first_digit
				}
			}
		}

	}
	return ""
}
