package utils

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/coffemanfp/beppin/database/models"
	errs "github.com/coffemanfp/beppin/errors"
)

// SelectLanguages - Select a languages list.
func SelectLanguages(db *sql.DB, limit, offset int) (languages models.Languages, err error) {
	if db == nil {
		err = errs.ErrClosedDatabase
		return
	}

	query := `
	SELECT
		code, status, created_at, updated_at
	FROM
		languages
	LIMIT
		$1
	OFFSET
		$2
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the select languages statement: %v", err)
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(limit, offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("failed to select languages: %w", errs.ErrNotExistentObject)
			return
		}

		err = fmt.Errorf("failed to select languages: %v", err)
		return
	}

	var language models.Language

	for rows.Next() {
		err = rows.Scan(
			&language.Code,
			&language.Status,
			&language.CreatedAt,
			&language.UpdatedAt,
		)
		if err != nil {
			err = fmt.Errorf("failed to scan language: %v", err)
			return
		}

		languages = append(languages, language)
	}

	return
}
