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
	exists, err := ExistsUser(db, userID, username)
	if err != nil {
		return
	}

	if !exists {
		err = errors.New(errs.ErrNotExistentObject)
		return
	}

	query := `
		SELECT
			users.id, languages.id, languages.code, username, password, name, last_name, birthday, theme, users.created_at, users.updated_at
		FROM
			users
		INNER JOIN
			languages
		ON
			users.language_id = languages.id
		WHERE
			users.id = $1
			
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the select user statement:\n%s", err)

		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(userID).Scan(
		&user.ID,
		&user.Language.ID,
		&user.Language.Code,
		&user.Username,
		&user.Password,
		&user.Name,
		&user.LastName,
		&user.Birthday,
		&user.Theme,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		err = fmt.Errorf("failed to select the user:\n%s", err)
	}
	return
}
