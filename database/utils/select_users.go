package utils

import (
	"database/sql"
	"errors"
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
)

// SelectUsers - Select a users list.
func SelectUsers(db *sql.DB, limit, offset int) (users models.Users, err error) {
	if db == nil {
		err = errs.ErrClosedDatabase
		return
	}

	query := `
		SELECT
			id, language, avatar, username, email, name, last_name, birthday, theme, currency, created_at, updated_at
		FROM	
			users
		ORDER BY
			id
		LIMIT
			$1
		OFFSET
			$2
	`

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

	// Helper value for null database retorning
	nullData := struct {
		AvatarURL *sql.NullString
		Name      *sql.NullString
		LastName  *sql.NullString
		Birthday  *sql.NullTime
		UpdatedAt *sql.NullTime
	}{}

	var user models.User

	for rows.Next() {
		err = rows.Scan(
			&user.ID,
			&user.Language,
			&nullData.AvatarURL,
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

		// Check if isn't empty to access its value
		if nullData.AvatarURL != nil {
			user.Avatar.URL = nullData.AvatarURL.String
		}
		if nullData.Name != nil {
			user.Name = nullData.Name.String
		}
		if nullData.LastName != nil {
			user.LastName = nullData.LastName.String
		}
		if nullData.Birthday != nil {
			user.Birthday = &nullData.Birthday.Time
		}
		if nullData.UpdatedAt != nil {
			user.UpdatedAt = &nullData.UpdatedAt.Time
		}

		users = append(users, user)
		// Empty the value to avoid overwrite
		user = models.User{}
	}

	return
}
