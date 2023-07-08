package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"github.com/jaswdr/faker"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mike-neutron/go-batch-request-task/src"
)


func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })
	app.Get("/batch", func(c *fiber.Ctx) error {
		ctx := context.Background()
		fake := faker.New()

		// Make batch of products
		batch := src.Batch{}
		countItems := rand.Intn(100)
		for i := 0; i < countItems; i++ {
			batch = append(batch, src.Item{
				Product: src.Product{Name: fake.Car().Maker()+ " " + fake.Car().Model()},
				Stock: int32(rand.Intn(100)),
			})
		}

		serviceInstance := &src.SetStockService{}
		serviceInstance.Process(ctx, batch)

		fmt.Println(batch)

        return c.SendString("Test")
    })

	log.Fatal(app.Listen(":8080"))
}
