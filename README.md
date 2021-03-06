# gohasha [![Build Status](https://travis-ci.org/saromanov/gohasha.svg?branch=master)](https://travis-ci.org/saromanov/gohasha)

Inspired by [hasha](https://github.com/sindresorhus/hasha)

## Install
``` go get github.com/saromanov/gohasha ```
## Usage
```go
import
(
	"gohasha"
	"fmt"
)

func main() {
	hashstr, _ := gohasha.GoHasha(&gohasha.GohashaOptions{Data: "Value"})
	fmt.Println(hashstr) //returns hash for Value
}
```
or read from file
```go
result, _ := gohasha.GoHasha(&gohasha.GohashaOptions{Filepath: "testfile", Algorithm: "crc64"})
fmt.Println(result) //returns hash for data from testfile with crc64
```

## API
### gohasha.GohashaOptions
* Data - data for hashing
* Filepath -  path to file for hashing
* BufferReader - Buffer reader for hashing
* Object - any object for hashing
* Algorithm - md5, crc32 or adler32
* HashingFunc - your hashing function
Example:
```go
hashstr, _ := gohasha.GoHasha(&gohasha.GohashaOptions{Data: "Value",
HashingFunc: function(str string)string{
    //Just as example
	return str
}})
```

### gohasha.GoHasha(opt GohashaOptions)
Return hashed data


## Licence
MIT