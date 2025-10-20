package campaign

import "emailn/internal/contract"

type Service interface {
	Create(newCampaign contract.NewCampaign) (string, map[string]string)
	GetBy(id string) (*contract.CampaignResponse, map[string]string)
	Get() []*contract.CampaignResponse
}
