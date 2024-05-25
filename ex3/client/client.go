package main

import (
	"encoding/json"
	"log"
	"github.com/63070028/backend-challenge/ex3/client/services"
	"github.com/gofiber/fiber/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Initialize a new Fiber app
	app := fiber.New()

	// Define a route for the GET method on the root path '/'
	app.Get("/", func(c fiber.Ctx) error {
		// Send a string response to the client
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/beef/summary", func(c fiber.Ctx) error {
		creds := insecure.NewCredentials()
		bc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
		if err != nil {
			log.Fatal(err)
		}
		defer bc.Close()

		beefSerivceClient := services.NewBeefSerivceClient(bc)
		beefService := services.NewBeefService(beefSerivceClient)

		data, err := beefService.GetBeef()

		if err != nil {
			log.Fatal(err)
		}

		res := make(map[string]map[string]int32)
		res["beefs"] = map[string]int32{}

		for key, value := range data.Beef {
			res["beefs"][key] = value
		}

		json, _ := json.Marshal(res)

		return c.SendString(string(json))
	})

	// Start the server on port 8080
	log.Fatal(app.Listen(":8080"))
}
