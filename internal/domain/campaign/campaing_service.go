package campaign

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internal-errors"
)

type CampaingService struct {
	Repository
}

func (s *CampaingService) Create(newCampaign contract.NewCampaign) (string, map[string]string) {
	campaign, mapErr := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)

	if len(mapErr) > 0 {
		return "", mapErr
	}

	id, repositoryErr := s.Repository.Save(campaign)

	if repositoryErr != nil {
		return "", map[string]string{"error": internalerrors.ErrInternal.Error()}
	}

	return id, nil
}

func (s *CampaingService) GetBy(id string) (*contract.CampaignResponse, map[string]string) {
	campaign, err := s.GetById(id)

	if err != nil {
		return nil, map[string]string{"error": err.Error()}
	}

	return &contract.CampaignResponse{
		Id:      campaign.Id,
		Name:    campaign.Name,
		Content: campaign.Content,
		Status:  string(campaign.Status),
	}, nil
}
