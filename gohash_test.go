package gohasha

import
(
	"testing"
)

const
(
	MD5LENGTH = 32
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

func TestDataCrc64 (t *testing.T) {
	result, err := GoHasha(&GohashaOptions{Data: "foobar", Algorithm:"crc64"})
	if err != nil {
		t.Errorf("found error")
	}

	if len(result) != CRC64LENGTH {
		t.Errorf("%d not match %d", len(result),CRC64LENGTH)
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

