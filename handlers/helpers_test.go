package handlers_test

import (
	"testing"

	errs "github.com/coffemanfp/beppin-server/errors"
	"github.com/stretchr/testify/assert"
)

func assertInvalidParam(t *testing.T, param string, err error) {
	t.Helper()

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), errs.ErrInvalidParam)
	assert.Contains(t, err.Error(), param)
}
