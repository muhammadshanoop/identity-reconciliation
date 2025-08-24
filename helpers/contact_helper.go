package helpers

import "github.com/muhammadshanoop/identity-reconciliation/models"

func FindPrimaryContactID(contacts *[]models.Contact) uint {
	var primaryContactID uint
	for _, contact := range *contacts {
		if contact.LinkPrecedence == models.Primary {
			primaryContactID = contact.ID
			break
		}
	}
	return primaryContactID
}
