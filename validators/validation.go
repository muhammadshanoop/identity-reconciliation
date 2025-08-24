package validators

import (
	"errors"
	"fmt"

	"github.com/muhammadshanoop/identity-reconciliation/models"
)

func ValidateRequest(contactDetails *models.ContactDetails) error {
	// validation: at least one must be non-nil
	if contactDetails.Email == nil && contactDetails.PhoneNumber == nil {
		return errors.New("either email or phoneNumber must be provided")
	}

	// Safe logging / debug
	if contactDetails.Email != nil {
		fmt.Println("Email:", *contactDetails.Email)
	}
	if contactDetails.PhoneNumber != nil {
		fmt.Println("Phone:", *contactDetails.PhoneNumber)
	}

	return nil
}
