package utils

import (
	"database/sql"
	"errors"
	"fmt"

	dbm "github.com/coffemanfp/beppin-server/database/models"
	errs "github.com/coffemanfp/beppin-server/errors"
)

// UpdateUser - Updates a user.
func UpdateUser(db *sql.DB, userID int, username string, user dbm.User) (err error) {
	exists, err := ExistsUser(db, userID, "")
	if err != nil {
		return
	}

	if !exists {
		err = errors.New(errs.ErrNotExistentObject)
		return
	}

	previousUserData, err := SelectUser(db, userID, "")
	if err != nil {
		return
	}

	var languageConditional string
	var languageIDOrCode interface{}

	if user.Language.Code != "" && user.Language.ID == 0 {
		languageConditional = "(SELECT id FROM languages WHERE code = $1)"

		languageIDOrCode = user.Language.Code
	} else {
		languageConditional = "$1"
	}

	user = fillUserEmptyFields(user, previousUserData)

	if languageIDOrCode == nil {
		languageIDOrCode = user.Language.ID
	}

	query := fmt.Sprintf(`
		UPDATE
			users
		SET
			language_id = %s,
			username = $2,
			password = $3,
			name = $4,
			last_name = $5,
			birthday = $6,
			theme = $7,
			updated_at = NOW()
		WHERE 
			id = $8 OR username = $9
	`, languageConditional)

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the update user statement:\n%s", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		languageIDOrCode,
		user.Username,
		user.Password,
		user.Name,
		user.LastName,
		user.Birthday,
		user.Theme,
		userID,
		username,
	)
	if err != nil {
		err = fmt.Errorf("failed to execute the update user statement:\n%s", err)
	}
	return
}

func fillUserEmptyFields(user dbm.User, previousUserData dbm.User) dbm.User {

	switch "" {
	case user.Language.Code:
		user.Language.Code = previousUserData.Language.Code
	case user.Username:
		user.Username = previousUserData.Username
	case user.Password:
		user.Password = previousUserData.Password
	case user.Name:
		user.Name = previousUserData.Name
	case user.LastName:
		user.LastName = previousUserData.LastName
	case user.Theme:
		user.Theme = previousUserData.Theme
	}

	if user.Language.ID == 0 {
		user.Language.ID = previousUserData.Language.ID
	}

	if user.Birthday == nil {
		user.Birthday = previousUserData.Birthday
	}

	return user
}
