package campaign

import "emailn/internal/contract"

type Service interface {
	Create(newCampaign contract.NewCampaign) (string, map[string]string)
	Repository
	GetBy(id string) (*contract.CampaignResponse, map[string]string)
}
