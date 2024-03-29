package handlers

import (
	"backend/models"
	"backend/storage"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateUser(c *fiber.Ctx) error {
	var reqUser models.User
	if err := c.BodyParser(&reqUser); err != nil {
		return err
	}

	reqUser.ID = uuid.NewString()

	respUser, err := storage.CreateUser(reqUser)
	if err != nil {
		return err
	}

	return c.JSON(respUser)
}

func GetUsers(c *fiber.Ctx) error {
	page := c.Query("page")
	intPage, err := strconv.Atoi(page)
	if err != nil {
		return err
	}

	limit := c.Query("limit")
	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		return err
	}

	respUsers, err := storage.GetAllUsers(intPage, intLimit)
	if err != nil {
		return err
	}

	return c.JSON(respUsers)
}

func UpdateUser(c *fiber.Ctx) error {
	userId := c.Params("id")

	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	updatedUser, err := storage.UpdateUserById(userId, user)
	if err != nil {
		return err
	}

	return c.JSON(updatedUser)
}

func DeleteUser(c *fiber.Ctx) error {
	userId := c.Params("id")

	respUser, err := storage.DeleteUserById(userId)
	if err != nil {
		return err
	}

	return c.JSON(respUser)
}

func GetIdByPath(c *fiber.Ctx) error {
	userId := c.Params("id")

	respUser, err := storage.GetUserById(userId)
	if err != nil {
		return err
	}

	return c.JSON(respUser)
}
