package utils

import (
	"database/sql"
	"fmt"

	"github.com/coffemanfp/beppin/database/models"
	errs "github.com/coffemanfp/beppin/errors"
)

// InsertUser - Insert a user.
func InsertUser(db *sql.DB, user models.User) (id int, err error) {
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
			users(language, username, password, email, name, last_name, birthday, theme)
		VALUES
			($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING
			id
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the insert user statement: %v", err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		user.Language.Code,
		user.Username,
		user.Password,
		user.Name,
		user.Email,
		user.LastName,
		user.Birthday.Time,
		user.Theme,
	).Scan(&id)
	if err != nil {
		err = fmt.Errorf("failed to execute insert user statement: %v", err)
	}
	return
}
