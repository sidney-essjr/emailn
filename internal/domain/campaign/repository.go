package campaign

type Repository interface {
	Save(campaign *Campaign) (string, error)
	Get() map[string]Campaign
	GetById(id string) (*Campaign, error)
}
