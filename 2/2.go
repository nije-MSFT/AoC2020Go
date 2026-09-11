package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("./2/input.txt")
	defer file.Close()

	var count1, count2 int = 0, 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if isPasswordValid(scanner.Text()) {
			count1++
		}
		if isPasswordValid2(scanner.Text()) {
			count2++
		}
	}

	fmt.Println("Number of passwords valid:", count1)
	fmt.Println("Number of passwords valid:", count2)
}

func isPasswordValid(line string) bool {
	split := strings.Split(line, ": ")
	password := split[1]
	split = strings.Split(split[0], " ")
	character := split[1]
	split = strings.Split(split[0], "-")
	min, _ := strconv.Atoi(split[0])
	max, _ := strconv.Atoi(split[1])

	count := strings.Count(password, character)

	return count >= min && count <= max
}

func isPasswordValid2(line string) bool {
	split := strings.Split(line, ": ")
	password := []rune(split[1])
	split = strings.Split(split[0], " ")
	character := ([]rune(split[1]))[0]
	split = strings.Split(split[0], "-")
	first, _ := strconv.Atoi(split[0])
	second, _ := strconv.Atoi(split[1])
	return (password[first-1] == character) != (password[second-1] == character)
}
