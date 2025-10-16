package endpoints

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) CampaignGetById(w http.ResponseWriter, r *http.Request) (any, int, map[string]string) {
	id := chi.URLParam(r, "id")

	if id == "" {
		return nil, http.StatusBadRequest, map[string]string{"error": "missing id parameter"}
	}

	campaign, err := h.CampaignService.GetById(id)

	if err != nil {
		return nil, http.StatusNotFound, map[string]string{"error": err.Error()}
	}

	return campaign, http.StatusOK, nil
}
