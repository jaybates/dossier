package campaign

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Pull all campaigns from DB
func GetAllCampaigns(c *fiber.Ctx) error {
	err := c.SendString("GetAllCampaigns -> Campaign -> Handler")
	return err
}

// Pull campaign via UUID from DB
func GetCampaign(c *fiber.Ctx) error {
	msg := fmt.Sprintf("GetCampaign -> Campaign -> HandlerView Campaign: %s", c.Params("campaign"))
	err := c.SendString(msg)
	return err
}

// Pull all campaigns from DB
func CreateCampaign(c *fiber.Ctx) error {
	err := c.SendString("CreateCampaign -> Campaign -> Handler")
	return err
}

// update campaign via UUID from DB
func UpdateCampaign(c *fiber.Ctx) error {
	msg := fmt.Sprintf("UpdateCampaign -> Campaign -> HandlerUpate Campaign: %s", c.Params("campaign"))
	err := c.SendString(msg)
	return err
}

// Pull campaign via UUID from DB
func DeleteCampaign(c *fiber.Ctx) error {
	msg := fmt.Sprintf("DeleteCampaign -> Campaign -> Handler Delete Campaign: %s", c.Params("campaign"))
	err := c.SendString(msg)
	return err
}
