package jobs

import (
	"os"

	"github.com/hibiken/asynq"
)

func NewAsynqClient() *asynq.Client {
	redisAddr := os.Getenv("REDIS_ADDR")
	return asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
}
