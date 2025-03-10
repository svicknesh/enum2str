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

func TestFull(t *testing.T) {

	fmt.Println(enum2str.String(BadAlgo, "unknown", "ECDSA-256", "ECDSA-384", "ECDSA-521"))

}
