package helpers

import (
	"github.com/muhammadshanoop/identity-reconciliation/database"
	"github.com/muhammadshanoop/identity-reconciliation/models"
)

func FindOrCreateContact(contactDetails *models.ContactDetails) (*uint, error) {
	var existingContact []models.Contact
	// Search by either email OR phone
	err := database.DB.Where("email = ? OR phone_number = ?", contactDetails.Email, contactDetails.PhoneNumber).
		Find(&existingContact).Error
	if err != nil {
		return nil, err
	}

	primaryID := checkIfSameContactPresent(contactDetails, &existingContact)
	if primaryID != nil {
		return primaryID, nil
	}

	primaryID = checkTwoPrimaryContactPresent(existingContact)
	if primaryID != nil {
		return primaryID, nil
	}

	primaryID = shouldCreatePrimaryContact(existingContact, contactDetails)
	if primaryID != nil {
		return primaryID, nil
	}

	primaryID = shouldCreateSecondaryContact(existingContact, contactDetails)
	if primaryID != nil {
		return primaryID, nil
	}
	return nil, nil
}

func shouldCreateSecondaryContact(existingContact []models.Contact, contactDetails *models.ContactDetails) *uint {
	primaryContactID := FindPrimaryContactID(&existingContact)
	newContact := models.Contact{
		Email:          contactDetails.Email,
		PhoneNumber:    contactDetails.PhoneNumber,
		LinkPrecedence: models.Secondary,
		LinkedID:       &primaryContactID,
	}
	if err := database.DB.Create(&newContact).Error; err != nil {
		return nil
	}
	return &primaryContactID
}

func shouldCreatePrimaryContact(existingContact []models.Contact, contactDetails *models.ContactDetails) *uint {
	if len(existingContact) == 0 {
		contact := models.Contact{
			Email:          contactDetails.Email,
			PhoneNumber:    contactDetails.PhoneNumber,
			LinkPrecedence: models.Primary,
		}
		if err := database.DB.Create(&contact).Error; err != nil {
			return nil
		}
		existingContact = append(existingContact, contact)
		primaryID := FindPrimaryContactID(&existingContact)
		return &primaryID
	}
	return nil
}

func checkTwoPrimaryContactPresent(existingContact []models.Contact) *uint {
	primaryIDs := []uint{}
	for _, c := range existingContact {
		if c.LinkPrecedence == models.Primary {
			primaryIDs = append(primaryIDs, c.ID)
		}
	}

	// If exactly two primary contacts, demote the one with smaller ID
	if len(primaryIDs) == 2 {
		// Find smaller and larger ID
		var minID, maxID uint
		if primaryIDs[0] < primaryIDs[1] {
			minID = primaryIDs[0]
			maxID = primaryIDs[1]
		} else {
			minID = primaryIDs[1]
			maxID = primaryIDs[0]
		}

		// Update the smaller ID to secondary and link it to the larger primary
		if err := database.DB.Model(&models.Contact{}).
			Where("id = ?", maxID).
			Updates(map[string]interface{}{
				"link_precedence": models.Secondary,
				"linked_id":       minID,
			}).Error; err != nil {
			return nil
		}
		return &minID
	}
	return nil
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

func checkIfSameContactPresent(contactDetails *models.ContactDetails, existingContact *[]models.Contact) *uint {
	for _, c := range *existingContact {
		if ((c.Email != nil && contactDetails.Email != nil && *c.Email == *contactDetails.Email) &&
			(c.PhoneNumber != nil && contactDetails.PhoneNumber != nil && *c.PhoneNumber == *contactDetails.PhoneNumber)) ||
			contactDetails.Email == nil || contactDetails.PhoneNumber == nil {
			primaryID := FindPrimaryContactID(existingContact)
			return &primaryID
		}
	}
	return nil
}
