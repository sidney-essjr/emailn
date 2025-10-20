package campaign

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internal-errors"
)

type CampaignService struct {
	repository Repository
}

func NewCampaignService(repository Repository) Service {
	return &CampaignService{
		repository: repository,
	}
}

func (s *CampaignService) Create(newCampaign contract.NewCampaign) (string, map[string]string) {
	campaign, mapErr := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)

	if len(mapErr) > 0 {
		return "", mapErr
	}

	id, repositoryErr := s.repository.Save(campaign)

	if repositoryErr != nil {
		return "", map[string]string{"error": internalerrors.ErrInternal.Error()}
	}

	return id, nil
}

func (s *CampaignService) GetBy(id string) (*contract.CampaignResponse, map[string]string) {
	campaign, err := s.repository.GetById(id)

	if err != nil {
		return nil, map[string]string{"error": internalerrors.ErrInternal.Error()}
	}

	return &contract.CampaignResponse{
		Id:      campaign.Id,
		Name:    campaign.Name,
		Content: campaign.Content,
		Status:  string(campaign.Status),
	}, nil
}

func (s *CampaignService) Get() []*contract.CampaignResponse {
	campaigns := s.repository.Get()
	result := make([]*contract.CampaignResponse, 0, len(campaigns))

	for _, campaign := range campaigns {
		result = append(result, &contract.CampaignResponse{
			Id:      campaign.Id,
			Name:    campaign.Name,
			Content: campaign.Content,
			Status:  string(campaign.Status),
		})
	}

	return result
}
