package controllers_test

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/coffemanfp/beppin-server/database/models"
// 	dbu "github.com/coffemanfp/beppin-server/database/utils"
// 	"github.com/labstack/echo"
// )

// var (
// 	user = models.User{
// 		Language: models.Language{
// 			Code: "es-ES",
// 		},
// 		Name:     "Franklin",
// 		LastName: "Peñaranda",
// 		Username: "coffemanfp",
// 		Password: "1234",
// 		Theme:    "dark",
// 	}
// )

// // TestCreateProduct
// func TestCreateProduct(t *testing.T) {
// 	e := echo.New()
// 	req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(userJSON))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()

// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	mock.ExpectQuery(`
// 		SELECT
// 			EXISTS(
// 				SELECT
// 					id
// 				FROM
// 					users
// 				WHERE
// 					id = $1 OR username = $2
// 			)
// 	`)

// 	mock.ExpectedExec("INSERT INTO users").
// 		WithArgs(
// 			user.Language.Code,
// 			user.name,
// 			user.lastName,
// 			user.username,
// 			user.password,
// 			user.theme,
// 		).
// 		WillReturnResult(NewResult(1, 1))

// 	err = dbu.InsertUser(db, user)

// 	c := e.NewContext(req, rec)
// }
