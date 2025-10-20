package endpoints

import (
	"context"
	"emailn/internal/contract"
	"emailn/internal/domain/campaign"
	internalerrors "emailn/internal/internal-errors"
	mock_test "emailn/internal/mock-test"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	campaignResponse = contract.CampaignResponse{
		Id:      "8ausewf",
		Name:    "Name X",
		Content: "Content X",
		Status:  string(campaign.StatusPending),
	}
)

func Test_Campaigns_GetById_return_Status_Ok(t *testing.T) {
	//arrange
	assert := assert.New(t)
	service := new(mock_test.ServiceMock)
	handler := Handler{CampaignService: service}
	service.On("GetBy", mock.Anything).Return(&campaignResponse, nil)

	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "123234")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	//act
	obj, status, mapErr := handler.CampaignGetById(res, req)

	//assert
	assert.Equal(http.StatusOK, status)
	assert.Equal(0, len(mapErr))
	assert.Equal(campaignResponse.Id, obj.(*contract.CampaignResponse).Id)
}

func Test_Campaigns_GetById_return_Status_BadRequest(t *testing.T) {
	//arrange
	assert := assert.New(t)
	service := new(mock_test.ServiceMock)
	handler := Handler{CampaignService: service}
	service.On("GetBy", mock.Anything).Return(&campaignResponse, nil)

	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	//act
	obj, status, mapErr := handler.CampaignGetById(res, req)

	//assert
	assert.Equal(http.StatusBadRequest, status)
	assert.Equal("missing id parameter", mapErr["error"])
	assert.Equal(nil, obj)
}

func Test_Campaigns_GetById_return_Status_NotFound(t *testing.T) {
	//arrange
	assert := assert.New(t)
	service := new(mock_test.ServiceMock)
	handler := Handler{CampaignService: service}
	service.On("GetBy", mock.Anything).Return(nil, map[string]string{"error": internalerrors.ErrInternal.Error()})

	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "123234")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	//act
	obj, status, mapErr := handler.CampaignGetById(res, req)

	//assert
	assert.Equal(http.StatusNotFound, status)
	assert.Equal(internalerrors.ErrInternal.Error(), mapErr["error"])
	assert.Equal(nil, obj)
}
