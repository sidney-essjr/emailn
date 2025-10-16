package campaign

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internal-errors"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

var (
	newCampaign = contract.NewCampaign{
		Name:    "Name X",
		Content: "Content X",
		Emails:  []string{"email1@email.com", "email2@email.com"},
	}
	repository = new(repositoryMock)
	service    = CampaingService{Repository: repository}
)

func (r *repositoryMock) Save(campaign *Campaign) (string, error) {
	args := r.Called(campaign)
	return args.String(0), args.Error(1)
}

func (r *repositoryMock) Get() map[string]Campaign {
	// args := r.Called(campaign)
	return nil
}

func (r *repositoryMock) GetById(id string) (*Campaign, error) {
	// args := r.Called(campaign)
	return &Campaign{}, nil
}

func onSave(campaign *Campaign) bool {
	return campaign.Name == newCampaign.Name &&
		campaign.Content == newCampaign.Content &&
		len(campaign.Contacts) == len(newCampaign.Emails)
}

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	repository.On("Save", mock.Anything).Return("wer34", nil)

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_ValidadeDomainError(t *testing.T) {
	assert := assert.New(t)
	newCampaign.Name = ""

	_, mapErr := service.Create(newCampaign)

	assert.NotNil(mapErr)
	assert.Equal("Key: 'Campaign.Name' Error:Field validation for 'Name' failed on the 'min' tag", mapErr["Name"])
}

func Test_Create_SaveCampaign(t *testing.T) {
	repository.On("Save", mock.MatchedBy(onSave)).Return("", nil)
	newCampaign.Name = "Name X"
	service.Create(newCampaign)

	repository.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repository.On("Save", mock.MatchedBy(onSave)).Return("", errors.New("internal server error"))

	_, mapErr := service.Create(newCampaign)

	assert.Equal(internalerrors.ErrInternal.Error(), mapErr["error"])
}
