package utils

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/coffemanfp/beppin-server/database/models"
	errs "github.com/coffemanfp/beppin-server/errors"
)

// InsertUser - Insert a user.
func InsertUser(db *sql.DB, user models.User) (err error) {
	exists, err := ExistsUser(db, 0, user.Username)
	if err != nil {
		return
	}

	language, err := SelectLanguage(db, user.Language.ID, user.Language.Code)
	if err != nil {
		return
	}

	user.Language = language

	if exists {
		err = errors.New(errs.ErrExistentObject)
		return
	}

	query := `
		INSERT INTO
			users(language_id, username, password, name, last_name, birthday, theme)
		VALUES
			($1, $2, $3, $4, $5, $6, $7)
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the insert user statement:\n%s", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		user.Language.ID,
		user.Username,
		user.Password,
		user.Name,
		user.LastName,
		user.Birthday.Time,
		user.Theme,
	)
	if err != nil {
		err = fmt.Errorf("failed to execute insert user statement:\n%s", err)
	}
	return
}
