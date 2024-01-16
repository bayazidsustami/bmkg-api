package clients

import "github.com/gofiber/fiber/v2"

func GetClient() *fiber.Client {
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	return client
}
