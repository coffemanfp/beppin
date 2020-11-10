package utils

import (
	"database/sql"
	"fmt"

	"github.com/coffemanfp/beppin/models"
	errs "github.com/coffemanfp/beppin/errors"
)

// InsertProduct - Insert a product.
func InsertProduct(db *sql.DB, product models.Product) (id int, err error) {
	if db == nil {
		err = errs.ErrClosedDatabase
		return
	}

	query := `
		INSERT INTO
			products(user_id, name, description)
		VALUES
			($1, $2, $3)
		RETURNING
			id
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the insert product statement: %v", err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		product.UserID,
		product.Name,
		product.Description,
	).Scan(&id)
	if err != nil {
		err = fmt.Errorf("failed to execute insert product statement: %v", err)
	}
	return
}
