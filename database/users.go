package database

import (
	"fmt"

	"github.com/coffemanfp/beppin-server/database/models"
	dbu "github.com/coffemanfp/beppin-server/database/utils"
	errs "github.com/coffemanfp/beppin-server/errors"
)

func (dS defaultStorage) CreateUser(user models.User) (err error) {
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

	if user.Language.Code != "" {
		var language models.Language
		language, err = dbu.SelectLanguage(dS.db, user.Language)
		if err != nil {
			return
		}

		user.Language = language
	}

	err = dbu.InsertUser(dS.db, user)
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

func (dS defaultStorage) UpdateUser(userToUpdate, user models.User) (err error) {
	previousUserData, err := dbu.SelectUser(dS.db, userToUpdate)
	if err != nil {
		return
	}

	user = fillUserEmptyFields(user, previousUserData)

	err = dbu.UpdateUser(dS.db, userToUpdate, user)
	return
}

func (dS defaultStorage) UpdateAvatar(avatarURL string, userToUpdate models.User) (err error) {
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

	err = dbu.UpdateAvatar(dS.db, avatarURL, userToUpdate)
	return
}

func (dS defaultStorage) DeleteUser(userToDelete models.User) (err error) {
	err = dbu.DeleteUser(dS.db, userToDelete)
	return
}

func fillUserEmptyFields(user, previousUserData models.User) models.User {
	switch "" {
	case user.Language.Code:
		user.Language.Code = previousUserData.Language.Code

	case user.Username:
		user.Username = previousUserData.Username

	case user.Email:
		user.Email = previousUserData.Email

	case user.Password:
		user.Password = previousUserData.Password

	case user.Name:
		user.Name = previousUserData.Name

	case user.LastName:
		user.LastName = previousUserData.LastName

	case user.Theme:
		user.Theme = previousUserData.Theme
	}

	if user.Birthday == nil {
		user.Birthday = previousUserData.Birthday
	}

	return user
}
