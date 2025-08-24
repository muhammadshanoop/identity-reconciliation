package helpers

import (
	"github.com/muhammadshanoop/identity-reconciliation/database"
	"github.com/muhammadshanoop/identity-reconciliation/models"
)

func FindOrCreateContact(contactDetails *models.ContactDetails) (*[]models.Contact, error) {
	var existingContact []models.Contact
	// Search by either email OR phone
	err := database.DB.Where("email = ? OR phone_number = ?", contactDetails.Email, contactDetails.PhoneNumber).
		Find(&existingContact).Error
	if err != nil {
		return nil, err
	}

	// 1. If exact contact already exists (prevent duplicate)
	for _, c := range existingContact {
		if (c.Email != nil && contactDetails.Email != nil && *c.Email == *contactDetails.Email) &&
			(c.PhoneNumber != nil && contactDetails.PhoneNumber != nil && *c.PhoneNumber == *contactDetails.PhoneNumber) {
			// Found exact same contact → return without creating duplicate
			return &existingContact, nil
		}
	}

	// 2. If nothing found → create a new Primary contact
	if len(existingContact) == 0 {
		contact := models.Contact{
			Email:          contactDetails.Email,
			PhoneNumber:    contactDetails.PhoneNumber,
			LinkPrecedence: models.Primary,
		}
		if err := database.DB.Create(&contact).Error; err != nil {
			return nil, err
		}
		existingContact = append(existingContact, contact)
		return &existingContact, nil
	}

	// 3. If something exists → create a Secondary contact linked to Primary
	primaryContactID := FindPrimaryContactID(&existingContact)
	newContact := models.Contact{
		Email:          contactDetails.Email,
		PhoneNumber:    contactDetails.PhoneNumber,
		LinkPrecedence: models.Secondary,
		LinkedID:       &primaryContactID,
	}
	if err := database.DB.Create(&newContact).Error; err != nil {
		return nil, err
	}
	existingContact = append(existingContact, newContact)

	return &existingContact, nil
}

func GetAllLinkedContacts(primaryContactID uint) (*[]models.Contact, error) {
	var linkedContacts []models.Contact
	err := database.DB.Where("id =? OR linked_id=?", primaryContactID, primaryContactID).
		Find(&linkedContacts).
		Order("CASE WHEN link_precedence = 'primary' THEN 0 ELSE 1 END").
		Error
	if err != nil {
		return nil, err
	}
	return &linkedContacts, nil
}
