package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	inputs := readFile()
	groups := createGroups(inputs)
	var total int
	for _, group := range groups {
		fmt.Println(len(group))
		total += len(group)
	}
	fmt.Printf("Total %d", total)
}

func readFile() []string {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var inputs []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		inputs = append(inputs, input)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return inputs
}

func createGroups(inputs []string) []map[string]bool {
	group := make(map[string]bool)
	groups := make([]map[string]bool, 0)
	for i, line := range inputs {
		if line != "" {
			for _, character := range line {
				group[string(character)] = true
			}
		}
		if line == "" || i+1 == len(inputs) {
			// passports = append(passports, passport)
			groups = append(groups, group)
			// passport = map[string]string{}
			group = map[string]bool{}

		}
	}
	return groups
}
