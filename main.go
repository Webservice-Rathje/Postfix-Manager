package main

import (
	"github.com/Webservice-Rathje/Postfix-Manager/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	utils.GenerateSecret()

	app := fiber.New(fiber.Config{
		Prefork: true,
	})

}
