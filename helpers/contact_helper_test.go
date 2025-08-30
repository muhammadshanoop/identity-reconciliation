package helpers

import (
	"testing"

	"github.com/muhammadshanoop/identity-reconciliation/models"
	"github.com/stretchr/testify/assert"
)

func TestFindPrimaryContactID(t *testing.T) {
	t.Run("returns primary contact ID when primary exists", func(t *testing.T) {
		contacts := []models.Contact{
			{ID: 1, LinkPrecedence: models.Secondary},
			{ID: 2, LinkPrecedence: models.Primary},
		}

		id := FindPrimaryContactID(&contacts)
		assert.Equal(t, uint(2), id)
	})

	t.Run("returns zero when no primary exists", func(t *testing.T) {
		contacts := []models.Contact{
			{ID: 3, LinkPrecedence: models.Secondary},
			{ID: 4, LinkPrecedence: models.Secondary},
		}

		id := FindPrimaryContactID(&contacts)
		assert.Equal(t, uint(0), id)
	})

	t.Run("returns first primary if multiple exist", func(t *testing.T) {
		contacts := []models.Contact{
			{ID: 5, LinkPrecedence: models.Primary},
			{ID: 6, LinkPrecedence: models.Primary},
		}

		id := FindPrimaryContactID(&contacts)
		assert.Equal(t, uint(5), id) // first match wins
	})
}
