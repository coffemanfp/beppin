package database_test

import (
	"testing"

	"github.com/coffemanfp/bepping-server/config"
	"github.com/coffemanfp/bepping-server/database"
)

func TestGet(t *testing.T) {
	err := config.SetSettingsByFile("../config.yaml")
	if err != nil {
		t.Error("unexpected error:\n", err)
	}

	dbConn, err := database.Get()
	if err != nil {
		t.Error("unexpected error:\n", err)
	}
	defer database.CloseConn()

	err = dbConn.Ping()
	if err != nil {
		t.Errorf("failed to ping to the database:%s\b", err)
	}
}

func TestMaxConns(t *testing.T) {
	err := config.SetSettingsByFile("../config.yaml")
	if err != nil {
		t.Error("unexpected error:\n", err)
	}

	dbConn, err := database.Get()
	if err != nil {
		t.Error("unexpected error:\n", err)
	}
	defer database.CloseConn()

	maxConns := dbConn.MaxIdleConns()

	if maxConns != 1 {
		t.Errorf("max connections (%d) invalid, expected (%d)", maxConns, 1)
	}
}
