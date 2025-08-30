package services

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/muhammadshanoop/identity-reconciliation/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// --- Mock dependencies ---
type mockHelper struct {
	findFunc func(contactDetails *models.ContactDetails) (*uint, error)
	getFunc  func(primaryContactID uint) (*[]models.Contact, error)
}

func (m *mockHelper) FindOrCreateContact(contactDetails *models.ContactDetails) (*uint, error) {
	return m.findFunc(contactDetails)
}
func (m *mockHelper) GetAllLinkedContacts(primaryContactID uint) (*[]models.Contact, error) {
	return m.getFunc(primaryContactID)
}

type mockValidator struct {
	validateFunc func(contactDetails *models.ContactDetails) error
}

func (m *mockValidator) ValidateRequest(contactDetails *models.ContactDetails) error {
	return m.validateFunc(contactDetails)
}

func TestReconcileUser_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockH := &mockHelper{
		findFunc: func(contactDetails *models.ContactDetails) (*uint, error) {
			id := uint(1)
			return &id, nil
		},
		getFunc: func(primaryContactID uint) (*[]models.Contact, error) {
			email := "test@gmail.com"
			phoneNumber := "1234567891"
			contacts := []models.Contact{
				{ID: 1, Email: &email, PhoneNumber: &phoneNumber, LinkPrecedence: models.Primary},
			}
			return &contacts, nil
		},
	}
	mockV := &mockValidator{
		validateFunc: func(contactDetails *models.ContactDetails) error {
			return nil
		},
	}

	services := NewUserService(mockH, mockV)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/api/identify", strings.NewReader(`{"email":"test@example.com"}`))
	c.Request.Header.Set("Content-Type", "application/json")

	services.ReconcileUser(c)
	assert.Equal(t, http.StatusOK, w.Code)
	var got struct {
		Message models.IdentifyReponse `json:"message"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &got)
	require.NoError(t, err)

	expected := models.IdentifyReponse{
		ContactResponse: models.ContactDetailResponse{
			PrimaryContactID:    1,
			Emails:              []string{"test@gmail.com"},
			PhoneNumbers:        []string{"1234567891"},
			SecondaryContactIDs: nil,
		},
	}

	assert.Equal(t, expected, got.Message)
}

func TestReconcileUser_InvalidInput(t *testing.T) {
	gin.SetMode(gin.TestMode)

	services := NewUserService(&mockHelper{}, &mockValidator{})
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest("POST", "/api/identify", strings.NewReader(`{"email": null,"phoneNumber": 11}`))
	c.Request.Header.Set("Content-Type", "application/json")

	services.ReconcileUser(c)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid input")
}

func TestReconcileUser_ValidationError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockV := &mockValidator{
		validateFunc: func(cd *models.ContactDetails) error {
			return errors.New("validation failed")
		},
	}
	service := NewUserService(&mockHelper{}, mockV)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/api/identify", strings.NewReader(`{"email": null,"phoneNumber": null}`))
	c.Request.Header.Set("Content-Type", "application/json")

	service.ReconcileUser(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "validation failed")
}
