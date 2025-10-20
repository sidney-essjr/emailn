package mock_test

import (
	"emailn/internal/contract"

	"github.com/stretchr/testify/mock"
)

type ServiceMock struct {
	mock.Mock
}

func (s *ServiceMock) Create(newCampaign contract.NewCampaign) (string, map[string]string) {
	args := s.Called(newCampaign)

	if args.Get(1) == nil {
		return args.String(0), nil
	}
	return args.String(0), args.Get(1).(map[string]string)
}

func (s *ServiceMock) GetBy(id string) (*contract.CampaignResponse, map[string]string) {
	args := s.Called(id)

	if args.Get(0) == nil {
		return nil, args.Get(1).(map[string]string)
	}
	return args.Get(0).(*contract.CampaignResponse), nil
}

func (s *ServiceMock) Get() []*contract.CampaignResponse {
	args := s.Called()

	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).([]*contract.CampaignResponse)
}
