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
