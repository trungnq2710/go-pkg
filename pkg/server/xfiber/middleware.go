// Created at 11/19/2021 4:25 PM
// Developer: trungnq2710 (trungnq2710@gmail.com)

package xfiber

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func loggerMiddleware(c *fiber.Ctx) error {
	// TODO impl
	log.Printf("[%s] - [%s] => %s ", c.Method(), c.IP(), c.OriginalURL())
	return c.Next()
}

func recoverMiddleware(c *fiber.Ctx) error {
	// TODO impl
	return c.Next()
}
