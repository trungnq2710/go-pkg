// Created at 11/18/2021 3:05 PM
// Developer: trungnq2710 (trungnq2710@gmail.com)

package xstring

import (
	"crypto/rand"
	"math/big"
	"reflect"
	"runtime"
	"unicode"
)

func FunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func IsUniqueCharacters(chars []byte) bool {
	seen := make(map[byte]bool)
	for _, char := range chars {
		if _, ok := seen[char]; ok {
			return false
		}
		seen[char] = true
	}
	return true
}

func GenerateRandomString(length int) (string, error) {
	var result string
	for len(result) < length {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(127)))
		if err != nil {
			return "", err
		}
		n := num.Int64()
		if unicode.IsLetter(rune(n)) {
			result += string(n)
		}
	}
	return result, nil
}
