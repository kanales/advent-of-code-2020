package days

import (
	"bytes"
	"reflect"
	"testing"
)

// func TestParseRecord(t *testing.T) {
// 	input := []byte("1-3 a: abcde")
// 	expect := PasswordRecord{Low: 1, High: 3, Value: 'a', Password: []byte("abcde")}
// 	got := ParseRecord(input)
// 	if expect.Low != got.Low {
// 		t.Errorf(".Low = %v; want %v", got.Low, expect.Low)
// 	}
// 	if expect.High != got.High {
// 		t.Errorf(".High = %v; want %v", got.High, expect.High)
// 	}
// 	if expect.Value != got.Value {
// 		t.Errorf(".Value = %v; want %v", got.Value, expect.Value)
// 	}
// 	if !bytes.Equal(expect.Password, got.Password) {
// 		t.Errorf(".Password = %v; want %v", got.Password, expect.Password)
// 	}
// }

func TestParseRecords(t *testing.T) {
	input := []byte("1-3 a: abcde\n1-3 b: cdefg\n2-9 c: ccccccccc\n")
	expect := []PasswordRecord{
		{Low: 1, High: 3, Value: 'a', Password: []byte("abcde")},
		{Low: 1, High: 3, Value: 'b', Password: []byte("cdefg")},
		{Low: 2, High: 9, Value: 'c', Password: []byte("ccccccccc")},
	}
	got := ParseRecords(input)
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
		if !bytes.Equal(record2.Password, record.Password) {
			t.Errorf(".Password = %v; want %v", record2.Password, record.Password)
		}
	}
}

func TestIsPasswordCorrect(t *testing.T) {
	input := []PasswordRecord{
		{Low: 1, High: 3, Value: 'a', Password: []byte("abcde")},
		{Low: 1, High: 3, Value: 'b', Password: []byte("cdefg")},
		{Low: 2, High: 9, Value: 'c', Password: []byte("ccccccccc")},
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
		{Low: 1, High: 3, Value: 'a', Password: []byte("abcde")},
		{Low: 1, High: 3, Value: 'b', Password: []byte("cdefg")},
		{Low: 2, High: 9, Value: 'c', Password: []byte("ccccccccc")},
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
		{Low: 1, High: 3, Value: 'a', Password: []byte("abcde")},
		{Low: 1, High: 3, Value: 'b', Password: []byte("cdefg")},
		{Low: 2, High: 9, Value: 'c', Password: []byte("ccccccccc")},
	}
	expect := 2
	got := CountCorrectPasswords(input, IsPasswordCorrect1)
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("CountCorrectPasswords(%v)= %v; want %v", input, got, expect)
	}
}
