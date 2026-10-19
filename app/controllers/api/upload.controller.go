package api

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"peregerine/systems/types/responses"
	"peregerine/app/libraries"
	"peregerine/app/jobs/resume"
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

	// PDF parsing is done by the worker so the request returns immediately
	err = u.Queue.SetQueue(c.UserContext(), resume.TaskName, resume.Payload{
		Path: path,
		SaveAs: filename + ext,
		Original: file.Filename,
	})
	if err != nil {
		resp := u.ErrorResponse(false, err, "Failed to queue file processing")

		return c.Status(500).JSON(resp)
	}

	resp := u.SuccessResponse(true, fiber.Map{
		"orignal": file.Filename,
		"save_as": filename + ext,
		"save_path": path,
		"status": "processing",
	})

	return c.JSON(resp)
}

func (u UploadController) Update(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Index",
	})
}

// Detail returns the parsed text of an uploaded resume, ?id=<save_as>
func (u UploadController) Detail(c *fiber.Ctx) error {
	id := c.Query("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(
			u.ErrorResponse(false, nil, "id is required"),
		)
	}

	text, err := u.Redis.Get(c.UserContext(), resume.ResultGroup, id)
	if errors.Is(err, RedisService.ErrNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(
			u.ErrorResponse(false, err, "File not found or still processing"),
		)
	}
	if err != nil {
		return c.Status(500).JSON(
			u.ErrorResponse(false, err, "Redis is not available"),
		)
	}

	return c.JSON(u.SuccessResponse(true, fiber.Map{
		"save_as": id,
		"text": text,
	}))
}

func (u UploadController) Remove(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Index",
	})
}
