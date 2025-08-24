package services

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/muhammadshanoop/identity-reconciliation/helpers"
	"github.com/muhammadshanoop/identity-reconciliation/models"
	"github.com/muhammadshanoop/identity-reconciliation/validators"
)

func ReconcileUser(c *gin.Context) {
	var contactDetails models.ContactDetails
	if err := c.ShouldBindJSON(&contactDetails); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	err := validators.ValidateRequest(&contactDetails)
	if err != nil {
		log.Fatalf("Error in validating request : %v", err)
	}
	//create contact
	contacts, err := helpers.FindOrCreateContact(&contactDetails)
	if err != nil {
		log.Fatalf("Error in creating contact : %v", err)
	}

	primaryContactID := helpers.FindPrimaryContactID(contacts)
	linkedContacts, err := helpers.GetAllLinkedContacts(primaryContactID)
	if err != nil {
		log.Fatalf("Error in fetching linked contacts :%v", err)
	}

	response := formatResponse(primaryContactID, linkedContacts)
	c.JSON(200, gin.H{"message": response})
}

func formatResponse(primaryContactID uint, linkedContacts *[]models.Contact) models.IdentifyReponse {
	var contactResponse models.ContactDetailResponse
	contactResponse.PrimaryContactID = primaryContactID
	emailsSeen := make(map[string]bool)
	phonesSeen := make(map[string]bool)
	for _, contact := range *linkedContacts {
		if contact.Email != nil && !emailsSeen[*contact.Email] {
			contactResponse.Emails = append(contactResponse.Emails, *contact.Email)
			emailsSeen[*contact.Email] = true
		}

		if contact.PhoneNumber != nil && !phonesSeen[*contact.PhoneNumber] {
			contactResponse.PhoneNumbers = append(contactResponse.PhoneNumbers, *contact.PhoneNumber)
			phonesSeen[*contact.PhoneNumber] = true
		}

		if contact.LinkPrecedence == models.Secondary {
			contactResponse.SecondaryContactIDs = append(contactResponse.SecondaryContactIDs, contact.ID)
		}
	}
	return models.IdentifyReponse{
		ContactResponse: contactResponse,
	}
}
