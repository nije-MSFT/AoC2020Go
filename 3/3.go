package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("./3/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var treeMap [][]rune

	for scanner.Scan() {
		treeMap = append(treeMap, []rune(scanner.Text()))
	}

	treeCount := countTrees(treeMap, 1, 1)
	treeCount *= countTrees(treeMap, 3, 1)
	treeCount *= countTrees(treeMap, 5, 1)
	treeCount *= countTrees(treeMap, 7, 1)
	treeCount *= countTrees(treeMap, 1, 2)

	fmt.Println("Number of trees:", treeCount)
}

func countTrees(treeMap [][]rune, slopeX int, slopeY int) int {
	var x, y, treeCount int = 0, 0, 0

	for y < len(treeMap) {
		if treeMap[y][x] == '#' {
			treeCount++
		}
		x = (x + slopeX) % len(treeMap[0])
		y = y + slopeY
	}

	return treeCount
}
