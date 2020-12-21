package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type passport struct {
	Byr int
	Iyr int
	Eyr int
	Hgt string
	Hcl string
	Ecl string
	Pid int
	Cid int
}

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
	passports := createPassports(inputs)
	validPassportCount := countValidPassports(passports)
	fmt.Printf("Valid passports: %d\n", validPassportCount)
}

func countValidPassports(passports []passport) int {
	var count int
	for _, pp := range passports {
		fmt.Printf("byr: %d iry: %d, eyr: %d, hgt: %s, hcl: %s, ecl: %s, pid: %d, cid: %d\n", pp.Byr, pp.Iyr, pp.Eyr, pp.Hgt, pp.Hcl, pp.Ecl, pp.Pid, pp.Cid)
		if pp.Byr != 0 && pp.Iyr != 0 && pp.Eyr != 0 && pp.Hgt != "" && pp.Hcl != "" && pp.Ecl != "" && pp.Pid != 0 {
			count++
		}
	}
	return count
}

func createPassports(inputs []string) []passport {
	var passportlist []passport
	var Passport passport
	//var lf = regexp.MustCompile(`^\s*$`)
	var reByr = regexp.MustCompile(`(?m)byr:([0-9]*)`)
	var reIyr = regexp.MustCompile(`(?m)iyr:([0-9]*)`)
	var reEyr = regexp.MustCompile(`(?m)eyr:([0-9]*)`)
	var reHgt = regexp.MustCompile(`(?m)hgt:([0-9]*cm)`)
	var reHcl = regexp.MustCompile(`(?m)hcl:([#0-9a-zA-Z]*)`)
	var reEcl = regexp.MustCompile(`(?m)ecl:([0-9a-zA-Z]*)`)
	var rePid = regexp.MustCompile(`(?m)pid:([0-9]*)`)
	var reCid = regexp.MustCompile(`(?m)cid:([0-9]*)`)

	for i, line := range inputs {
		if line == "" {
			fmt.Println("New passport")
			fmt.Println(passportlist)
			passportlist = append(passportlist, Passport)
			Passport = passport{}
		} else {
			byr := reByr.FindStringSubmatch(line)
			if len(byr) > 0 {
				byrint, _ := strconv.Atoi(byr[1])
				Passport.Byr = byrint
			}
			iyr := reIyr.FindStringSubmatch(line)
			if len(iyr) > 0 {
				iyrint, _ := strconv.Atoi(iyr[1])
				Passport.Iyr = iyrint
			}
			eyr := reEyr.FindStringSubmatch(line)
			if len(eyr) > 0 {
				eyrint, _ := strconv.Atoi(eyr[1])
				Passport.Eyr = eyrint
			}
			hgt := reHgt.FindStringSubmatch(line)
			if len(hgt) > 0 {
				Passport.Hgt = hgt[1]
			}
			hcl := reHcl.FindStringSubmatch(line)
			if len(hcl) > 0 {
				Passport.Hcl = hcl[1]
			}
			ecl := reEcl.FindStringSubmatch(line)
			if len(ecl) > 0 {
				Passport.Ecl = ecl[1]
			}
			pid := rePid.FindStringSubmatch(line)
			if len(pid) > 0 {
				pidint, _ := strconv.Atoi(pid[1])
				Passport.Pid = pidint
			}
			cid := reCid.FindStringSubmatch(line)
			if len(cid) > 0 {
				cidint, _ := strconv.Atoi(cid[1])
				Passport.Cid = cidint
			}
		}
		if i+1 == len(inputs) {
			passportlist = append(passportlist, Passport)
		}
	}
	return passportlist
}
