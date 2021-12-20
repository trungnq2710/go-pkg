// Created at 11/30/2021 11:22 PM
// Developer: trungnq2710 (trungnq2710@gmail.com)

package xmiddleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/jwt/v2"

	"gitlab.com/jplatform/jengine/constant"
	"gitlab.com/jplatform/jengine/pkg/util/xrest"
)

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:     []byte(constant.TokenSecret),
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
