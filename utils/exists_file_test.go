package utils_test

import (
	"testing"

	"github.com/coffemanfp/beppin-server/utils"
	"github.com/stretchr/testify/assert"
)

func TestExistsFile(t *testing.T) {
	exists, err := utils.ExistsFile("../main.go")
	if err != nil {
		t.Errorf("unexpected error:\n%s", err)
	}

	if !exists {
		t.Errorf("exists value (%t) invalid, expected (%t)", exists, true)
	}
}

func TestFailedExistsFile(t *testing.T) {
	t.Run("shouldNotExists", func(t *testing.T) {
		exists, err := utils.ExistsFile("asdalksjdlkajsd")
		assert.Nil(t, err)
		assert.Equal(t, false, exists)
	})

	t.Run("noData", func(t *testing.T) {
		exists, err := utils.ExistsFile("")
		assert.Nil(t, err)
		assert.Equal(t, false, exists)
	})
}
