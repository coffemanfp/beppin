package utils

import (
	"database/sql"
	"fmt"

	"github.com/coffemanfp/beppin-server/models"
)

// InsertProduct - Insert a product.
func InsertProduct(db *sql.DB, product models.Product) (err error) {
	query := `
		INSERT INTO
			products(user_id, name, description)
		VALUES
			($1, $2, $3)
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the insert product statement:\n%s", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		product.UserID,
		product.Name,
		product.Description,
	)
	if err != nil {
		err = fmt.Errorf("failed to execute insert product statement:\n%s", err)
	}
	return
}
