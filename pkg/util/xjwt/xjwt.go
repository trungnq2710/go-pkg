// Created at 11/30/2021 10:59 PM
// Developer: trungnq2710 (trungnq2710@gmail.com)

package xjwt

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

var (
	tokenExpired     = errors.New("token is timed out, please log in again")
	tokenInvalid     = errors.New("token has been invalidated")
	tokenNotValidYet = errors.New("token not active yet")
	tokenMalformed   = errors.New("that's not even a token")
	tokenUnknown     = errors.New("couldn't handle this token")
)

type Claims struct {
	UID int64
	jwt.RegisteredClaims
}

func BuildClaims(uid, ttl int64) Claims {
	now := time.Now()
	return Claims{
		UID: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(ttl*24) * time.Hour)), //Expiration time
			IssuedAt:  jwt.NewNumericDate(now),                                        //Issuing time
			NotBefore: jwt.NewNumericDate(now),                                        //Begin Effective time
		}}
}

func CreateToken(uid, ttl int64, signed []byte) (string, int64, error) {
	claims := BuildClaims(uid, ttl)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signed)
	return tokenString, claims.ExpiresAt.Time.Unix(), err
}

func VerifyToken(tokensString string, tokenSecret []byte) error {
	_, err := getClaimFromToken(tokensString, tokenSecret)
	return err
}

func ParseToken(tokensString string, tokenSecret []byte) (claims *Claims, err error) {
	return getClaimFromToken(tokensString, tokenSecret)
}

func getClaimFromToken(tokensString string, tokenSecret []byte) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokensString, &Claims{}, secret(tokenSecret))
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, tokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, tokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, tokenNotValidYet
			} else {
				return nil, tokenUnknown
			}
		}
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func secret(tokenSecret []byte) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return tokenSecret, nil
	}
}

func GetUIDFromFiberCtx(c *fiber.Ctx) (int64, int, error) {
	tmp := c.Locals("user")

	token, ok := tmp.(*jwt.Token)
	if !ok {
		return 0, fiber.StatusInternalServerError, errors.New("internal token type invalid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fiber.StatusInternalServerError, errors.New("internal claims type invalid")
	}

	iUID := claims["UID"]
	if iUID == nil {
		return 0, fiber.StatusUnauthorized, errors.New("invalid or expired jwt")
	}

	uid, err := convertInterfaceToInt64(iUID)
	if err != nil {
		return 0, fiber.StatusInternalServerError, errors.New("internal uid type invalid")
	}

	return uid, fiber.StatusOK, nil
}

func convertInterfaceToInt64(t interface{}) (int64, error) {
	switch t := t.(type) {
	case int64:
		return t, nil
	case int:
		return int64(t), nil
	case string:
		return strconv.ParseInt(t, 10, 64)
	case float64:
		return int64(t), nil
	default:
		return 0, errors.Errorf("type %T not supported", t)
	}
}
