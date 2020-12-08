package day4

import (
	"bytes"
	"regexp"
	"strconv"
)

// Passport each of the passport records
type Passport struct {
	Fields map[string]string
}

var requiredFields = [...]string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
	//"cid",
}

var rePassport *regexp.Regexp

func init() {
	// Matches one line
	rePassport = regexp.MustCompile(`([^:\s]+):([^:\s]+)`)
}

// ParsePassport parses a single passport
func ParsePassport(input []byte) Passport {
	fieldPairs := rePassport.FindAllSubmatch(input, -1)
	fields := make(map[string]string)
	for _, pair := range fieldPairs {
		fields[string(pair[1])] = string(pair[2])
	}
	return Passport{Fields: fields}
}

// HasRequiredFields test if the passport has the right fields
func (passport *Passport) HasRequiredFields() bool {
	for _, field := range requiredFields {
		if _, ok := passport.Fields[field]; !ok {
			return false
		}
	}
	return true
}

func (passport *Passport) validateField(field string, pred func(string) bool) bool {
	f, ok := passport.Fields[field]
	if !ok {
		return false
	}
	return pred(f)
}

// Validate test if the passport has the right fields AND the fields are correct
func (passport *Passport) Validate() bool {
	intInRange := func(low int, high int) func(string) bool {
		return func(field string) bool {
			num, err := strconv.Atoi(field)
			if err != nil {
				return false
			}
			return (low <= num) && (num <= high)
		}
	}

	if !passport.validateField("byr", intInRange(1920, 2002)) {
		return false
	}

	if !passport.validateField("iyr", intInRange(2010, 2020)) {
		return false
	}

	if !passport.validateField("eyr", intInRange(2020, 2030)) {
		return false
	}

	if !passport.validateField("hgt", func(field string) bool {
		units := field[len(field)-2:]
		if units == "cm" {
			return intInRange(150, 193)(field[:len(field)-2])
		}

		if units == "in" {
			return intInRange(59, 76)(field[:len(field)-2])
		}

		return false
	}) {
		return false
	}

	if !passport.validateField("hcl", func(field string) bool {
		ok, _ := regexp.Match(`^#[0-9a-f]{6}$`, []byte(field))
		return ok
	}) {
		return false
	}

	if !passport.validateField("ecl", func(field string) bool {
		ok, _ := regexp.Match(`^(amb|blu|brn|gry|grn|hzl|oth)$`, []byte(field))
		return ok
	}) {
		return false
	}

	if !passport.validateField("pid", func(field string) bool {
		ok, _ := regexp.Match(`^\d{9}$`, []byte(field))
		return ok
	}) {
		return false
	}

	return true
}

// ParsePassports loads a slice of passport from bytes
func ParsePassports(input []byte) []Passport {
	lines := bytes.Split(input, []byte("\n\n"))
	passports := make([]Passport, len(lines))
	for i, line := range lines {
		passports[i] = ParsePassport(line)
	}
	return passports
}
