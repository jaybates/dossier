package router

import (
	"github.com/gofiber/fiber/v2"
	campaign "github.com/jaybates/dossier/handler"
)

func SetupRoutes(app *fiber.App) {

	// Setup up Group Routes

	campaigns := app.Group("campaigns")

	// Route Group - campaigns [START]

	campaigns.Get("/", campaign.GetAllCampaigns)

	campaigns.Get("/create", campaign.CreateCampaign)

	campaigns.Get("/update/:campaign", campaign.UpdateCampaign)

	campaigns.Get("/delete/:campaign", campaign.DeleteCampaign)

	campaigns.Get("/:campaign", campaign.GetCampaign)

	// Route Group - campaigns [START]

}
