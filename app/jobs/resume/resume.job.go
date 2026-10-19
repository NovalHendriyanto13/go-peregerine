package resume

import (
	"fmt"
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"go.uber.org/zap"
	"peregerine/app/libraries"
	RedisType "peregerine/systems/services/redis/types"
	LoggerService "peregerine/systems/services/logger"
)

const TaskName = "resume:parse"

// ResultGroup is the redis hash where parsed resume text is stored, keyed by SaveAs
const ResultGroup = "resumes"

// Payload is the task payload sent by the upload controller
type Payload struct {
	Path string `json:"path"`
	SaveAs string `json:"save_as"`
	Original string `json:"original"`
}

type ResumeHandler struct {
	redis RedisType.SettingService
	logger *LoggerService.Logger
}

func NewResumeHandler(redis RedisType.SettingService, logger *LoggerService.Logger) *ResumeHandler {
	return &ResumeHandler{ redis: redis, logger: logger }
}

func (h *ResumeHandler) Process(ctx context.Context, t *asynq.Task) error {
	var payload Payload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		// A malformed payload will never succeed, so don't retry it
		return fmt.Errorf("invalid payload: %v: %w", err, asynq.SkipRetry)
	}

	text, err := libraries.ReadPdfToText(payload.Path)
	if err != nil {
		return err
	}

	if err := h.redis.Set(ctx, ResultGroup, payload.SaveAs, text); err != nil {
		return err
	}

	h.logger.Log.Info("resume parsed", zap.String("file", payload.SaveAs))

	return nil
}
