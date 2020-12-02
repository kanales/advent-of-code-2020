package days

import (
	"reflect"
	"testing"
)

func TestParsePassword(t *testing.T) {
	input := "1-3 a: abcde\n1-3 b: cdefg\n2-9 c: ccccccccc\n"
	expect := []PasswordRecord{
		{Low: 1, High: 3, Value: "a", Password: "abcde"},
		{Low: 1, High: 3, Value: "b", Password: "cdefg"},
		{Low: 2, High: 9, Value: "c", Password: "ccccccccc"},
	}
	got := ParsePasswords(input)
	for i, record := range expect {
		record2 := got[i]
		if record2.Low != record.Low {
			t.Errorf(".Low = %v; want %v", record2.Low, record.Low)
		}
		if record2.High != record.High {
			t.Errorf(".High = %v; want %v", record2.High, record.High)
		}
		if record2.Value != record.Value {
			t.Errorf(".Value = %v; want %v", record2.Value, record.Value)
		}
		if record2.Password != record.Password {
			t.Errorf(".Password = %v; want %v", record2.Password, record.Password)
		}
	}
}

func TestIsPasswordCorrect(t *testing.T) {
	input := []PasswordRecord{
		{Low: 1, High: 3, Value: "a", Password: "abcde"},
		{Low: 1, High: 3, Value: "b", Password: "cdefg"},
		{Low: 1, High: 9, Value: "c", Password: "ccccccccc"},
	}

	if !IsPasswordCorrect1(input[0]) {
		t.Errorf("IsPasswordCorrect(%v)= false; want true", input[0])
	}

	if IsPasswordCorrect1(input[1]) {
		t.Errorf("IsPasswordCorrect(%v)= true; want false", input[1])
	}

	if !IsPasswordCorrect1(input[2]) {
		t.Errorf("IsPasswordCorrect(%v)= false; want true", input[2])
	}
}

func TestIsPasswordCorrect2(t *testing.T) {
	input := []PasswordRecord{
		{Low: 1, High: 3, Value: "a", Password: "abcde"},
		{Low: 1, High: 3, Value: "b", Password: "cdefg"},
		{Low: 1, High: 9, Value: "c", Password: "ccccccccc"},
	}

	if !IsPasswordCorrect2(input[0]) {
		t.Errorf("IsPasswordCorrect(%v)= false; want true", input[0])
	}

	if IsPasswordCorrect2(input[1]) {
		t.Errorf("IsPasswordCorrect(%v)= true; want false", input[1])
	}

	if IsPasswordCorrect2(input[2]) {
		t.Errorf("IsPasswordCorrect(%v)= true; want false", input[2])
	}
}

func TestCountCorrectPasswords(t *testing.T) {
	input := []PasswordRecord{
		{Low: 1, High: 3, Value: "a", Password: "abcde"},
		{Low: 1, High: 3, Value: "b", Password: "cdefg"},
		{Low: 1, High: 9, Value: "c", Password: "ccccccccc"},
	}
	expect := 2
	got := CountCorrectPasswords(input, IsPasswordCorrect1)
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("CountCorrectPasswords(%v)= %v; want %v", input, got, expect)
	}
}
