package helpers

import (
	"testing"

	"github.com/muhammadshanoop/identity-reconciliation/models"
	"github.com/stretchr/testify/assert"
)

func TestCheckIfSameContactPresent(t *testing.T) {
	email := "test@example.com"
	phone := "1234567890"

	t.Run("returns primary ID when email and phone match", func(t *testing.T) {
		contacts := []models.Contact{
			{ID: 1, Email: &email, PhoneNumber: &phone, LinkPrecedence: models.Primary},
		}
		details := &models.ContactDetails{
			Email:       &email,
			PhoneNumber: &phone,
		}

		result := checkIfSameContactPresent(details, &contacts)
		assert.NotNil(t, result)
		assert.Equal(t, uint(1), *result)
	})

	t.Run("returns primary ID when request has only email", func(t *testing.T) {
		contacts := []models.Contact{
			{ID: 2, Email: &email, PhoneNumber: nil, LinkPrecedence: models.Primary},
		}
		details := &models.ContactDetails{
			Email:       &email,
			PhoneNumber: nil,
		}

		result := checkIfSameContactPresent(details, &contacts)
		assert.NotNil(t, result)
		assert.Equal(t, uint(2), *result)
	})

	t.Run("returns nil when no match found", func(t *testing.T) {
		contacts := []models.Contact{
			{ID: 3, Email: &email, PhoneNumber: &phone, LinkPrecedence: models.Primary},
		}
		otherEmail := "other@example.com"
		otherPhone := "9876543210"

		details := &models.ContactDetails{
			Email:       &otherEmail,
			PhoneNumber: &otherPhone,
		}

		result := checkIfSameContactPresent(details, &contacts)
		assert.Nil(t, result)
	})
}
