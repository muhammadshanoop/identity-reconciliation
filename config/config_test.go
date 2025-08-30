package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	if err := os.Setenv("DB_HOST", "localhost"); err != nil {
		t.Fatalf("failed to set DB_HOST: %v", err)
	}
	if err := os.Setenv("DB_USERNAME", "user"); err != nil {
		t.Fatalf("failed to set DB_USERNAME: %v", err)
	}
	if err := os.Setenv("DB_PASSWORD", "pass"); err != nil {
		t.Fatalf("failed to set DB_PASSWORD: %v", err)
	}
	if err := os.Setenv("DB_DATABASE", "testdb"); err != nil {
		t.Fatalf("failed to set DB_HOST: %v", err)
	}
	if err := os.Setenv("DB_PORT", "5432"); err != nil {
		t.Fatalf("failed to set DB_USERNAME: %v", err)
	}

	dsn := GetDSN()
	dsnValue := "host=localhost user=user password=pass dbname=testdb port=5432 sslmode=disable TimeZone=Asia/Kolkata"

	assert.Equal(t, dsn, dsnValue)
}
