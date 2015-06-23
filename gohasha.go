package gohasha

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"hash/crc64"
	"io/ioutil"
	"io"
	"bytes"
)

type GohashaOptions struct {
	//path to file for hashing
	Filepath string
	//Data from hashing
	Data string
	//Reader buffer
	BufferReader io.Reader
	
	//Now, algorithms can be md5 or crc32
	Algorithm string
}

//GoHasha return hash in string from string data or file
func GoHasha(opt *GohashaOptions) (string, error) {
	if opt.Data != "" {
		return crypt(opt.Data, opt.Algorithm), nil
	}
	if opt.Filepath != "" {
		data, err := ioutil.ReadFile(opt.Filepath)
		if err != nil {
			return "", err
		} else {
			return crypt(string(data), opt.Algorithm), nil
		}
	}

	if opt.BufferReader != nil {
		buff := new(bytes.Buffer)
		buff.ReadFrom(opt.BufferReader)
		return crypt(buff.String(), opt.Algorithm), nil
	}
	return "", errors.New("Can't find data for hashing")
}

func crypt(data string, algorithm string) string {
	if algorithm == "crc64" {
		return crc64run(data)
	} else {
		return md5run(data)
	}
}

func md5run(data string) string {
	hasher := md5.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}

func crc64run(data string) string {
	hasher := crc64.New(crc64.MakeTable(crc64.ISO))
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}
