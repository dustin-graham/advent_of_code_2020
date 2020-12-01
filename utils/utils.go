package utils

import (
	"bufio"
	"os"
	"strconv"
)

func ReadInts(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	entries := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		entries = append(entries, value)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return entries, nil
}
