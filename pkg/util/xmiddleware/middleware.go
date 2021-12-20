// Created at 11/30/2021 11:22 PM
// Developer: trungnq2710 (trungnq2710@gmail.com)

package xmiddleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/jwt/v2"

	"github.com/trungnq2710/go-pkg/pkg/util/xrest"
)

func Protected(tokenSecret []byte) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:     tokenSecret,
		ErrorHandler:   errorHandler,
		SuccessHandler: successHandler(),
	})
}

func errorHandler(c *fiber.Ctx, err error) error {
	log.Printf("[%s] - [%s] => %s - err: %s ", c.Method(), c.IP(), c.OriginalURL(), err.Error())

	if err.Error() == "Missing or malformed JWT" {
		return xrest.ErrorWithStatus(c, fiber.StatusBadRequest, "missing or malformed jwt")
	}
	return xrest.ErrorWithStatus(c, fiber.StatusUnauthorized, "invalid or expired jwt")
}

func successHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Next()
	}
}
