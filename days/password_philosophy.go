package days

import (
	"strconv"
	"strings"
)

// PasswordRecord a line of the Password Philosophy (day 2) password file
type PasswordRecord struct {
	Low      int
	High     int
	Value    string
	Password string
}

// ParsePasswords generates a slice of password records from input
func ParsePasswords(data string) []PasswordRecord {
	lines := strings.Split(data, "\n")
	lines = lines[:len(lines)-1]

	records := make([]PasswordRecord, len(lines))
	for i, line := range lines {
		sublines := strings.SplitN(line, ": ", 2)
		check := sublines[0]
		pass := sublines[1]
		sublines = strings.SplitN(check, " ", 2)
		value := sublines[1]
		sublines = strings.SplitN(sublines[0], "-", 2)
		low, _ := strconv.Atoi(sublines[0])
		high, _ := strconv.Atoi(sublines[1])
		records[i] = PasswordRecord{Low: low, High: high, Value: value, Password: pass}
	}
	return records
}

// IsPasswordCorrect1 checks password for the past job
func IsPasswordCorrect1(record PasswordRecord) bool {
	counts := strings.Count(record.Password, record.Value)
	return record.Low <= counts && record.High >= counts
}

// IsPasswordCorrect2 checks password for the current job
func IsPasswordCorrect2(record PasswordRecord) bool {
	lowOK := record.Value == string(record.Password[record.Low-1])
	highOK := record.Value == string(record.Password[record.High-1])
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
