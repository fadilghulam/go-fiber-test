package handler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fadilghulam/go-fiber-test/database"
	"github.com/fadilghulam/go-fiber-test/dto"
	"github.com/gofiber/fiber/v2"
)

func GetTransactions(c *fiber.Ctx) error {
	db := database.DB.Db
	var transaction []dto.TransactionSale
	// find all transaction in the database
	db.Find(&transaction)
	// If no transaction found, return an error
	if len(transaction) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Transaction not found", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Transaction Found", "data": transaction})
}

func GetTransactionsDetails(c *fiber.Ctx) error {
	db := database.DB.Db

	idStr := c.Params("id")

	var ids []int
	idValues := strings.Split(idStr, ",")
	for _, idValue := range idValues {
		// Handle 'id' values in the format 'id[]'
		idInt, _ := strconv.Atoi(idValue)
		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'id' parameter"})
		// 	return
		// }
		ids = append(ids, idInt)
	}

	query := "SELECT * FROM transaction_sale ts"
	queryDetail := "SELECT * FROM transaction_sale_detail tsd"
	if len(ids) > 0 {
		idStr := strconv.Itoa(ids[0])

		for i := 1; i < len(ids); i++ {
			idStr += "," + strconv.Itoa(ids[i])
		}
		if len(ids) >= 1 && ids[0] != 0 {
			query += fmt.Sprintf(" WHERE ts.id IN (%s)", idStr)
			queryDetail += fmt.Sprintf(" WHERE tsd.transaction_id IN (%s)", idStr)
		}
	}

	// find all transaction in the database
	rows, err := db.Raw(query).Rows()
	if err != nil {
		return err
	}
	defer rows.Close()

	// Check if no rows were returned
	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return err
		}
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Transactions not found", "data": nil})
	}

	// Reset rows to the beginning
	rows.Close()
	rows, err = db.Raw(query).Rows()
	if err != nil {
		return err
	}
	defer rows.Close()

	rows2, err := db.Raw(queryDetail).Rows()
	if err != nil {
		return err
	}
	defer rows2.Close()

	var result []map[string]interface{}
	// Iterate over rows and scan into map
	for rows.Next() {
		var values = make(map[string]interface{})
		if err := db.ScanRows(rows, &values); err != nil {
			return err
		}
		result = append(result, values)
	}

	var resultDetail []map[string]interface{}
	// Iterate over rows and scan into map
	for rows2.Next() {
		var values = make(map[string]interface{})
		if err := db.ScanRows(rows2, &values); err != nil {
			return err
		}
		resultDetail = append(resultDetail, values)
	}

	// Return the result
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Transactions Found", "data": result, "data_detail": resultDetail})
}

func GetTransactionDetailsBy(c *fiber.Ctx) error {
	db := database.DB.Db

	idStr := c.Query("id")
	branchIDsBytes := c.Request().URI().QueryArgs().PeekMulti("id[]")

	var ids []int
	if idStr == "" {
		ids = dto.ConvertParam(branchIDsBytes)
	} else {
		idValues := strings.Split(idStr, ",")
		for _, idValue := range idValues {
			idInt, _ := strconv.Atoi(idValue)
			ids = append(ids, idInt)
		}
	}

	query := "SELECT * FROM transaction_sale ts"
	queryDetail := "SELECT * FROM transaction_sale_detail tsd"
	if len(ids) > 0 {
		idStr := strconv.Itoa(ids[0])

		for i := 1; i < len(ids); i++ {
			idStr += "," + strconv.Itoa(ids[i])
		}
		if len(ids) >= 1 && ids[0] != 0 {
			query += fmt.Sprintf(" WHERE ts.id IN (%s)", idStr)
			queryDetail += fmt.Sprintf(" WHERE tsd.transaction_id IN (%s)", idStr)
		}
	}

	// find all transaction in the database
	rows, err := db.Raw(query).Rows()
	if err != nil {
		return err
	}
	defer rows.Close()

	// Check if no rows were returned
	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return err
		}
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Transactions not found", "data": nil})
	}

	// Reset rows to the beginning
	rows.Close()
	rows, err = db.Raw(query).Rows()
	if err != nil {
		return err
	}
	defer rows.Close()

	rows2, err := db.Raw(queryDetail).Rows()
	if err != nil {
		return err
	}
	defer rows2.Close()

	var result []map[string]interface{}
	// Iterate over rows and scan into map
	for rows.Next() {
		var values = make(map[string]interface{})
		if err := db.ScanRows(rows, &values); err != nil {
			return err
		}
		result = append(result, values)
	}

	var resultDetail []map[string]interface{}
	// Iterate over rows and scan into map
	for rows2.Next() {
		var values = make(map[string]interface{})
		if err := db.ScanRows(rows2, &values); err != nil {
			return err
		}
		resultDetail = append(resultDetail, values)
	}

	// Return the result
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Transactions Found", "data": result, "data_detail": resultDetail})
}

func GetParam(c *fiber.Ctx) error {
	// db := database.DB.Db

	idStr := c.Query("id")
	branchIDsBytes := c.Request().URI().QueryArgs().PeekMulti("id[]")

	var ids []int
	if idStr == "" {
		ids = dto.ConvertParam(branchIDsBytes)
	} else {
		idValues := strings.Split(idStr, ",")
		for _, idValue := range idValues {
			idInt, _ := strconv.Atoi(idValue)
			ids = append(ids, idInt)
		}
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Success", "data": ids})
}
