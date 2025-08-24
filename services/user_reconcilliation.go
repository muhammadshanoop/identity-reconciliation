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

	primaryContactID := findPrimaryContactID(contacts)
	linkedContacts, err := helpers.GetAllLinkedContacts(primaryContactID)
	if err != nil {
		log.Fatalf("Error in fetching linked contacts :%v", err)
	}

	response := formatResponse(primaryContactID, linkedContacts)
	c.JSON(200, gin.H{"message": response})
}

func findPrimaryContactID(contacts *[]models.Contact) uint {
	var primaryContactID uint
	for _, contact := range *contacts {
		if contact.LinkPrecedence == models.Primary {
			primaryContactID = contact.ID
			break
		}
	}
	return primaryContactID
}

func formatResponse(primaryContactID uint, linkedContacts *[]models.Contact) models.IdentifyReponse {
	var contactResponse models.ContactDetailResponse
	contactResponse.PrimaryContactID = primaryContactID
	for _, contact := range *linkedContacts {
		contactResponse.Emails = append(contactResponse.Emails, *contact.Email)
		contactResponse.PhoneNumbers = append(contactResponse.PhoneNumbers, *contact.PhoneNumber)
		if contact.LinkPrecedence == models.Secondary {
			contactResponse.SecondaryContactIDs = append(contactResponse.SecondaryContactIDs, contact.ID)
		}
	}
	return models.IdentifyReponse{
		ContactResponse: contactResponse,
	}
}
