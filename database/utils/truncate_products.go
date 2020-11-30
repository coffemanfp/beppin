package utils

import (
	"fmt"
)

// TruncateProducts deletes all the product records.
func TruncateProducts(dbtx DBTX, cascade bool) (err error) {
	var query string

	if cascade {
		query = `
			TRUNCATE TABLE
				products
			CASCADE
		`
	} else {
		query = `
			TRUNCATE TABLE
				products
		`
	}

	stmt, err := dbtx.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the truncate table products statement: %v", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		err = fmt.Errorf("failed to exec the truncate table products statement: %v", err)
	}
	return
}
