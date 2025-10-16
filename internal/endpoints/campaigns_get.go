package endpoints

import (
	"net/http"
)

func (h *Handler) CampaignGet(w http.ResponseWriter, r *http.Request) (any, int, map[string]string) {
	return h.CampaignService.Get(), 200, nil
}
