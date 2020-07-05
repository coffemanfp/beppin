package utils

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/coffemanfp/beppin-server/database/models"
	errs "github.com/coffemanfp/beppin-server/errors"
)

// SelectUser - Selects a user.
func SelectUser(db *sql.DB, userID int, username string) (user models.User, err error) {
	query := `
		SELECT
			id, language, username, name, last_name, birthday, theme, created_at, updated_at
		FROM
			users
		WHERE
			id = $1
			
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the select user statement:\n%s", err)

		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(userID).Scan(
		&user.ID,
		&user.Language.Code,
		&user.Username,
		&user.Name,
		&user.LastName,
		&user.Birthday,
		&user.Theme,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.New(errs.ErrNotExistentObject)
			return
		}

		err = fmt.Errorf("failed to select the user:\n%s", err)
		return
	}
	return
}
