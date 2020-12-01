package main

import (
	"github.com/dustin-graham/advent_of_code_2020/utils"
	"log"
)

func main() {
	entries, err := utils.ReadInts("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	part1(entries)
	part2(entries)
}

func part1(entries []int) {
	for i, outer := range entries {
		for j, inner := range entries {
			if i != j && outer+inner == 2020 {
				log.Printf("answer: %d", outer*inner)
				return
			}
		}
	}
}

func part2(entries []int) {
	for i, a := range entries {
		for j, b := range entries {
			for k, c := range entries {
				if i != j && i != k && j != k && a+b+c == 2020 {
					log.Printf("answer: %d", a*b*c)
					return
				}
			}
		}
	}

}
