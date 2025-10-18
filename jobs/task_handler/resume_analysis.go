package task_handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hibiken/asynq"
	"github.com/supermario64bit/atsanalyzer/dto"
	"github.com/supermario64bit/atsanalyzer/service"
)

func HandleResumeAnalysisTask(ctx context.Context, t *asynq.Task) error {
	log.Println("Resume Analysis Task Handler started")
	var p dto.ResumeAnalysisJob
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	log.Println("Analaysing resume...")
	svc := service.NewResumeAnalysisService()
	_, appErr := svc.Analyse(&p)

	if appErr != nil {
		log.Println("Error occured while analysing. Error: " + appErr.Error.Error())
		return appErr.Error
	}
	log.Println("Resume Analysis Task Handler completed")
	return nil
}
