package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// type Passport struct {
// 	Byr int
// 	Iyr int
// 	Eyr int
// 	Hgt string
// 	Hcl string
// 	Ecl string
// 	Pid int
// 	Cid int
// }

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
	passports := createPassports(inputs)
	validPassportCount := countValidPassports(passports)
	fmt.Printf("Valid passports: %d\n", validPassportCount)
}

func countValidPassports(passports []map[string]string) int {
	var count int
	for _, pp := range passports {
		if pp["byr"] != "" && pp["iyr"] != "" && pp["eyr"] != "" && pp["hgt"] != "" && pp["hcl"] != "" && pp["ecl"] != "" && pp["pid"] != "" {
			count++
		}
	}
	return count
}

func extractValuePairs(line string, passport map[string]string) map[string]string {
	pairs := strings.Split(line, " ")
	for _, pair := range pairs {
		passport[pair[0:3]] = pair[4:]
	}
	return passport
}

func createPassports(inputs []string) []map[string]string {
	passportlist := make([]map[string]string, 0)
	passport := make(map[string]string)
	for i, line := range inputs {
		if line != "" {
			passport = extractValuePairs(line, passport)
		}
		if i+1 == len(inputs) || line == "" {
			passportlist = append(passportlist, passport)
		}
	}
	return passportlist
}
