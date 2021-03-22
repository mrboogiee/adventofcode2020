package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.
*/

var (
	wanted          = "shiny gold"
	bagsDefinitions map[string]map[string]int //color name, sub bags names, amount
)

func main() {
	inputs := readFile()
	bagsDefinitions = make(map[string]map[string]int)
	for _, line := range inputs {
		convertLine(line)
	}
	fmt.Print(getAmountOfGoldBags())

}

func readFile() []string {
	file, err := os.Open("./example.txt")
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

func convertLine(line string) /*bag*/ {
	bagContents := strings.Split(line, " bags contain ")
	bagsDefinitions[bagContents[0]] = make(map[string]int)
	bags := strings.Split(bagContents[1], ", ")
	for _, subBag := range bags {
		reBags := regexp.MustCompile(`(?m)([0-9]+) (([a-zA-Z]*) ([a-zA-Z]*))`)
		for _, match := range reBags.FindAllStringSubmatch(subBag, -1) {
			amountOfBags, err := strconv.Atoi(match[1])
			if err != nil {
				log.Fatalln(err)
			}
			bagsDefinitions[bagContents[0]][match[2]] = amountOfBags
		}
	}
}

func getAmountOfGoldBags() int {
	var bagsToReturn int
	for _, subBags := range bagsDefinitions {
		var canReturn bool
		if canContainWanted(subBags) {
			canReturn = true
		}
		if canReturn {
			bagsToReturn++
		}
	}
	return bagsToReturn
}

func canContainWanted(bags map[string]int) bool {
	for bag := range bags {
		if bag == wanted {
			return true
		} else {
			for name, subBags := range bagsDefinitions {
				if bag == name {
					return canContainWanted(subBags)
				}
			}
		}
	}
	return false
}
