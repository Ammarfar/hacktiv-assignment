package helpers

import (
	"assignment/models"
	"strconv"
)

func GenerateUser(total int) []models.User {
	var user = models.User{}
	var users = []models.User{}

	for i := 1; i < total+1; i++ {
		order := strconv.Itoa(i)
		users = append(users, *user.NewUser("User "+order, "Address "+order, "Job "+order, "Reason "+order))
	}

	return users
}
