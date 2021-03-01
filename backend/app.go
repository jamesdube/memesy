package main

import (
	"backend/client"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)
import "github.com/gofiber/fiber/v2"

type CreateMemeRequest struct {
	Name string
	TopText string
	BottomText string
}

type CreateMemeResponse struct {
	Name string
	Image string
}

func main() {

	app := fiber.New()

	app.Use(cors.New())

	app.Get("/memes",getMemes)
	app.Post("/memes",createMemes)

	err := app.Listen(":3120")

	log.Fatalf("error: %s", err)
}

func getMemes(ctx *fiber.Ctx) error {

	log.Printf("Request for meme pictures")

	c := client.NewMemeClient()
	memes := c.GetMemes()

	return ctx.JSON(&fiber.Map{
		"success": true,
		"data":    memes,
	})
}

func createMemes(ctx *fiber.Ctx) error {

	var cmr CreateMemeRequest

	err := ctx.BodyParser(&cmr)
	if err != nil {
		return err
	}

	log.Printf("Create new meme of %s with text %s",cmr.Name,cmr.TopText)

	c := client.NewMemeClient()
	m := c.CreateMeme(cmr.Name,cmr.TopText,cmr.BottomText)

	return ctx.Status(200).JSON(&fiber.Map{
		"Name":"meme",
		"Image": "data:image/png;base64," +m,
	})

}
