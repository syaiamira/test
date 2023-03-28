package handler

import "github.com/gofiber/fiber/v2"

func HiApaKhabar(ctx *fiber.Ctx) error {

	greetings := "Hi, apa khabar?"

	return ctx.Status(200).JSON(greetings)
}
