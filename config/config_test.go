package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_USERNAME", "user")
	os.Setenv("DB_PASSWORD", "pass")
	os.Setenv("DB_DATABASE", "testdb")
	os.Setenv("DB_PORT", "5432")

	dsn := GetDSN()
	dsnValue := "host=localhost user=user password=pass dbname=testdb port=5432 sslmode=disable TimeZone=Asia/Kolkata"

	assert.Equal(t, dsn, dsnValue)
}
