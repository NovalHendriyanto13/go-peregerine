package api

import (
	"github.com/gofiber/fiber/v2"
	"peregerine/systems/types/responses"
	"peregerine/app/libraries"
	"path/filepath"
	RedisService "peregerine/systems/services/redis/types"
	QueueService "peregerine/systems/services/queue/types"
)

type UploadController struct {
	responses.BaseResponse
	Redis RedisService.SettingService
	Queue QueueService.QueueMethod
}

func (u UploadController) Index(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Index",
	})
}

func (u UploadController) Create(c *fiber.Ctx) error {
	file, err := c.FormFile("resume")
	if err != nil {
		resp := u.ErrorResponse(false, err, "Please upload file")

		return c.Status(500).JSON(resp)
	}

	filename := libraries.RandomString(8)
	ext := filepath.Ext(file.Filename)

	path := "./storage/public/resumes/" + filename + ext
	if err := c.SaveFile(file, path); err != nil {
		resp := u.ErrorResponse(false, err, "There an error when uploading file")

		return c.Status(500).JSON(resp)
	}

	readText, err := libraries.ReadPdfToText(path)
	if err != nil {
		resp := u.ErrorResponse(false, err, "Error Reading file")

		return c.Status(500).JSON(resp)
	}

	u.Queue.SetQueue("testing", fiber.Map{
		"testing": file.Filename,
	})

	resp := u.SuccessResponse(true, fiber.Map{
		"orignal": file.Filename,
		"save_as": filename + ext,
		"save_path": path,
		"text": readText,
	})

	return c.JSON(resp)
}

func (u UploadController) Update(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Index",
	})
}

func (u UploadController) Detail(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Index",
	})
}

func (u UploadController) Remove(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Index",
	})
}