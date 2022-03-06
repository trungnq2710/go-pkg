// Created at 11/27/2021 10:05 PM
// Developer: trungnq2710 (trungnq2710@gmail.com)

package xrest

import "github.com/gofiber/fiber/v2"

func Error(c *fiber.Ctx, err string) error {
	return c.JSON(&fiber.Map{
		"status":  STATUS_ERROR,
		"success": false,
		"message": err,
	})
}

func ErrorWithHTTPStatus(c *fiber.Ctx, httpStatus int, status, err string) error {
	return c.Status(httpStatus).
		JSON(fiber.Map{
			"status":  status,
			"success": false,
			"message": err,
		})
}

func ErrorWithStatus(c *fiber.Ctx, status, err string) error {
	return c.Status(fiber.StatusOK).
		JSON(fiber.Map{
			"status":  status,
			"success": false,
			"message": err,
		})
}

func WithoutData(c *fiber.Ctx, msg string) error {
	if msg == "" {
		msg = "success"
	}
	return c.JSON(&fiber.Map{
		"status":  STATUS_SUCCESS,
		"success": true,
		"message": msg,
	})
}

func WithData(c *fiber.Ctx, msg string, data interface{}) error {
	if msg == "" {
		msg = "success"
	}
	return c.JSON(&fiber.Map{
		"status":  STATUS_SUCCESS,
		"success": true,
		"message": msg,
		"data":    data,
	})
}

func WithDataset(c *fiber.Ctx, msg string, dataset interface{}) error {
	if msg == "" {
		msg = "success"
	}
	return c.JSON(&fiber.Map{
		"status":  STATUS_SUCCESS,
		"success": true,
		"message": msg,
		"dataset": dataset,
	})
}

func WithOffsetID(c *fiber.Ctx, msg string, dataset interface{}, total int32) error {
	if msg == "" {
		msg = "success"
	}
	return c.JSON(&fiber.Map{
		"status":  STATUS_SUCCESS,
		"success": true,
		"message": msg,
		"dataset": dataset,
		"total":   total,
	})
}

func WithPagination(c *fiber.Ctx, msg string, dataset interface{}, paginate interface{}) error {
	if msg == "" {
		msg = "success"
	}
	return c.JSON(&fiber.Map{
		"status":   STATUS_SUCCESS,
		"success":  true,
		"message":  msg,
		"dataset":  dataset,
		"paginate": paginate,
	})
}
