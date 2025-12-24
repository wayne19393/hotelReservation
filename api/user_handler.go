package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wayne19393/hotelReservation/types"
)

func HandleGetUsers(c *fiber.Ctx) error { //this is a handler alwyas the same syntax tbh
	u := types.User{
		Firstname: "James",
		Lastname:  "Bond",
	}
	return c.JSON(u)
}
func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"user": "Kia"})
}
