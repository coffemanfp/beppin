package database

import (
	"fmt"

	dbu "github.com/coffemanfp/beppin/database/utils"
	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
)

func (dS defaultStorage) SignUp(user models.User) (newUser models.User, err error) {
	identifier := user.GetIdentifier()
	if identifier == nil {
		err = fmt.Errorf("failed to create user: %w (user)", errs.ErrNotProvidedOrInvalidObject)
		return
	}

	exists, err := dS.ExistsUser(user)
	if err != nil {
		return
	}

	if exists {
		err = fmt.Errorf("failed to create (%v) user: %w (user)", identifier, errs.ErrExistentObject)
		return
	}

	if user.Language != "" {
		var language models.Language
		language, err = dbu.SelectLanguage(dS.db, models.Language{Code: user.Language})
		if err != nil {
			return
		}

		user.Language = language.Code
	}

	newUser, err = dbu.SignUp(dS.db, user)
	return
}

func (dS defaultStorage) ExistsUser(user models.User) (exists bool, err error) {
	exists, err = dbu.ExistsUser(dS.db, user)
	return
}

func (dS defaultStorage) Login(userToLogin models.User) (user models.User, match bool, err error) {
	user, match, err = dbu.Login(dS.db, userToLogin)
	return
}

func (dS defaultStorage) GetUser(userToFind models.User) (user models.User, err error) {
	user, err = dbu.SelectUser(dS.db, userToFind)
	return
}

func (dS defaultStorage) GetUsers(limit, offset int) (users models.Users, err error) {
	users, err = dbu.SelectUsers(dS.db, limit, offset)
	return
}

func (dS defaultStorage) UpdateUser(userToUpdate, user models.User) (userUpdated models.User, err error) {
	previousUserData, err := dbu.SelectUser(dS.db, userToUpdate)
	if err != nil {
		return
	}

	fmt.Println(previousUserData)

	user = fillUserEmptyFields(user, previousUserData)

	fmt.Println(user)

	userUpdated, err = dbu.UpdateUser(dS.db, userToUpdate, user)
	return
}

func (dS defaultStorage) UpdateAvatar(avatarURL string, userToUpdate models.User) (id int, err error) {
	identifier := userToUpdate.GetIdentifier()
	if identifier == nil {
		err = fmt.Errorf("failed to check user: %w (user)", errs.ErrNotProvidedOrInvalidObject)
		return
	}

	exists, err := dS.ExistsUser(userToUpdate)
	if err != nil {
		return
	}

	if !exists {
		err = fmt.Errorf("failed to check (%v) user: %w (user)", identifier, errs.ErrNotExistentObject)
		return
	}

	id, err = dbu.UpdateAvatar(dS.db, avatarURL, userToUpdate)
	return
}

func (dS defaultStorage) DeleteUser(userToDelete models.User) (id int, err error) {
	id, err = dbu.DeleteUser(dS.db, userToDelete)
	return
}

func fillUserEmptyFields(user, previousUserData models.User) models.User {
	if user.ID == 0 {
		user.ID = previousUserData.ID
	}
	if user.Language == "" {
		user.Language = previousUserData.Language
	}
	if user.Username == "" {
		user.Username = previousUserData.Username
	}
	if user.Email == "" {
		user.Email = previousUserData.Email
	}
	if user.Password == "" {
		user.Password = previousUserData.Password
	}
	if user.Name == "" {
		user.Name = previousUserData.Name
	}
	if user.LastName == "" {
		user.LastName = previousUserData.LastName
	}
	if user.Theme == "" {
		user.Theme = previousUserData.Theme
	}
	if user.Currency == "" {
		user.Currency = previousUserData.Currency
	}
	if user.Birthday == nil {
		user.Birthday = previousUserData.Birthday
	}
	if user.CreatedAt == nil {
		user.CreatedAt = previousUserData.CreatedAt
	}
	if user.UpdatedAt == nil {
		user.UpdatedAt = previousUserData.UpdatedAt
	}

	return user
}
