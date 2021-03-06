package utils

import (
	"database/sql"
	"errors"
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
)

// SelectLanguages - Select a languages list.
func SelectLanguages(dbtx DBTX, limit, offset int) (languages models.Languages, err error) {
	if dbtx == nil {
		err = errs.ErrClosedDatabase
		return
	}

	query := `
	SELECT
		id, code, status, created_at, updated_at
	FROM
		languages
	ORDER BY
		id
	LIMIT
		$1
	OFFSET
		$2
	`

	stmt, err := dbtx.Prepare(query)
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
	var nullData nullLanguageData

	for rows.Next() {
		err = rows.Scan(
			&language.ID,
			&language.Code,
			&language.Status,
			&language.CreatedAt,
			&nullData.UpdatedAt,
		)
		if err != nil {
			err = fmt.Errorf("failed to scan language: %v", err)
			return
		}

		nullData.setResults(&language)
		languages = append(languages, language)

		// Empty the value to avoid overwrite
		language = models.Language{}
	}

	return
}
