package utils

import (
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
)

// SignUp - Inserts a basic user and returns the token data.
func SignUp(dbtx DBTX, user models.User) (newUser models.User, err error) {
	if dbtx == nil {
		err = errs.ErrClosedDatabase
		return
	}

	query := `
		INSERT INTO
			users(username, password, email, name, last_name)
		VALUES
			($1, $2, $3, $4, $5)
		RETURNING
			id, language, username, theme, currency
	`

	stmt, err := dbtx.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the insert user statement: %v", err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		user.Username,
		user.Password,
		user.Email,
		user.Name,
		user.LastName,
	).Scan(
		&newUser.ID,
		&newUser.Language,
		&newUser.Username,
		&newUser.Theme,
		&newUser.Currency,
	)

	if err != nil {
		err = fmt.Errorf("failed to execute insert user statement: %v", err)
	}
	return
}
