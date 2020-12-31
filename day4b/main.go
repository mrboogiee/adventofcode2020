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
	passport := make(map[string]string)
	passports := make([]map[string]string, 0)
	for i, line := range inputs {
		if line != "" {
			fields := strings.Split(line, " ")
			for _, field := range fields {
				keyvalue := strings.Split(field, ":")
				passport[keyvalue[0]] = keyvalue[1]
			}
		}
		if line == "" || i+1 == len(inputs) {
			passports = append(passports, passport)
			passport = map[string]string{}
		}
	}
	var count int
	for _, passport := range passports {
		var validPassport = true
		if passport["byr"] != "" && passport["iyr"] != "" && passport["eyr"] != "" && passport["hgt"] != "" && passport["hcl"] != "" && passport["ecl"] != "" && passport["pid"] != "" {
			for field, value := range passport {
				if field != "cid" {
					if !validateField(field, value) {
						validPassport = false
					}
				}
			}
			if validPassport {
				count++
			}
		}

	}
	fmt.Println(count)
}

func validateField(field, value string) bool {
	switch field {
	case "byr":
		byr, err := strconv.Atoi(value)
		if err == nil {
			if 1920 <= byr && byr <= 2002 {
				return true
			}
		}
	case "iyr":
		iyr, err := strconv.Atoi(value)
		if err == nil {
			if 2010 <= iyr && iyr <= 2020 {
				return true
			}
		}
	case "eyr":
		eyr, err := strconv.Atoi(value)
		if err == nil {
			if 2020 <= eyr && eyr <= 2030 {
				return true
			}
		}
	case "hgt":
		var re = regexp.MustCompile(`(?m)([0-9]{3}) ?(cm)|([0-9]{2}) ?(in)`)
		//for i, match := range re.FindAllString(value, -1) {
		for _, match := range re.FindAllStringSubmatch(value, -1) {
			if match[2] == "cm" {
				// fmt.Println(match[1], match[2], "found at index", i)
				length, err := strconv.Atoi(match[1])
				if err == nil {
					if 150 <= length && length <= 193 {
						return true
					}
				}
			} else if match[4] == "in" {
				// fmt.Println(match[3], match[4], "found at index", i)
				length, err := strconv.Atoi(match[3])
				if err == nil {
					if 59 <= length && length <= 76 {
						return true
					}
				}
			}
		}
	case "hcl":
		var re = regexp.MustCompile(`^#[0-9a-fA-F]{6}$`)
		if re.MatchString(value) {
			return true
		}
		// }
	case "ecl":
		if value == "amb" || value == "blu" || value == "brn" || value == "gry" || value == "grn" || value == "hzl" || value == "oth" {
			return true
		}
	case "pid":
		_, err := strconv.Atoi(value)
		if err == nil {
			if len(value) == 9 {
				return true
			}
		}
	}
	return false
}
