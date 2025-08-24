package helpers

import (
	"github.com/muhammadshanoop/identity-reconciliation/database"
	"github.com/muhammadshanoop/identity-reconciliation/models"
)

func FindOrCreateContact(contactDetails *models.ContactDetails) (*[]models.Contact, error) {
	var existingContact []models.Contact
	err := database.DB.Where("email=? OR phone_number=?", contactDetails.Email, contactDetails.PhoneNumber).Find(&existingContact).Error
	if err != nil {
		return nil, err
	}
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
	}
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
