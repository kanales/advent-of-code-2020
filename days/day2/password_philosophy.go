package day2

import (
	"bytes"
	"regexp"
	"strconv"
)

// PasswordRecord a line of the Password Philosophy (day 2) password file
type PasswordRecord struct {
	Low      int
	High     int
	Value    byte
	Password []byte
}

var rePassword *regexp.Regexp

func init() {
	// Example: `1-3 a: abcde`
	rePassword = regexp.MustCompile(`(?m)^(\d+)-(\d+) ([[:alpha:]]): ([[:alpha:]]+)$`)
}

// ParseRecords generates a slice of password records from input
func ParseRecords(data []byte) []PasswordRecord {
	matches := rePassword.FindAllSubmatch(data, -1)

	records := make([]PasswordRecord, len(matches))
	for i, match := range matches {
		low, _ := strconv.Atoi(string(match[1]))
		high, _ := strconv.Atoi(string(match[2]))
		records[i] = PasswordRecord{
			Low:      low,
			High:     high,
			Value:    match[3][0],
			Password: match[4]}
	}
	return records
}

// IsPasswordCorrect1 checks password for the past job
func IsPasswordCorrect1(record PasswordRecord) bool {
	counts := bytes.Count(record.Password, []byte{record.Value})
	return record.Low <= counts && record.High >= counts
}

// IsPasswordCorrect2 checks password for the current job
func IsPasswordCorrect2(record PasswordRecord) bool {
	lowOK := record.Value == record.Password[record.Low-1]
	highOK := record.Value == record.Password[record.High-1]
	return lowOK != highOK
}

// CountCorrectPasswords returns how many passwords pass with the criterion
func CountCorrectPasswords(records []PasswordRecord, criterion func(PasswordRecord) bool) int {
	counter := 0
	for _, record := range records {
		if criterion(record) {
			counter++
		}
	}
	return counter
}
