REPO=github.com/nstoker/congenial-memory

setup: run_services
	@go run ./cmd/db/setup.go

run_services:
	@docker-compose up --build -d

run_server:
	@MONGO_URL=mongodb://mongo_user:mongo_secret@0.0.0.0:27017/kudos PORT=:444 go run cmd/main.go

run_client:
	@bin/bash -c ""cd $$GOPATH/src/$$REPO/pkg/http/web/app && yarn serve"
