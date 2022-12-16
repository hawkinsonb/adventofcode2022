package main

import(
	"os"
	"fmt"
	"bufio"
	"strconv"
)

func main() {
	fmt.Printf(`Elf: %v`, WhatElfHasTheMostCalories("./input.txt"))
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func WhatElfHasTheMostCalories(filePath string) int {
	file, err  := os.Open(filePath)
	check(err)

    scanner := bufio.NewScanner(file)

	elf := 0
	elfCalories := 0

	currentHighestElf := 0
	currentHighestElfCalories := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if (elfCalories > currentHighestElfCalories) {
				currentHighestElf = elf
				currentHighestElfCalories = elfCalories
			}

			elfCalories = 0
			elf++
		} else {
			integer, _ := strconv.Atoi(line)
			elfCalories += integer
		}
    }

	return currentHighestElf
}
