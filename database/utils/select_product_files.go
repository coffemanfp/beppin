package utils

import (
	"database/sql"
	"errors"
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
)

// SelectProductFiles - Selects the files of a product.
func SelectProductFiles(db *sql.DB, productToFind models.Product) (files models.Files, err error) {
	if db == nil {
		err = errs.ErrClosedDatabase
		return
	}

	identifier := productToFind.GetIdentifier()
	if identifier == nil {
		err = fmt.Errorf("failed to select product: %w (product)", errs.ErrNotProvidedOrInvalidObject)
		return
	}

	query := `
		SELECT
			files.id, path, files.created_at, files.updated_at
		FROM
			files
		INNER JOIN
			products
		INNER JOIN
			files_products
		ON
			files_products.product_id = products.id
		ON
			files_products.file_id = files.id
		WHERE
			products.id = $1
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the select (%v) product statement: %v", identifier, err)

		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(productToFind.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("failed to select (%v) product: %w (product)", identifier, errs.ErrNotExistentObject)
			return
		}

		err = fmt.Errorf("failed to select (%v) product: %v", identifier, err)
	}

	var nullData nullFileData
	var file models.File

	for rows.Next() {
		err = rows.Scan(
			&file.ID,
			&file.Path,
			&file.CreatedAt,
			&nullData.UpdatedAt,
		)
		if err != nil {
			err = fmt.Errorf("failed to scan file: %v", err)
			return
		}

		nullData.setResults(&file)
		file.SetURL()
		files = append(files, file)

		// Empty the value to avoid overwrite
		file = models.File{}
	}
	return
}
