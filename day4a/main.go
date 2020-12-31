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
	for _,passport := range passports{
		if passport["byr"]!="" && passport["iyr"] !="" && passport["eyr"] != "" && passport["hgt"]!="" && passport["hcl"]!=""&&passport["ecl"]!=""&&passport["pid"]!=""{
			count++
		}
	}
	fmt.Println(count)
}
