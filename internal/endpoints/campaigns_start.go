package endpoints

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) CampaignStart(w http.ResponseWriter, r *http.Request) (any, int, error) {
	id := chi.URLParam(r, "id")
	err := h.CampaignService.Start(id)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return nil, 200, nil
}
