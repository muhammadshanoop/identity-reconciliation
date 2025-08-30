package services

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadshanoop/identity-reconciliation/helpers"
	"github.com/muhammadshanoop/identity-reconciliation/models"
	"github.com/muhammadshanoop/identity-reconciliation/validators"
)

type UserService interface {
	ReconcileUser(c *gin.Context)
}
type Helper interface {
	FindOrCreateContact(contactDetails *models.ContactDetails) (*uint, error)
	GetAllLinkedContacts(primaryContactID uint) (*[]models.Contact, error)
}
type Validator interface {
	ValidateRequest(contactDetails *models.ContactDetails) error
}
type DefaultHelper struct{}
type DefaultValidator struct{}

func (DefaultHelper) FindOrCreateContact(contactDetails *models.ContactDetails) (*uint, error) {
	return helpers.FindOrCreateContact(contactDetails)
}
func (DefaultHelper) GetAllLinkedContacts(primaryContactID uint) (*[]models.Contact, error) {
	return helpers.GetAllLinkedContacts(primaryContactID)
}
func (DefaultValidator) ValidateRequest(contactDetails *models.ContactDetails) error {
	return validators.ValidateRequest(contactDetails)
}

type userServiceImpl struct {
	helper    Helper
	validator Validator
}

func NewUserService(helper Helper, validator Validator) UserService {
	return &userServiceImpl{helper, validator}
}

// global default instance
var UserServiceInstance UserService = NewUserService(DefaultHelper{}, DefaultValidator{})

func (u *userServiceImpl) ReconcileUser(c *gin.Context) {
	var contactDetails models.ContactDetails
	if err := c.ShouldBindJSON(&contactDetails); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	err := u.validator.ValidateRequest(&contactDetails)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	primaryContactID, err := u.helper.FindOrCreateContact(&contactDetails)
	if err != nil || primaryContactID == nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response := u.formatResponse(*primaryContactID)
	c.JSON(200, gin.H{"message": response})
}

func (u *userServiceImpl) formatResponse(primaryContactID uint) models.IdentifyReponse {
	var contactResponse models.ContactDetailResponse
	contactResponse.PrimaryContactID = primaryContactID
	linkedContacts, _ := u.helper.GetAllLinkedContacts(primaryContactID)

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
