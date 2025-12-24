package main

import (
	"flag"
	"github.com/gofiber/fiber/v2"
	"github.com/wayne19393/hotelReservation/api"
)

func main() {
	listenAddr := flag.String("listen", ":8080", "Listen address of API server")
	flag.Parse()
	app := fiber.New()
	appv1 := app.Group("/api/v1")
	appv1.Get("/user", api.HandleGetUsers)
	appv1.Get("/user/:id", api.HandleGetUser)
	app.Listen(*listenAddr)
}
