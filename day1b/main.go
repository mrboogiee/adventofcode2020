package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var inputs []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		input, _ := strconv.Atoi(scanner.Text())
		inputs = append(inputs, input)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, primary := range inputs {
		for _, secondary := range inputs {
			for _, tertiary := range inputs {
				if primary+secondary+tertiary == 2020 {
					fmt.Println(primary * secondary * tertiary)
					os.Exit(0)
				}
			}
		}
	}
}
