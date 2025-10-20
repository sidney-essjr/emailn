package main

import (
	"emailn/internal/domain/campaign"
	"emailn/internal/endpoints"
	"emailn/internal/infrastructure/database"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	service := campaign.NewCampaignService(&database.CampaignRepository{
		Campaigns: make(map[string]campaign.Campaign),
	})

	handler := endpoints.Handler{
		CampaignService: service,
	}

	r.Post("/campaigns", endpoints.HandlerError(handler.CampaignPost))
	r.Get("/campaigns", endpoints.HandlerError(handler.CampaignGet))
	r.Get("/campaigns/{id}", endpoints.HandlerError(handler.CampaignGetById))

	http.ListenAndServe(":3000", r)

}
