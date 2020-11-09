package utils

import (
	"database/sql"
	"fmt"

	"github.com/coffemanfp/beppin/database/models"
	errs "github.com/coffemanfp/beppin/errors"
)

// InsertUser - Insert a user.
func InsertUser(db *sql.DB, user models.User) (newUser models.User, err error) {
	if db == nil {
		err = errs.ErrClosedDatabase
		return
	}

	if user.Language.Code != "" {
		var language models.Language
		language, err = SelectLanguage(db, user.Language)
		if err != nil {
			return
		}

		user.Language = language
	}

	query := `
		INSERT INTO
			users(username, password, email)
		VALUES
			($1, $2, $3)
		RETURNING
			id, language, username, theme, currency
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the insert user statement: %v", err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		user.Username,
		user.Password,
		user.Email,
	).Scan(
		&newUser.ID,
		&newUser.Language.Code,
		&newUser.Username,
		&newUser.Theme,
		&newUser.Currency,
	)

	if err != nil {
		err = fmt.Errorf("failed to execute insert user statement: %v", err)
	}
	return
}
