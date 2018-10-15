package util

import (
	"testing"
	"fmt"
)

func TestGetRandomString(t *testing.T) {
	fmt.Printf("GetRandomString result:%s\n", GetRandomString(8))
}

func TestMD5(t *testing.T) {
	fmt.Printf("MD5 result:%s\n", MD5("1g1hlg42" + "#123@HGD"))
}