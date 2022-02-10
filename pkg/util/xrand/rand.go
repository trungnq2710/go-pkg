// Created at 12/27/2021 11:14 AM
// Developer: trungnq2710 (trungnq2710@gmail.com)

package xrand

import (
	"math/rand"
	"time"
)

func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func RandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min+1)
}

func RandomStringNumber(len int) string {
	rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(48, 57))
	}
	return string(bytes)
}
