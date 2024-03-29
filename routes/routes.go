package routes

import (
	"backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	//Create user
	app.Post("/users/create", handlers.CreateUser)

	//Get users' list
	app.Get("/users/list", handlers.GetUsers)

	//Get user by id
	app.Get("/users/:id", handlers.GetIdByPath)

	//Update user	
	app.Put("/users/update", handlers.UpdateUser)

	//Delete user
	app.Delete("/users/delete/:id", handlers.DeleteUser)
}
