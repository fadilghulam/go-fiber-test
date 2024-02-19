package handler

import (
	"github.com/fadilghulam/go-fiber-test/database"
	"github.com/fadilghulam/go-fiber-test/dto"
	"github.com/gofiber/fiber/v2"
)

func getTransactions(c *fiber.Ctx) error {
	db := database.DB.Db
	// find all users in the database
	transactions := db.Raw("SELECT * FROM transaction_sale").Scan(&dto.TransactionSale{})
	// If no user found, return an error
	if transactions == nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Users not found", "data": nil})
	}
	// return users
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Users Found", "data": transactions})
}
