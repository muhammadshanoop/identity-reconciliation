package helpers

import (
	"testing"

	"github.com/muhammadshanoop/identity-reconciliation/models"
	"github.com/muhammadshanoop/identity-reconciliation/testutils"
	"github.com/stretchr/testify/assert"
)

func TestShouldCreatePrimaryContact(t *testing.T) {
	testutils.SetupTestDB()

	email := "primary@example.com"
	phone := "1112223333"
	details := &models.ContactDetails{
		Email:       &email,
		PhoneNumber: &phone,
	}

	t.Run("creates primary contact when no existing contacts", func(t *testing.T) {
		result := shouldCreatePrimaryContact([]models.Contact{}, details)
		assert.NotNil(t, result)
		assert.Equal(t, uint(1), *result)
	})
	t.Run("returns nil when contacts already exist", func(t *testing.T) {
		// Insert a dummy contact
		existing := []models.Contact{
			{ID: 2, Email: &email, PhoneNumber: &phone, LinkPrecedence: models.Primary},
		}

		result := shouldCreatePrimaryContact(existing, details)
		assert.Nil(t, result)
	})
}

func TestShouldCreateSecondaryContact(t *testing.T) {
	db := testutils.SetupTestDB()

	email := "secondary@example.com"
	phone := "9998887777"
	details := &models.ContactDetails{
		Email:       &email,
		PhoneNumber: &phone,
	}

	// First create a primary
	primary := models.Contact{
		Email:          &email,
		PhoneNumber:    &phone,
		LinkPrecedence: models.Primary,
	}
	db.Create(&primary)

	existing := []models.Contact{primary}

	t.Run("creates secondary contact linked to primary", func(t *testing.T) {
		result := shouldCreateSecondaryContact(existing, details)
		assert.NotNil(t, result)
		assert.Equal(t, primary.ID, *result)

		var contacts []models.Contact
		db.Find(&contacts)
		assert.Len(t, contacts, 2)
		assert.Equal(t, models.Secondary, contacts[1].LinkPrecedence)
	})
}

func TestCheckTwoPrimaryContactPresent(t *testing.T) {
	db := testutils.SetupTestDB()

	email1 := "first@example.com"
	email2 := "second@example.com"

	// Create two primaries
	primary1 := models.Contact{Email: &email1, LinkPrecedence: models.Primary}
	primary2 := models.Contact{Email: &email2, LinkPrecedence: models.Primary}
	db.Create(&primary1)
	db.Create(&primary2)

	existing := []models.Contact{primary1, primary2}

	t.Run("demotes one primary to secondary", func(t *testing.T) {
		result := checkTwoPrimaryContactPresent(existing)
		assert.NotNil(t, result)

		var updated models.Contact
		db.First(&updated, primary2.ID) // second one should be updated
		assert.Equal(t, models.Secondary, updated.LinkPrecedence)
		assert.Equal(t, primary1.ID, *updated.LinkedID)
	})
}

func TestFindOrCreateContact(t *testing.T) {
	testutils.SetupTestDB()
	email := "unique@example.com"
	phone := "1234567890"
	details := &models.ContactDetails{
		Email:       &email,
		PhoneNumber: &phone,
	}

	t.Run("creates a new primary contact", func(t *testing.T) {
		id, err := FindOrCreateContact(details)
		assert.NoError(t, err)
		assert.NotNil(t, id)
		assert.Equal(t, uint(1), *id)
	})

	t.Run("finds existing contact without creating new one", func(t *testing.T) {
		id, err := FindOrCreateContact(details)
		assert.NoError(t, err)
		assert.NotNil(t, id)
		assert.Equal(t, uint(1), *id) // should reuse same primary
	})
}

func TestGetAllLinkedContacts(t *testing.T) {
	db := testutils.SetupTestDB()

	email1 := "p@example.com"
	primary := models.Contact{Email: &email1, LinkPrecedence: models.Primary}
	db.Create(&primary)

	email2 := "s@example.com"
	secondary := models.Contact{
		Email:          &email2,
		LinkPrecedence: models.Secondary,
		LinkedID:       &primary.ID,
	}
	db.Create(&secondary)

	t.Run("returns all linked contacts", func(t *testing.T) {
		contacts, err := GetAllLinkedContacts(primary.ID)
		assert.NoError(t, err)
		assert.NotNil(t, contacts)
		assert.Len(t, *contacts, 2)
		assert.Equal(t, models.Primary, (*contacts)[0].LinkPrecedence)
		assert.Equal(t, models.Secondary, (*contacts)[1].LinkPrecedence)
	})
}
