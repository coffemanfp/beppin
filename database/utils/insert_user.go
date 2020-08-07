package utils

import (
	"database/sql"
	"fmt"

	"github.com/coffemanfp/beppin-server/database/models"
	errs "github.com/coffemanfp/beppin-server/errors"
)

// InsertUser - Insert a user.
func InsertUser(db *sql.DB, user models.User) (err error) {
	identifier := user.GetIdentifier()
	if identifier == nil {
		err = fmt.Errorf("failed to insert user: %w (user)", errs.ErrNotProvidedOrInvalidObject)
		return
	}

	exists, err := ExistsUser(db, user)
	if err != nil {
		return
	}

	if exists {
		err = fmt.Errorf("failed to check (%v) user: %w (user)", identifier, errs.ErrExistentObject)
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
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the insert user statement: %v", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		user.Language.Code,
		user.Username,
		user.Password,
		user.Name,
		user.Email,
		user.LastName,
		user.Birthday.Time,
		user.Theme,
	)
	if err != nil {
		err = fmt.Errorf("failed to execute insert user statement: %v", err)
	}
	return
}
