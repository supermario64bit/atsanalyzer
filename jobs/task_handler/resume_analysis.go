package task_handler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/supermario64bit/atsanalyzer/dto"
	"github.com/supermario64bit/atsanalyzer/service"
)

func HandleResumeAnalysisTask(ctx context.Context, t *asynq.Task) error {
	var p dto.ResumeRequest
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	svc := service.NewResumeAnalysisService()
	_, appErr := svc.Analyse(&p)

	return appErr.Error
}
