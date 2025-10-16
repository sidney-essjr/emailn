package database

import (
	"emailn/internal/domain/campaign"
	"fmt"
)

type CampaignRepository struct {
	Campaigns map[string]campaign.Campaign
}

func (r *CampaignRepository) Save(campaign *campaign.Campaign) (string, error) {
	r.Campaigns[campaign.Id] = *campaign
	return campaign.Id, nil
}

func (r *CampaignRepository) Get() map[string]campaign.Campaign {
	return r.Campaigns
}

func (r *CampaignRepository) GetById(id string) (*campaign.Campaign, error) {
	if c, ok := r.Campaigns[id]; ok {
		return &c, nil
	}
	return nil, fmt.Errorf("campaign with id %s not found", id)
}
