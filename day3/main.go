package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var inputs []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		input := scanner.Text()
		inputs = append(inputs, input)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	scenario1 := processSlope(1, 1, inputs)
	scenario2 := processSlope(3, 1, inputs)
	scenario3 := processSlope(5, 1, inputs)
	scenario4 := processSlope(7, 1, inputs)
	scenario5 := processSlope(1, 2, inputs)
	fmt.Printf("Amount of trees: %d\n", scenario1*scenario2*scenario3*scenario4*scenario5)
}

func processSlope(right, down int, inputs []string) int {
	var currentX = 0
	var currentY = 0
	var amountOfTrees int
	for j, row := range inputs {
		if j == currentY {
			for i, item := range row {
				if i == currentX {
					if item == '#' {
						amountOfTrees++
					}
					currentX += right
					if currentX > 30 {
						currentX = currentX - 31
					}
					break
				}
			}
			currentY += down
		}
	}
	return amountOfTrees
}
