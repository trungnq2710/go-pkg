// Created at 11/30/2021 10:18 PM
// Developer: trungnq2710 (trungnq2710@gmail.com)

package xpassword

import "golang.org/x/crypto/bcrypt"

func Generate(raw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), 10)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func Verify(hash string, raw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(raw))
}
