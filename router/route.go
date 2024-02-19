package router

import (
	"github.com/fadilghulam/go-fiber-test/handler"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	// grouping
	api := app.Group("/api")

	// users
	user := api.Group("/user")
	user.Get("/", handler.GetAllUsers)
	user.Get("/:id", handler.GetSingleUser)
	user.Post("/", handler.CreateUser)
	user.Put("/:id", handler.UpdateUser)
	user.Delete("/:id", handler.DeleteUserByID)

	transaction := api.Group("/transaction")
	transaction.Get("/", handler.GetTransactions)
	transaction.Get("/:id", handler.GetTransactionsDetails)
}
