package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func readLines(path string) []string {
	currDir, err := os.Getwd()
	checkErr(err)
	filepath := currDir + path

	file, err := os.Open(filepath)
	checkErr(err)

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func getDigitValues(data []string) []int {
	re := regexp.MustCompile("\\d")
	var values []int

	for _, val := range data {
		v := ""
		for _, c := range val {
			stringVal := string(c)
			match := re.Match([]byte(stringVal))
			if match {
				v = v + stringVal
				break
			}
		}
		for i := len(val) - 1; i >= 0; i-- {
			stringVal := string(val[i])
			match := re.Match([]byte(stringVal))
			if match {
				v = v + stringVal
				break
			}
		}
		if len(v) == 0 {
			values = append(values, 0)
		}
		if len(v) == 1 {
			v = v + v
			newVal, err := strconv.Atoi(v)
			checkErr(err)
			values = append(values, newVal)
		} else {
			newVal, err := strconv.Atoi(v)
			checkErr(err)
			values = append(values, newVal)
		}
	}
	return values
}

func part1() {
	data := readLines("/2023/day1/input.txt")
	values := getDigitValues(data)
	var total int
	for idx, _ := range values {
		total += values[idx]
	}
	fmt.Print(total)
}

func getDesiredChars() map[string]int {
	m := make(map[string]int)

	// Digit characters
	m["1"] = 1
	m["2"] = 2
	m["3"] = 3
	m["4"] = 4
	m["5"] = 5
	m["6"] = 6
	m["7"] = 7
	m["8"] = 8
	m["9"] = 9
	m["0"] = 0

	// Spelled out
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	m["four"] = 4
	m["five"] = 5
	m["six"] = 6
	m["seven"] = 7
	m["eight"] = 8
	m["nine"] = 9
	m["zero"] = 0

	return m
}

func getIndices() {}

func part2() {
	fmt.Println("")
}

func main() {
	part1()
	part2()
}
