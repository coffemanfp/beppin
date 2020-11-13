package utils

import (
	"database/sql"
	"errors"
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
)

// Login - Select a user by his username and password, and checks if exists.
func Login(db *sql.DB, userToLogin models.User) (user models.User, match bool, err error) {
	if db == nil {
		err = errs.ErrClosedDatabase
		return
	}

	match = true

	query := `
		SELECT
			id, avatar, language, username, email, theme, currency
		FROM
			users
		WHERE
			username = $1 AND password = $2 OR email = $3 AND password = $2
			
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the login (%s) user statement: %v", userToLogin.Username, err)
		return
	}
	defer stmt.Close()

	var nullData nullUserData

	err = stmt.QueryRow(
		userToLogin.Username,
		userToLogin.Password,
		userToLogin.Email,
	).Scan(
		&user.ID,
		&nullData.AvatarURL,
		&user.Language,
		&user.Username,
		&user.Email,
		&user.Theme,
		&user.Currency,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = nil
			match = false
			return
		}

		err = fmt.Errorf("failed to login (%s) user: %v", userToLogin.Username, err)
	}

	nullData.setResults(&user)
	return
}
