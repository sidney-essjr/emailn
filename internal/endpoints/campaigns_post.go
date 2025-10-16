package endpoints

import (
	"emailn/internal/contract"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, r *http.Request) (any, int, map[string]string) {
	var request contract.NewCampaign
	err := render.DecodeJSON(r.Body, &request)

	if err != nil {
		render.Status(r, 400)
		render.JSON(w, r, err.Error())
	}

	id, mapErr := h.CampaignService.Create(request)
	return map[string]string{"id": id}, 201, mapErr
}
