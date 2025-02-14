package repositories

import "mingloo/web-api/internal/models"

var users = []models.User{}

func FetchUsers() []models.User {
	return users
}

func SaveUser(user models.User) {
	users = append(users, user)
}
