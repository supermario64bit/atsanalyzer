package jobs

import (
	"log"
	"os"

	"github.com/hibiken/asynq"
	"github.com/joho/godotenv"
	"github.com/supermario64bit/atsanalyzer/jobs/routes"
)

func main() {
	err := godotenv.Load(".env.worker")
	if err != nil {
		log.Fatal(err)
	}

	redisAddr := os.Getenv("REDIS_ADDR")
	// Connect to Redis
	server := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{
			Concurrency: 10, // process up to 10 jobs at once
		},
	)

	// Create a mux and register handlers
	mux := asynq.NewServeMux()
	routes.RegisterHandlers(mux)

	log.Println("🚀 Asynq Worker started...")
	if err := server.Run(mux); err != nil {
		log.Fatalf("❌ Worker crashed: %v", err)
	}
}
