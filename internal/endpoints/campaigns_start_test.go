package endpoints

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Start_200(t *testing.T) {
	campaignId := "testeId"
	setup()

	service.On("Start", mock.MatchedBy(func(id string) bool {
		return id == campaignId
	})).Return(nil)

	rr, req := httpTest("PATCH", "/")
	req = addParameter("id", campaignId, req)

	_, status, err := handler.CampaignStart(rr, req)

	assert.Equal(t, 200, status)
	assert.Nil(t, err)
}

func Test_Start_Err(t *testing.T) {
	expectedErr := errors.New("generic err")
	setup()

	service.On("Start", mock.Anything).Return(expectedErr)

	rr, req := httpTest("PATCH", "/")

	_, status, err := handler.CampaignStart(rr, req)

	assert.NotEqual(t, 200, status)
	assert.Equal(t, expectedErr, err)
}
