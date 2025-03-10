package enum2str_test

import (
	"fmt"
	"testing"

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

func TestFull(t *testing.T) {

	// uint8 test
	fmt.Println(enum2str.String(EC256, "unknown", "ECDSA-256", "ECDSA-384", "ECDSA-521"))
	fmt.Println(enum2str.String(BadAlgo, "unknown", "ECDSA-256", "ECDSA-384", "ECDSA-521"))

	// int8 test
	fmt.Println(enum2str.String(KeyTypeDSA, "unknown", "RSA", "DSA"))

	// int test
	fmt.Println(enum2str.String(ResponseSuccess, "unknown", "success", "fail"))

}
