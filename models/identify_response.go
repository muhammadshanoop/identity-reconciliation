package models

type IdentifyReponse struct {
	ContactResponse ContactDetailResponse `json:"contact"`
}

type ContactDetailResponse struct {
	PrimaryContactID    uint     `json:"primaryContatctId"`
	Emails              []string `json:"emails"`
	PhoneNumbers        []string `json:"phoneNumbers"`
	SecondaryContactIDs []uint   `json:"secondaryContactIds"`
}
