package utils

import (
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
)

// InsertProductFile - Inserts a product file.
func InsertProductFile(dbtx DBTX, productID int64, fileID int64) (err error) {
	if dbtx == nil {
		err = errs.ErrClosedDatabase
		return
	}

	query := `
		INSERT INTO
			files_products(file_id, product_id)
		VALUES
			($1, $2)
	`

	stmt, err := dbtx.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the insert file product statement: %v", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(fileID, productID)
	if err != nil {
		err = fmt.Errorf("failed to execute insert file product statement: %v", err)
	}
	return
}
