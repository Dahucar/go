package main

import (
	"fmt"

	"example.com/go_auth/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	client, ctx, cancel, err, auth_collection := database.Connect()

	fmt.Println("client", client)
	fmt.Println("ctx", ctx)
	fmt.Println("cancel", cancel)
	fmt.Println("err", err)
	fmt.Println("auth_collection", auth_collection)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {

		filter := bson.D{{Key: "autorefid", Value: "100"}}
		// document := bson.D{
		// 	{"name", "go_user"},
		// 	{"surname", ""},
		// 	{"email", "test@test,cl"},
		// }
		database.AddUser(auth_collection, ctx, filter)
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
