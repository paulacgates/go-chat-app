package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pusher/pusher-http-go/v5"
)

func HandleRequest() {
	pusherClient := pusher.Client{
		AppID:   "1582021",
		Key:     "3c5d31c1dbeda1b7789f",
		Secret:  "99f64cadbea19141865f",
		Cluster: "sa1",
		Secure:  true,
	}

	app := fiber.New()
	app.Use(cors.New())

	app.Post("api/messages", func(c *fiber.Ctx) error {
		var data map[string]string

		err := c.BodyParser(&data)
		if err != nil {
			return err
		}
		pusherClient.Trigger("chat", "message", data)

		return c.JSON([]string{})
	})
	log.Fatal(app.Listen(":8080"))
}
