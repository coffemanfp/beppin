package utils

import (
	"database/sql"
	"fmt"

	"github.com/coffemanfp/beppin-server/config"
	"github.com/coffemanfp/beppin-server/database/models"
)

// SelectUsers - Select a users list.
func SelectUsers(db *sql.DB, limit int, offset int) (users models.Users, err error) {
	query := `
		SELECT
			users.id, languages.code, username, password, name, last_name, birthday, theme, users.created_at, users.updated_at
		FROM	
			users
		INNER JOIN
			languages
		ON
			languages.id = users.language_id
		LIMIT
			$1
		OFFSET
			$2
	`

	settings, err := config.GetSettings()
	if err != nil {
		return
	}

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
		err = fmt.Errorf("failed to select the users:\n%s", err)
		return
	}

	var user models.User

	for rows.Next() {
		err = rows.Scan(
			&user.ID,
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
			err = fmt.Errorf("failed to scan a user:\n%s", err)
			return
		}

		users = append(users, user)
	}

	return
}
