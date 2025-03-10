## Golang `Enum` to `String` helper library

A simple 1-purpose library that converts an enum, specifically those using `iota` to return a string instance.

Since I do alot of types and I needed the string value, having to repeat myself becomes a chore, hence why I wrote this very simple library.

### Assumptions

- Only supports `uint8`, `int8`, `int` types.
- Numbers are sequential, any gaps expects the user to fill an empty string by the user, as the string slice doesn't handle gaps.
- Since `iota` values typically start at `1`, `0` is reserved for `unknown`, allowing the user to process unknown values passed using a single string. 

**Note on index** - it is highly recommended to support the `unknown` string because this library checks if the number given is less than 0 or greater than the length of string given, at which point it defaults to `0`, thus returning `unknown` for any odd input.

### Examples

```go
import (
	"fmt"

	"github.com/svicknesh/enum2str"
)

type Algorithm uint8

const (
	EC256 Algorithm = iota + 1
	EC384
	EC521
)

const (
	BadAlgo Algorithm = 100
)

type KeyType int8

const (
	KeyTypeRSA KeyType = iota + 1
	KeyTypeDSA
)

type Response int

const (
	ResponseSuccess Response = iota + 1
	ResponseFail
)

func main() {

	// uint8 test
	fmt.Println(enum2str.String(EC256, "unknown", "ECDSA-256", "ECDSA-384", "ECDSA-521"))
	fmt.Println(enum2str.String(BadAlgo, "unknown", "ECDSA-256", "ECDSA-384", "ECDSA-521"))

	// int8 test
	fmt.Println(enum2str.String(KeyTypeDSA, "unknown", "RSA", "DSA"))

	// int test
	fmt.Println(enum2str.String(ResponseSuccess, "unknown", "success", "fail"))

}
```
