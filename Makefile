MONGO_URL=mongodb://mongo_user:mongo_secret@0.0.0.0:27017
PORT=:4444
REPO=github.com/nstoker/congenial-memory

setup: run_services
	MONGO_URL=${MONGO_URL} go run ./cmd/db/setup.go

run_services:
	@docker-compose up --build -d

run_server:
	@MONGO_URL=${MONGO_URL} PORT=${PORT} go run cmd/main.go

run_client:
	/bin/bash -c "cd ${GOPATH}/src/${REPO}/pkg/http/web/app && yarn serve"
