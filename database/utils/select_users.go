package utils

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/coffemanfp/beppin-server/config"
	"github.com/coffemanfp/beppin-server/database/models"
	errs "github.com/coffemanfp/beppin-server/errors"
)

// SelectUsers - Select a users list.
func SelectUsers(db *sql.DB, limit, offset uint64) (users models.Users, err error) {
	if db == nil {
		err = errs.ErrClosedDatabase
		return
	}

	query := `
		SELECT
			id, language, avatar, username, email, name, last_name, birthday, theme, created_at, updated_at
		FROM	
			users
		LIMIT
			$1
		OFFSET
			$2
	`

	settings := config.GetSettings()

	if limit == 0 {
		limit = settings.MaxElementsPerPagination
	}

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the select users statement:\n%s", err)
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(limit, offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("failed to select users: %w", errs.ErrNotExistentObject)
			return
		}

		err = fmt.Errorf("failed to select users: %v", err)
		return
	}

	var user models.User

	for rows.Next() {
		err = rows.Scan(
			&user.ID,
			&user.Language.Code,
			&user.AvatarURL,
			&user.Username,
			&user.Email,
			&user.Name,
			&user.LastName,
			&user.Birthday,
			&user.Theme,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			err = fmt.Errorf("failed to scan user: %v", err)
			return
		}

		users = append(users, user)
	}

	return
}
