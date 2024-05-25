package main

import (
	"fmt"
	"log"
	"net"

	"github.com/63070028/backend-challenge/ex3/server/services"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	services.RegisterBeefSerivceServer(s, services.NewBeefSerivceServer())

	fmt.Println("gRPC server listening on port 50051")

	err = s.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}

}

// func main() {
// 	// Initialize a new Fiber app
// 	app := fiber.New()

// 	// Define a route for the GET method on the root path '/'
// 	app.Get("/", func(c fiber.Ctx) error {
// 		// Send a string response to the client
// 		return c.SendString("Hello, World 👋!")
// 	})

// 	app.Get("/beef/summary", func(c fiber.Ctx) error {
// 		data, _ := http.Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
// 		body, _ := io.ReadAll(data.Body)

// 		var beefs = map[string]map[string]uint32{}
// 		beefs["beef"] = map[string]uint32{}

// 		for _, b := range regexp.MustCompile(`\s+|\.|,`).Split(string(body), -1) {
// 			if b != "" {
// 				if _, ok := beefs["beef"][strings.ToLower(b)]; !ok {
// 					beefs["beef"][strings.ToLower(b)] = 1
// 				} else {
// 					beefs["beef"][strings.ToLower(b)]++
// 				}
// 			}
// 		}

// 		json, _ := json.Marshal(beefs)

// 		return c.SendString(string(json))
// 	})

// 	// Start the server on port 8080
// 	log.Fatal(app.Listen(":8080"))
// }
