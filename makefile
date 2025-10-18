run_devmode:
	@echo "Running the application..."
	go run cmd/server/main.go
IMAGE_NAME=supermario64bit/ats_resume_analyzer_backend

build_docker:
	@read -p "Enter tag (e.g., v1.0.0 or latest): " TAG; \
	echo "Building image $(IMAGE_NAME):$$TAG"; \
	docker build -t $(IMAGE_NAME):$$TAG .; \
	echo "Build complete!"

# Push the image to Docker Hub
push_docker:
	@read -p "Enter tag to push: " TAG; \
	echo "Pushing image $(IMAGE_NAME):$$TAG to Docker Hub..."; \
	docker push $(IMAGE_NAME):$$TAG; \
	echo "Push complete!"

run_worker:
	@echo "⚙️ Running the Asynq worker in dev mode..."
	go run cmd/worker/main.go

run_asynqmon_docker:
	@echo "🌐 Starting AsynqMon UI..."
	docker run -d \
		--name asynqmon \
		-p 8081:8080 \
		--rm \
		hibiken/asynqmon \
		--redis-addr=redis:6379
	@echo "✅ AsynqMon started at http://localhost:8081"
stop_asynqmon_docker:
	@echo "🛑 Stopping AsynqMon..."
	@docker rm -f asynqmon 2>/dev/null || true
	@echo "✅ AsynqMon stopped"
