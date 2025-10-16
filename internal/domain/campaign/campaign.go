package campaign

import (
	internalerrors "emailn/internal/internal-errors"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string `validate:"email"`
}

type Status string

const (
	StatusPending Status = "pending"
	StatusSending Status = "sending"
	StatusDone    Status = "done"
)

type Campaign struct {
	Id        string    `validate:"required"`
	Name      string    `validate:"min=5,max=24"`
	CreatedOn time.Time `validate:"required"`
	Content   string    `validate:"min=5,max=1024"`
	Contacts  []Contact `validate:"min=1,dive"`
	Status
}

func NewCampaign(name string, content string, emails []string) (*Campaign, map[string]string) {

	contacts := make([]Contact, len(emails))
	for index, value := range emails {
		contacts[index].Email = value
	}

	campaign := &Campaign{
		Id:        xid.New().String(),
		Name:      name,
		CreatedOn: time.Now(),
		Content:   content,
		Contacts:  contacts,
		Status:    StatusPending,
	}

	mapErr := internalerrors.ValidateStruct(campaign)
	if len(mapErr) > 0 {
		return nil, mapErr
	}

	return campaign, nil
}

func (s Status) IsValid() bool {
	switch s {
	case StatusPending, StatusSending, StatusDone:
		return true
	default:
		return false
	}
}
