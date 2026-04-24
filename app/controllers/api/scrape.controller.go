package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/go-rod/rod"
	"peregerine/systems/types/responses"
	"peregerine/app/requests/scrapes"
	"peregerine/configs"
	logger "peregerine/systems/services/logger"
)

type ScrapeController struct {
	responses.BaseResponse
	Logger *logger.Logger
}

func (s ScrapeController) Index(c *fiber.Ctx) error {
	var req requests.ScrapeRequest

	if err := c.BodyParser(&req); err != nil {
		resp := s.ErrorResponse(false, err, "Invalid Request")
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	browser := rod.New().
		ControlURL(configs.ChromeUrl).
		MustConnect()
	defer browser.MustClose()

	page := browser.MustPage(req.UrlReq)

	page.MustWaitLoad()

	title := page.MustInfo().Title
	page.MustElement(`flt-semantics[arial-label="Login With Microsoft"]`).
		MustClick()

	page.MustWaitNavigation()
	fmt.Println("Redirect to: ", page.MustInfo().URL)

	page.MustWaitLoad()
	

	// el, err := page.Element("h1")
	// if err != nil {
	// 	r
	// }

	resp := s.SuccessResponse(true, fiber.Map{
		"title": title,
	})

	return c.JSON(resp)
}

func (s ScrapeController) Create(c *fiber.Ctx) error {
	return nil
}

func (s ScrapeController) Detail(c *fiber.Ctx) error {
	return nil
}

func (s ScrapeController) Update(c *fiber.Ctx) error {
	return nil
}

func (s ScrapeController) Remove(c *fiber.Ctx) error {
	return nil
}