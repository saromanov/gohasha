# gohasha
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

## Licence
MIT