package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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
	re := regexp.MustCompile(`(?m)([0-9]*)-([0-9]*) ([a-z]): (.*)`)
	var valid int
	for _, input := range inputs {
		//fmt.Println(input)
		splitvalues := re.FindStringSubmatch(input)
		//fmt.Println(splitvalues[1])
		minimum, _ := strconv.Atoi(splitvalues[1])
		maximum, _ := strconv.Atoi(splitvalues[2])
		letter := splitvalues[3]
		password := splitvalues[4]
		fmt.Printf("minimum: %d, maximum %d, letter %s, password %s\n", minimum, maximum, letter, password)
		rePolicyMatch := regexp.MustCompile(letter)
		matches := rePolicyMatch.FindAllStringIndex(password, -1)
		fmt.Println(matches)
		// if len(matches) <= maximum && len(matches) >= minimum {
		// 	valid++
		// }
		// fmt.Println(matches[0][1])
		// fmt.Println(matches[1][1])
		var minfound, maxfound bool
		for _, match := range matches {
			if match[1] == minimum {
				minfound = true
			}
			if match[1] == maximum {
				maxfound = true
			}
		}
		if (minfound || maxfound) && !(minfound && maxfound) {
			valid++
		}
	}
	fmt.Println(valid)
}
