package database

import "github.com/coffemanfp/beppin/database/models"

// UserStorage reprensents all implementations for user utils.
type UserStorage interface {
	CreateUser(user models.User) (int, error)
	Login(userToLogin models.User) (models.User, bool, error)
	ExistsUser(user models.User) (bool, error)
	GetUser(userToFind models.User) (models.User, error)
	GetUsers(limit, offset int) (models.Users, error)
	UpdateUser(userToUpdate, user models.User) (int, error)
	UpdateAvatar(avatarURL string, user models.User) (int, error)
	DeleteUser(userToDelete models.User) (int, error)
}
