package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
		input := scanner.Text()
		inputs = append(inputs, input)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	var seatIDs []int
	for _, seatLocation := range inputs {
		seatIDs = append(seatIDs, getSeatID(seatLocation))
	}
	sort.Ints(seatIDs)

	var mySeatID = seatIDs[0]
	for _, seatID := range seatIDs {
		if seatID == mySeatID {
			mySeatID++
		}
	}
	fmt.Println(mySeatID)

	mySeatID = seatIDs[len(seatIDs)-1]
	for i := len(seatIDs) - 1; i >= 0; i-- {
		if seatIDs[i] == mySeatID {
			mySeatID--
		}
	}
	fmt.Println(mySeatID)

}

func getSeatID(location string) int {
	var rowDefinition = location[0:7]
	var columnDefinition = location[7:10]
	return getRowNumber(rowDefinition)*8 + getColumnNumber(columnDefinition)

}

func getRowNumber(rowDefinition string) int {
	MinRowNumber := 0
	maxRowNumber := 128
	for _, character := range rowDefinition {
		if string(character) == "F" {
			maxRowNumber = maxRowNumber - ((maxRowNumber - MinRowNumber) / 2)
		} else if string(character) == "B" {
			MinRowNumber = MinRowNumber + ((maxRowNumber - MinRowNumber) / 2)
		}
	}
	if maxRowNumber-1 != MinRowNumber {
		log.Fatalln(MinRowNumber, maxRowNumber)
	}
	return MinRowNumber
}

func getColumnNumber(columnDefinition string) int {
	minColumnNumber := 0
	maxColumnNumber := 8
	for _, character := range columnDefinition {
		if string(character) == "L" {
			maxColumnNumber = maxColumnNumber - ((maxColumnNumber - minColumnNumber) / 2)
		} else if string(character) == "R" {
			minColumnNumber = minColumnNumber + ((maxColumnNumber - minColumnNumber) / 2)
		}
	}
	if maxColumnNumber-1 != minColumnNumber {
		log.Fatalln(minColumnNumber, maxColumnNumber)
	}
	return minColumnNumber
}
