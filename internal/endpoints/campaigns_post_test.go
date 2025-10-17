package endpoints

import (
	"bytes"
	"emailn/internal/contract"
	"emailn/internal/domain/campaign"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type serviceMock struct {
	mock.Mock
	campaign.Repository
}

var (
	reqBody = contract.NewCampaign{
		Name:    "Name X",
		Content: "Content X",
		Emails:  []string{"email1@email.com", "email2@email.com"},
	}
	service = new(serviceMock)
)

func (s *serviceMock) Create(newCampaign contract.NewCampaign) (string, map[string]string) {
	id := "2"
	errors := map[string]string{}
	return id, errors
}

func (s *serviceMock) GetBy(id string) (*contract.CampaignResponse, map[string]string) {
	return nil, nil
}

func Test_campaign_post_should_create_new_campaing(t *testing.T) {
	assert := assert.New(t)
	expectedId := map[string]string{"id": "2"}
	service.On("Create", mock.Anything).Return(expectedId, map[string]string{})

	handler := Handler{CampaignService: service}
	jsonData, _ := json.Marshal(reqBody)
	buffer := bytes.NewBuffer(jsonData)
	req, _ := http.NewRequest("POST", "/", buffer)
	res := httptest.NewRecorder()
	id, status, mapErr := handler.CampaignPost(res, req)

	assert.Equal(expectedId, id)
	assert.Equal(201, status)
	assert.Equal(len(mapErr), 0)
}
