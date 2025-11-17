package endpoints

import (
	"context"
	internalmock "emailn/internal/test/mock"
	"net/http"
	"net/http/httptest"

	"github.com/go-chi/chi/v5"
)

var (
	service *internalmock.CampaignServiceMock
	handler = Handler{}
)

func setup() Handler {
	service = new(internalmock.CampaignServiceMock)
	handler = Handler{CampaignService: service}
	return handler
}

func httpTest(method, url string) (*httptest.ResponseRecorder, *http.Request) {
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest(method, url, nil)
	return rr, req
}

func addParameter(keyParameter, valueParameter string, req *http.Request) *http.Request {
	chiContext := chi.NewRouteContext()
	chiContext.URLParams.Add(keyParameter, valueParameter)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiContext))
	return req
}
