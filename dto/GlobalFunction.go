package dto

import (
	"strconv"
	"strings"
)

func ConvertParam(param [][]byte) []int {

	var idValuesTest []string
	for _, idBytes := range param {
		idValuesTest = append(idValuesTest, string(idBytes))
	}

	idString := strings.Join(idValuesTest, ",")

	var ids []int
	idValues := strings.Split(idString, ",")
	for _, idValue := range idValues {
		// Handle 'id' values in the format 'id[]'
		idInt, _ := strconv.Atoi(idValue)
		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'id' parameter"})
		// 	return
		// }
		ids = append(ids, idInt)
	}

	return ids
}
