package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1ValidPasswords, err := validatePassword(func(minOccurrence, maxOccurrence int, requiredRune rune, passwordRunes []rune) bool {
		occurrences := 0
		for _, passwordRune := range passwordRunes {
			if passwordRune == requiredRune {
				occurrences++
			}
		}
		if minOccurrence <= occurrences && maxOccurrence >= occurrences {
			return true
		}
		return false
	}, false)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("part 1 valid password count: %d", part1ValidPasswords)

	part2ValidPasswordCount, err := validatePassword(func(firstPosition, secondPosition int, requiredRune rune, passwordRunes []rune) bool {
		isValid := func(position int) bool {
			// policy is 1-based
			if len(passwordRunes) > (position - 1) {
				return passwordRunes[position-1] == requiredRune
			}
			return false
		}
		firstPositionValid := isValid(firstPosition)
		secondPositionValid := isValid(secondPosition)
		return (firstPositionValid || secondPositionValid) && !(firstPositionValid && secondPositionValid)
	}, false)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("part 2 valid password count: %d", part2ValidPasswordCount)
}

type passwordValidator func(firstPosition, secondPosition int, requiredRune rune, passwordRunes []rune) bool

func validatePassword(validate passwordValidator, describe bool) (int, error) {
	file, err := os.Open("day2/input.txt")
	if err != nil {
		return 0, nil
	}
	defer file.Close()

	validEntries := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		rangeSummaryParts := strings.Split(parts[0], "-")
		firstPosition, err := strconv.Atoi(rangeSummaryParts[0])
		if err != nil {
			return 0, nil
		}
		secondPosition, err := strconv.Atoi(rangeSummaryParts[1])
		if err != nil {
			return 0, nil
		}

		requiredRune := []rune(strings.TrimSuffix(parts[1], ":"))[0]

		password := parts[2]

		passwordRunes := []rune(password)
		if validate(firstPosition, secondPosition, requiredRune, passwordRunes) {
			if describe {
				log.Printf("valid password. firstPosition: %d. secondPosition: %d. requiredRune: %s. password: %s", firstPosition, secondPosition, string(requiredRune), string(passwordRunes))
			}
			validEntries++
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, nil
	}

	return validEntries, nil
}
