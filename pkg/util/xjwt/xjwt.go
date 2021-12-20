// Created at 11/30/2021 10:59 PM
// Developer: trungnq2710 (trungnq2710@gmail.com)

package xjwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"

	"gitlab.com/jplatform/jengine/constant"
)

var (
	TokenExpired     = errors.New("token is timed out, please log in again")
	TokenInvalid     = errors.New("token has been invalidated")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenUnknown     = errors.New("couldn't handle this token")
)

type Claims struct {
	UID int32
	jwt.RegisteredClaims
}

func BuildClaims(uid int32, ttl int64) Claims {
	now := time.Now()
	return Claims{
		UID: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(ttl*24) * time.Hour)), //Expiration time
			IssuedAt:  jwt.NewNumericDate(now),                                        //Issuing time
			NotBefore: jwt.NewNumericDate(now),                                        //Begin Effective time
		}}
}

func CreateToken(userID int32, ttl int64, signed []byte) (string, int64, error) {
	claims := BuildClaims(userID, ttl)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signed)
	return tokenString, claims.ExpiresAt.Time.Unix(), err
}

func VerifyToken(tokensString string) error {
	_, err := getClaimFromToken(tokensString)
	return err
}

func ParseToken(tokensString string) (claims *Claims, err error) {
	return getClaimFromToken(tokensString)
}

func getClaimFromToken(tokensString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokensString, &Claims{}, secret())
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenUnknown
			}
		}
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(constant.TokenSecret), nil
	}
}