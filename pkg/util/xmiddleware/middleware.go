// Created at 11/30/2021 11:22 PM
// Developer: trungnq2710 (trungnq2710@gmail.com)

package xmiddleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/jwt/v3"
	"github.com/trungnq2710/go-pkg/pkg/util/xrest"
)

type HasPermitFunc func(c *fiber.Ctx) bool

func Protected(tokenSecret []byte, hasPermit HasPermitFunc) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:     tokenSecret,
		ErrorHandler:   errorHandler,
		SuccessHandler: successHandler(hasPermit),
	})
}

func errorHandler(c *fiber.Ctx, err error) error {
	log.Printf("[%s] - [%s] => %s - err: %s ", c.Method(), c.IP(), c.OriginalURL(), err.Error())

	if err.Error() == "Missing or malformed JWT" {
		return xrest.ErrorWithStatus(c, xrest.STATUS_BAD_REQUEST, "missing or malformed jwt")
	}
	return xrest.ErrorWithStatus(c, xrest.STATUS_TOKEN_INVALID, "invalid or expired jwt")
}

func successHandler(hasPermit HasPermitFunc) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if hasPermit == nil || hasPermit(c) {
			return c.Next()
		}

		return xrest.ErrorWithStatus(c, xrest.STATUS_PERMIT_DENIED, "permit denied")
	}
}
