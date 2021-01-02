package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Group struct {
	members int
	answers map[string]int
}

func main() {
	inputs := readFile()
	groups := createGroups(inputs)
	var total int
	for _, group := range groups {
		allYesCount := 0
		fmt.Println(group)
		for _, count := range group.answers {
			fmt.Println(count, group.members)
			if count == group.members {
				allYesCount++
			}
		}
		fmt.Println(allYesCount)
		total += allYesCount
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

func createGroups(inputs []string) []Group {
	group := Group{}
	group.answers = make(map[string]int)
	groups := []Group{}
	for i, line := range inputs {
		if line != "" {
			group.members++
			for _, character := range line {
				group.answers[string(character)]++
			}
		}
		if line == "" || i+1 == len(inputs) {
			groups = append(groups, group)
			group = Group{}
			group.answers = make(map[string]int)

		}
	}
	return groups
}
