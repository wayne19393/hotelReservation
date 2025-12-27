package api

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/wayne19393/hotelReservation/db"
	"github.com/wayne19393/hotelReservation/types"
)

type UserHandler struct {
	userStore db.UserStore
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	var (
		id  = c.Params("id")
		ctx = context.Background()
	)
	user, err := h.userStore.GetUserByID(ctx, id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error { //this is a handler alwyas the same syntax tbh
	u := types.User{
		Firstname: "James",
		Lastname:  "Bond",
	}
	return c.JSON(u)
}
