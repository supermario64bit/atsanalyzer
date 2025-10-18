package routes

import (
	"github.com/hibiken/asynq"
	"github.com/supermario64bit/atsanalyzer/jobs/task_handler"
)

func RegisterHandlers(mux *asynq.ServeMux) {
	mux.HandleFunc(ResumeAnalyseRoute, task_handler.HandleResumeAnalysisTask)
}
