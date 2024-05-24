package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v3"
)

type Beef struct {
	Beef map[string]uint32
}

func main() {
	// Initialize a new Fiber app
	app := fiber.New()

	// Define a route for the GET method on the root path '/'
	app.Get("/", func(c fiber.Ctx) error {
		// Send a string response to the client
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/beef/summary", func(c fiber.Ctx) error {
		data, _ := http.Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
		body, _ := io.ReadAll(data.Body)

		var beefs Beef
		beefs.Beef = map[string]uint32{}

		for _, b := range regexp.MustCompile(`\s+|\.|,`).Split(string(body), -1) {
			if b != "" {
				if _, ok := beefs.Beef[strings.ToLower(b)]; !ok {
					beefs.Beef[strings.ToLower(b)] = 1
				} else {
					beefs.Beef[strings.ToLower(b)]++
				}
			}
		}

		json, _ := json.Marshal(beefs)

		return c.SendString(string(json))
	})

	// Start the server on port 8080
	log.Fatal(app.Listen(":8080"))
}
