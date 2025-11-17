package campaign_test

import (
	"emailn/internal/contract"
	"emailn/internal/domain/campaign"
	internalerrors "emailn/internal/internal-errors"
	internalmock "emailn/internal/test/mock"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	newCampaign = contract.NewCampaign{
		Name:      "Name test",
		Content:   "Body Hi!",
		Emails:    []string{"teste1@test.com"},
		CreatedBy: "test@test.com",
	}
	repositoryMock     *internalmock.RepositoryMock
	service            = campaign.ServiceImp{}
	pendingCampaign, _ = campaign.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)
	startedCampaing, _ = campaign.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)
)

func setUp() {
	repositoryMock = new(internalmock.RepositoryMock)
	service.Repository = repositoryMock
	startedCampaing.Status = campaign.Started
}

func setUpgetByMockAnythingReturnCampaign(campaign *campaign.Campaign) {
	repositoryMock.On("GetBy", mock.Anything).Return(campaign, nil)
}

func setupSendEmailWithErr(err error) {
	sendMail := func(campaign *campaign.Campaign) error {
		return err
	}
	service.SendMail = sendMail
}

func Test_Create_WithValidData_ReturnsID(t *testing.T) {
	setUp()
	repositoryMock.On("Create", mock.Anything).Return(nil)

	id, err := service.Create(newCampaign)

	assert.NotNil(t, id)
	assert.Nil(t, err)
}

func Test_Create_WithEmptyPayload_DoesNotReturnInternalError(t *testing.T) {
	_, err := service.Create(contract.NewCampaign{})

	assert.False(t, errors.Is(err, internalerrors.ErrInternal))
}

func Test_Create_WithValidData_CallsRepositoryCreate(t *testing.T) {
	setUp()
	repositoryMock.On("Create", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		if campaign.Name != newCampaign.Name ||
			campaign.Content != newCampaign.Content ||
			len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}

		return true
	})).Return(nil)

	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)
}

func Test_Create_WhenRepositoryFails_ReturnsInternalError(t *testing.T) {
	setUp()
	repositoryMock.On("Create", mock.Anything).Return(errors.New("error to save on database"))

	_, err := service.Create(newCampaign)

	assert.ErrorIs(t, internalerrors.ErrInternal, err)
}

func Test_GetBy_WhenExists_ReturnsCorrectInfo(t *testing.T) {
	setUp()
	repositoryMock.On("GetBy", mock.MatchedBy(func(id string) bool {
		return id == pendingCampaign.ID
	})).Return(pendingCampaign, nil)

	campaignReturned, _ := service.GetBy(pendingCampaign.ID)

	assert.Equal(t, pendingCampaign.ID, campaignReturned.ID)
	assert.Equal(t, pendingCampaign.Name, campaignReturned.Name)
	assert.Equal(t, pendingCampaign.Content, campaignReturned.Content)
	assert.Equal(t, pendingCampaign.Status, campaignReturned.Status)
}

func Test_GetBy_WhenRepoError_ReturnsInternalError(t *testing.T) {
	setUp()
	repositoryMock.On("GetBy", mock.Anything).Return(nil, errors.New("Something wrong'"))

	_, err := service.GetBy("Invalid_campaign")

	assert.Equal(t, internalerrors.ErrInternal.Error(), err.Error())
}

func Test_Start_WhenCampaignNotDefined_ReturnsError(t *testing.T) {
	setUp()
	repositoryMock.On("GetBy", mock.Anything).Return(nil, errors.New(""))

	err := service.Start("")

	assert.NotNil(t, err)
}

func Test_Start_WhenStatusNotPending_ReturnsError(t *testing.T) {
	setUp()
	setUpgetByMockAnythingReturnCampaign(startedCampaing)

	err := service.Start("campaignTest")

	assert.NotNil(t, err)
}

func Test_Start_WhenMailSent_UpdatesStatusToStarted(t *testing.T) {
	setUp()
	setUpgetByMockAnythingReturnCampaign(pendingCampaign)
	repositoryMock.On("Update", mock.MatchedBy(func(updatedCampaign *campaign.Campaign) bool {
		return pendingCampaign.ID == updatedCampaign.ID && updatedCampaign.Status == campaign.Started
	})).Return(nil)
	setupSendEmailWithErr(nil)
	service.Start("campaignTest")

	assert.Equal(t, campaign.Started, pendingCampaign.Status)
}

func Test_DeliverCampaignEmail_WhenFail_StatusIsFail(t *testing.T) {
	setUp()
	setupSendEmailWithErr(errors.New("error to send email"))
	repositoryMock.On("Update", mock.MatchedBy(func(campaignToUpdate *campaign.Campaign) bool {
		return pendingCampaign.ID == campaignToUpdate.ID && campaignToUpdate.Status == campaign.Fail
	})).Return(nil)

	service.DeliverCampaignEmail(pendingCampaign)

	repositoryMock.AssertExpectations(t)
}

func Test_DeliverCampaignEmail_WhenSuccess_StatusIsDone(t *testing.T) {
	setUp()
	setupSendEmailWithErr(nil)
	repositoryMock.On("Update", mock.MatchedBy(func(campaignToUpdate *campaign.Campaign) bool {
		return pendingCampaign.ID == campaignToUpdate.ID && campaignToUpdate.Status == campaign.Done
	})).Return(nil)

	service.DeliverCampaignEmail(pendingCampaign)

	repositoryMock.AssertExpectations(t)
}
