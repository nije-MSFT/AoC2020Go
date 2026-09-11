package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Passport struct {
	byr, iyr, eyr, hgt, hcl, ecl, pid, cid string
}

func main() {
	file, _ := os.Open("./4/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var passport Passport
	var validPassports = 0

	for scanner.Scan() {
		if scanner.Text() == "" {
			if passport.byr != "" && passport.iyr != "" && passport.eyr != "" && passport.hgt != "" && passport.hcl != "" && passport.ecl != "" && passport.pid != "" {
				validPassports++
			}
			passport = Passport{}
			continue
		}
		splitLine := strings.Split(scanner.Text(), " ")

		for _, token := range splitLine {
			keyValue := strings.Split(token, ":")

			switch keyValue[0] {
			case "byr":
				passport.byr = keyValue[1]
			case "iyr":
				passport.iyr = keyValue[1]
			case "eyr":
				passport.eyr = keyValue[1]
			case "hgt":
				passport.hgt = keyValue[1]
			case "hcl":
				passport.hcl = keyValue[1]
			case "ecl":
				passport.ecl = keyValue[1]
			case "pid":
				passport.pid = keyValue[1]
			case "cid":
				passport.cid = keyValue[1]
			}
		}
	}

	if passport.byr != "" && passport.iyr != "" && passport.eyr != "" && passport.hgt != "" && passport.hcl != "" && passport.ecl != "" && passport.pid != "" {
		validPassports++
	}

	fmt.Println("Valid Passports:", validPassports)
}
