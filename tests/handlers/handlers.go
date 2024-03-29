package handlers

import (
	"backend/models"
	"backend/tests/storage"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MockHandlers struct {
	MockUserService storage.MockStorage
}

func (m *MockHandlers) CreateUser(c *fiber.Ctx) error {
	var reqUser models.User
	if err := c.BodyParser(&reqUser); err != nil {
		return err
	}

	reqUser.ID = uuid.NewString()

	respUser, err := m.MockUserService.CreateUser(reqUser)
	if err != nil {
		return err
	}

	return c.JSON(respUser)
}

func (m *MockHandlers) GetUsers(c *fiber.Ctx) error {
	page := c.Query("page")
	if page == "" {
		page = "1"
	}
	intPage, err := strconv.Atoi(page)
	if err != nil {
		return err
	}

	limit := c.Query("limit")
	if limit == "" {
		limit = "10"
	}
	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		return err
	}

	respUsers, err := m.MockUserService.GetAllUsers(intPage, intLimit)
	if err != nil {
		return err
	}

	return c.JSON(respUsers)
}

func (m *MockHandlers) UpdateUser(c *fiber.Ctx) error {
	userId := c.Query("id")
	if userId == "" {
		userId = "12345"
	}

	var user models.User
	if err := c.BodyParser(&user); err != nil {
		fmt.Println(user)
		fmt.Println(err, "update 1")
		return err
	}
	fmt.Println(user)
	user.ID = userId

	updatedUser, err := m.MockUserService.UpdateUserById(userId, user)
	if err != nil {
		fmt.Println(err, "update 2")
		return err
	}

	return c.JSON(updatedUser)
}

func (m *MockHandlers) DeleteUser(c *fiber.Ctx) error {
	userId := c.Query("id")

	respUser, err := m.MockUserService.DeleteUserById(userId)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return c.JSON(respUser)
}

func (m *MockHandlers) GetIdByPath(c *fiber.Ctx) error {
	userId := c.Query("id")

	respUser, err := m.MockUserService.GetUserById(userId)
	if err != nil {
		return err
	}

	return c.JSON(respUser)
}
