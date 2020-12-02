package utils

import (
	"database/sql"
	"errors"
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
)

// SelectUsers - Select a users list.
func SelectUsers(dbtx DBTX, limit, offset int) (users models.Users, err error) {
	if dbtx == nil {
		err = errs.ErrClosedDatabase
		return
	}

	query := `
		SELECT
			users.id, language, files.id, files.path, username, email, name, last_name, birthday, theme, currency, users.created_at, users.updated_at
		FROM	
			users
		LEFT JOIN
			files
		ON
			users.avatar_id = files.id
		ORDER BY
			users.id
		LIMIT
			$1
		OFFSET
			$2
	`

	stmt, err := dbtx.Prepare(query)
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
	var nullData nullUserData

	for rows.Next() {
		err = rows.Scan(
			&user.ID,
			&user.Language,
			&nullData.AvatarID,
			&nullData.AvatarPath,
			&user.Username,
			&user.Email,
			&nullData.Name,
			&nullData.LastName,
			&nullData.Birthday,
			&user.Theme,
			&user.Currency,
			&user.CreatedAt,
			&nullData.UpdatedAt,
		)
		if err != nil {
			err = fmt.Errorf("failed to scan user: %v", err)
			return
		}

		nullData.setResults(&user)
		if user.Avatar != nil {
			user.Avatar.SetURL()
		}
		users = append(users, user)

		// Empty the value to avoid overwrite
		user = models.User{}
	}

	return
}
