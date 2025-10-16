package campaign

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
)

var (
	name    = "Campaign X"
	content = "Body Content"
	emails  = []string{"email1@email.com", "email2@email.com"}
	faker   = gofakeit.New(0)
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	//arrange
	assert := assert.New(t)

	//act
	campaign, _ := NewCampaign(name, content, emails)

	//assert
	assert.NotEmpty(campaign.Id)
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(emails))
	assert.Equal(StatusPending, campaign.Status)
	assert.True(campaign.Status.IsValid())
}

func Test_NewCampaign_IdIsNotNil(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, emails)

	assert.NotNil(campaign.Id)
}

func Test_NewCampaign_CreatedOnMustBeNow(t *testing.T) {
	assert := assert.New(t)
	now := time.Now()

	campaign, _ := NewCampaign(name, content, emails)

	assert.WithinDuration(now, campaign.CreatedOn, time.Second)
}

func Test_NewCampaign_MustValidateRequiredNameMin(t *testing.T) {
	assert := assert.New(t)

	_, mapErr := NewCampaign("", content, emails)

	assert.NotNil(mapErr)
	assert.Equal("Key: 'Campaign.Name' Error:Field validation for 'Name' failed on the 'min' tag", mapErr["Name"])
}

func Test_NewCampaign_MustValidateRequiredNameMax(t *testing.T) {
	assert := assert.New(t)

	_, mapErr := NewCampaign(faker.LetterN(25), content, emails)

	assert.NotNil(mapErr)
	assert.Equal("Key: 'Campaign.Name' Error:Field validation for 'Name' failed on the 'max' tag", mapErr["Name"])
}

func Test_NewCampaign_MustValidateRequiredContentMin(t *testing.T) {
	assert := assert.New(t)

	_, mapErr := NewCampaign(name, "", emails)

	assert.NotNil(mapErr)
	assert.Equal("Key: 'Campaign.Content' Error:Field validation for 'Content' failed on the 'min' tag", mapErr["Content"])
}

func Test_NewCampaign_MustValidateRequiredContentMax(t *testing.T) {
	assert := assert.New(t)

	_, mapErr := NewCampaign(name, faker.LetterN(1025), emails)

	assert.NotNil(mapErr)
	assert.Equal("Key: 'Campaign.Content' Error:Field validation for 'Content' failed on the 'max' tag", mapErr["Content"])
}

func Test_NewCampaign_MustValidateRequiredEmailList(t *testing.T) {
	assert := assert.New(t)

	_, mapErr := NewCampaign(name, content, []string{})

	assert.NotNil(mapErr)
	assert.Equal("Key: 'Campaign.Contacts' Error:Field validation for 'Contacts' failed on the 'min' tag", mapErr["Contacts"])
}
