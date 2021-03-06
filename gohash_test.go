package gohasha

import (
	"testing"
)

const (
	MD5LENGTH   = 32
	CRC64LENGTH = 16
)

func TestDataMd5(t *testing.T) {
	result, err := GoHasha(&GohashaOptions{Data: "foobar"})
	if err != nil {
		t.Errorf("Found error")
	}

	if len(result) != MD5LENGTH {
		t.Errorf("%d not match %d", len(result), MD5LENGTH)
	}
}

func TestDataCrc64(t *testing.T) {
	result, err := GoHasha(&GohashaOptions{Data: "foobar", Algorithm: "crc64"})
	if err != nil {
		t.Errorf("found error")
	}

	if len(result) != CRC64LENGTH {
		t.Errorf("%d not match %d", len(result), CRC64LENGTH)
	}
}

func TestDataObject(t *testing.T) {
	value := map[string]string{"Seven": "days"}
	result, err := GoHasha(&GohashaOptions{Object:value, Algorithm: "crc64"})
	if err != nil {
		t.Errorf("found error")
	}

	if len(result) != CRC64LENGTH {
		t.Errorf("%d not match %d", len(result), CRC64LENGTH)
	}
}

func TestUserFunction(t *testing.T) {
	result, err := GoHasha(&GohashaOptions{Data: "foobar", HashingFunc: func(str string) string{
		return str
	}})
	if err != nil {
		t.Errorf("found error")
	}

	if result != "foobar" {
		t.Errorf("%s not match %s", result, "foobar")
	}

}

func TestReadFromFileMd5(t *testing.T) {
	result, _ := GoHasha(&GohashaOptions{Filepath: "gohasha.go"})
	if len(result) != MD5LENGTH {
		t.Errorf("%d not match %d", len(result), MD5LENGTH)
	}
}

func TestReadFromFileCrc64(t *testing.T) {
	result, _ := GoHasha(&GohashaOptions{Filepath: "gohasha.go", Algorithm: "crc64"})
	if len(result) != CRC64LENGTH {
		t.Errorf("%d not match %d", len(result), CRC64LENGTH)
	}
}

func TestReadFromWebPage(t *testing.T) {
	result, _ := GoHasha(&GohashaOptions{Webpage: "https://golang.org"})
	if len(result) != MD5LENGTH {
		t.Errorf("%d not match %d", len(result), CRC64LENGTH)
	}
}
