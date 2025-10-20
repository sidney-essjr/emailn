package campaign

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internal-errors"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

var (
	campaignAllFilds = Campaign{
		Id:        "8ausewf",
		Name:      "Name X",
		Content:   "Content X",
		Status:    StatusPending,
		CreatedOn: time.Now(),
		Contacts:  []Contact{{Email: "email1@email.com"}, {Email: "email2@email.com"}},
	}
	newCampaign = contract.NewCampaign{
		Name:    "Name X",
		Content: "Content X",
		Emails:  []string{"email1@email.com", "email2@email.com"},
	}
	responseCampaign = contract.CampaignResponse{
		Id:      "8ausewf",
		Name:    "Name X",
		Content: "Content X",
		Status:  string(StatusPending),
	}
)

func newService() (repo *repositoryMock, service Service) {
	repo = new(repositoryMock)
	service = NewCampaignService(repo)
	return
}

func (r *repositoryMock) Save(campaign *Campaign) (string, error) {
	args := r.Called(campaign)
	return args.String(0), args.Error(1)
}

func (r *repositoryMock) Get() map[string]Campaign {
	// args := r.Called(campaign)
	return nil
}

func (r *repositoryMock) GetById(id string) (*Campaign, error) {
	args := r.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Campaign), args.Error(1)
}

func onSave(campaign *Campaign) bool {
	return campaign.Name == newCampaign.Name &&
		campaign.Content == newCampaign.Content &&
		len(campaign.Contacts) == len(newCampaign.Emails)
}

func onGetBy(id string) bool {
	return id == responseCampaign.Id
}

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	repo, service := newService()
	repo.On("Save", mock.Anything).Return("wer34", nil)

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_ValidadeDomainError(t *testing.T) {
	assert := assert.New(t)
	_, service := newService()
	newCampaign.Name = ""

	_, mapErr := service.Create(newCampaign)

	assert.NotNil(mapErr)
	assert.Equal("Key: 'Campaign.Name' Error:Field validation for 'Name' failed on the 'min' tag", mapErr["Name"])
}

func Test_Create_SaveCampaign(t *testing.T) {
	repo, service := newService()
	repo.On("Save", mock.MatchedBy(onSave)).Return("", nil)
	newCampaign.Name = "Name X"
	service.Create(newCampaign)

	repo.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repo, service := newService()
	repo.On("Save", mock.Anything).Return("", errors.New("internal server error"))

	id, mapErr := service.Create(newCampaign)

	assert.Equal("", id)
	assert.Equal(internalerrors.ErrInternal.Error(), mapErr["error"])
}

func Test_GetBy_ReturnResponseCampaign(t *testing.T) {
	assert := assert.New(t)
	repo, service := newService()
	id := "8ausewf"
	repo.On("GetById", mock.MatchedBy(onGetBy)).Return(&campaignAllFilds, nil)

	campaign, _ := service.GetBy(id)

	assert.Equal(&responseCampaign, campaign)
}

func Test_GetBy_ReturnError(t *testing.T) {
	assert := assert.New(t)
	repo, service := newService()
	id := "8ausewf"
	repo.On("GetById", mock.Anything).Return(nil, errors.New("something error"))

	_, mapErr := service.GetBy(id)

	assert.Equal(internalerrors.ErrInternal.Error(), mapErr["error"])
}
