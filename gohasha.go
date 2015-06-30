package gohasha

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"hash/adler32"
	"hash/crc64"
	"io"
	"io/ioutil"
	"net/http"
)

type GohashaOptions struct {
	//path to file for hashing
	Filepath string
	//Data from hashing
	Data string
	//Reader buffer
	BufferReader io.Reader

	//Get data from webpage
	Webpage string

	//Now, algorithms can be md5 or crc32
	Algorithm string

	//Your hashing function
	HashingFunc func(string) string
}

//GoHasha return hash in string from string data or file
func GoHasha(opt *GohashaOptions) (string, error) {
	data := opt.Data
	/*if opt.Data != "" {
		return crypt(opt.Data, opt.Algorithm), nil
	}*/
	if opt.Filepath != "" {
		datab, err := ioutil.ReadFile(opt.Filepath)
		if err != nil {
			return "", err
		} else {
			data = string(datab)
		}
	}

	if opt.BufferReader != nil {
		buff := new(bytes.Buffer)
		buff.ReadFrom(opt.BufferReader)
		data = buff.String()
	}

	if opt.Webpage != "" {
		resp, err := http.Get(opt.Webpage)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		result, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		data = string(result)
	}

	if data == "" {
		return "", errors.New("Can't find data for hashing")
	} else {
		return crypt(data, opt.Algorithm, opt.HashingFunc), nil
	}
}

//Choose algorithms for crypt
func crypt(data string, algorithm string, hashfunc func(string) string) string {
	if hashfunc != nil {
		return hashfunc(data)
	}
	if algorithm == "crc64" {
		return crc64run(data)
	} else if algorithm == "adler32" {
		return fmt.Sprint(adler32run(data))
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

func adler32run(data string) uint32 {
	return adler32.Checksum([]byte(data))
}
