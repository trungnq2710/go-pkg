// Created at 11/19/2021 4:25 PM
// Developer: trungnq2710 (trungnq2710@gmail.com)

package xfiber

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"
	"runtime"
)

var defaultStackTraceBufLen = 1024

var loggerMiddleware = logger.New(logger.Config{
	Format:     "${time} - [${ip}:${port}] - ${pid}:${locals:requestid} - ${status} - ${method} ${path}\n",
	TimeFormat: "02-01-2006 15:04:05",
})

func recoverMiddleware(c *fiber.Ctx) (err error) {
	defer func() {
		if r := recover(); r != nil {
			stackTraceHandler(c, r)
			c.Status(fiber.StatusInternalServerError).
				JSON(fiber.Map{
					"status":  "error",
					"success": false,
					"message": fmt.Errorf("%v", r).Error(),
				})
		}
	}()

	return c.Next()
}

func stackTraceHandler(c *fiber.Ctx, e interface{}) {
	//debug.PrintStack()
	buf := make([]byte, defaultStackTraceBufLen)
	buf = buf[:runtime.Stack(buf, false)]
	_, _ = os.Stderr.WriteString(fmt.Sprintf("panic: %v\n%s\n", e, buf))
}
