package testutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupTestDB(t *testing.T) {
	db := SetupTestDB()
	assert.NotNil(t, db)
}
