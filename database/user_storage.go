package database

import "github.com/coffemanfp/beppin-server/database/models"

// UserStorage reprensents all implementations for user utils.
type UserStorage interface {
	CreateUser(user models.User) error
	Login(userToLogin models.User) (models.User, bool, error)
	GetUser(userToFind models.User) (models.User, error)
	GetUsers(limit, offset int) (models.Users, error)
	UpdateUser(userToUpdate, user models.User) error
	UpdateAvatar(avatarURL string, user models.User) error
	DeleteUser(userToDelete models.User) error
}
