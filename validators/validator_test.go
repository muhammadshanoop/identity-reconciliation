package validators

import (
	"testing"

	"github.com/muhammadshanoop/identity-reconciliation/models"
	"github.com/stretchr/testify/assert"
)

func TestValdiateRequest(t *testing.T) {
	t.Run("should fail when both email and phone are nil", func(t *testing.T) {
		contact := &models.ContactDetails{
			Email:       nil,
			PhoneNumber: nil,
		}
		err := ValidateRequest(contact)
		assert.NotNil(t, err)
		assert.Equal(t, "either email or phoneNumber must be provided", err.Error())
	})

	t.Run("should pass when email provided", func(t *testing.T) {
		email := "testing@gmail.com"
		contact := &models.ContactDetails{
			Email:       &email,
			PhoneNumber: nil,
		}
		err := ValidateRequest(contact)
		assert.Nil(t, err)
	})

	t.Run("should pass when mobile provided", func(t *testing.T) {
		phoneNumber := "9948849938"
		contact := &models.ContactDetails{
			Email:       nil,
			PhoneNumber: &phoneNumber,
		}
		err := ValidateRequest(contact)
		assert.Nil(t, err)
	})
	t.Run("should pass when both email and phone are provided", func(t *testing.T) {
		email := "test@example.com"
		phone := "1234567890"
		contact := &models.ContactDetails{
			Email:       &email,
			PhoneNumber: &phone,
		}

		err := ValidateRequest(contact)
		assert.Nil(t, err)
	})
}
