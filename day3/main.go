package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	part1()
	part2()
}

func part1() {
	trees, err := findTrees(3, 1, false)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Trees: %d", trees)
}

func part2() {
	oneOne, err := findTrees(1, 1, false)
	if err != nil {
		log.Fatal(err)
	}
	threeOne, err := findTrees(3, 1, false)
	if err != nil {
		log.Fatal(err)
	}
	fiveOne, err := findTrees(5, 1, false)
	if err != nil {
		log.Fatal(err)
	}
	sevenOne, err := findTrees(7, 1, false)
	if err != nil {
		log.Fatal(err)
	}
	oneTwo, err := findTrees(1, 2, false)
	if err != nil {
		log.Fatal(err)
	}

	product := oneOne * threeOne * fiveOne * sevenOne * oneTwo
	log.Printf("tree product: %d", product)
}

func findTrees(right, down int, describe bool) (int, error) {
	file, err := os.Open("day3/input.txt")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	trees := 0
	xPos := 0
	yPos := 0
	nextY := 0
	treeRune := '#'
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if yPos != nextY {
			yPos++
			continue
		}
		line := scanner.Text()
		lineRunes := []rune(line)
		if lineRunes[xPos] == treeRune {
			if describe {
				log.Printf("hit tree at line: %d, x: %d", yPos+1, xPos+1)
			}
			trees++
		}
		xPos = (xPos + right) % len(lineRunes)
		yPos++
		nextY += down
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return trees, nil
}
