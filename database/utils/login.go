package utils

import (
	"database/sql"
	"fmt"

	"github.com/coffemanfp/beppin-server/database/models"
)

// Login - Select a user by his username and password, and checks if exists.
func Login(db *sql.DB, username string, password string) (user models.User, match bool, err error) {
	match = true

	query := `
		SELECT
			id, language, username, theme
		FROM
			users
		WHERE
			username = $1 AND password = $2
			
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the login user statement:\n%s", err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(username, password).Scan(
		&user.ID,
		&user.Language.Code,
		&user.Username,
		&user.Theme,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			match = false
			return
		}

		err = fmt.Errorf("failed to select the user login:\n%s", err)
	}
	return
}
