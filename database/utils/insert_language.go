package utils

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/coffemanfp/beppin-server/database/models"
	errs "github.com/coffemanfp/beppin-server/errors"
)

// InsertLanguage - Insert a language.
func InsertLanguage(db *sql.DB, language models.Language) (err error) {
	exists, err := ExistsLanguage(db, 0, language.Code)
	if err != nil {
		return
	}

	if exists {
		err = errors.New(errs.ErrExistentObject)
		return
	}

	query := `
		INSERT INTO
			languages(code, status)
		VALUES
			($1, $2)
		ON CONFLICT DO
			NOTHING
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the insert language statement:\n%s", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		language.Code,
		language.Status,
	)
	if err != nil {
		err = fmt.Errorf("failed to execute insert language statement:\n%s", err)
	}
	return
}
