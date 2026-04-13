package resume

import (
	"fmt"
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
)

const TaskName = "testing"

type ResumeHandler struct {}

func NewResumeHandler() *ResumeHandler {
	return &ResumeHandler{}
}

func (h *ResumeHandler) Process(c context.Context, t *asynq.Task) error {
	var payload map[string] interface{}
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return err
	}
	fmt.Println(payload)


	return nil
}