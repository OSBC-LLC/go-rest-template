build:
	@echo "Building go-rest-service"
	@go build -o bin/orch-rest-template -v .
	@echo "done."

test:
	@echo "running tests..."
	@go test -v ./tests -coverprofile=./coverage.out -coverpkg ./...

report:
	@echo "showing test coverage report..."
	@go tool cover -html=coverage.out

db:
	@echo "building docker images"
	@go mod tidy
	@go mod vendor
	@docker compose build && rm -rf vendor/
	@echo done.

du:
	@echo "starting docker images"
	@docker compose up

dp:
	@echo "starting ONLY docker postgres"
	@docker compose -f docker-compose.psql.yml up

dn:
	@echo "starting ONLY nginx proxy"
	@docker build -t nginx-proxy ./nginx && docker run -p 8880:80 nginx-proxy

da: db du

newtest:
	@read -p "name: " NAME \
	&& cd tests && ginkgo generate $${NAME}
