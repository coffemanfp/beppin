package handlers_test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/coffemanfp/beppin/config"
	dbu "github.com/coffemanfp/beppin/database/utils"
	"github.com/coffemanfp/beppin/helpers"
	"github.com/coffemanfp/beppin/models"
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/stretchr/testify/assert"
)

// Example data
var (
	exampleLanguages models.Languages
	exampleUsers     models.Users
	exampleProducts  models.Products
)

// Database
var (
	db       *sql.DB
	fixtures *testfixtures.Loader
)

func TestMain(m *testing.M) {
	var err error

	log.SetFlags(log.Llongfile)

	config.SetDefaultSettings()

	db, err = sql.Open("postgres", config.GlobalSettings.Database.URL)
	if err != nil {
		log.Fatalln(err)
	}

	fixtures, err = testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("postgres"),
		testfixtures.Directory("../testdata/fixtures/"),
	)
	if err != nil {
		log.Fatalln(err)
	}

	if err = fixtures.EnsureTestDatabase(); err != nil {
		log.Fatalln(err)
	}

	if err = fixtures.Load(); err != nil {
		log.Fatalln(err)
	}

	// Populates example vars with the database data
	dbLanguages, err := dbu.SelectLanguages(db, 20, 0)
	if err != nil {
		return
	}

	exampleLanguages = helpers.ShouldParseDBModelToModel(dbLanguages).(models.Languages)

	dbUsers, err := dbu.SelectUsers(db, 20, 0)
	if err != nil {
		return
	}

	exampleUsers = helpers.ShouldParseDBModelToModel(dbUsers).(models.Users)

	// Populates example vars with the database data
	dbProducts, err := dbu.SelectProducts(db, 20, 0)
	if err != nil {
		return
	}

	exampleProducts = helpers.ShouldParseDBModelToModel(dbProducts).(models.Products)

	os.Exit(m.Run())
}

func TestDatabasePing(t *testing.T) {
	err := db.Ping()
	assert.Nil(t, err)
}
