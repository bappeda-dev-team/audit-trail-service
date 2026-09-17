APP_NAME=audit-trail-service

.PHONY: build run clean test swagger build-image build-docker format check-format dev

build:
	@echo ">> building app ${APP_NAME}"
	@mkdir -p bin
	go build -o ./bin/$(APP_NAME) ./cmd/api
	@echo ">> done build in ./bin/${APP_NAME}"

build-docker:
	@mkdir -p bin
	CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" \
	-o ./bin/$(APP_NAME) ./cmd/api

format:
	goimports -w .

check-format:
	@test -z "$$(gofmt -l .)" || (echo "gofmt required:" && gofmt -l . && exit 1)
	@test -z "$$(goimports -l .)" || (echo "goimports required:" && goimports -l . && exit 1)

run:
	go run ./cmd/api

dev: format run

test:
	go test -race ./...

swagger:
	swag init -d cmd/api,internal/api,internal/web,internal/audit/domain,internal/audit/web

clean:
	rm -f ./bin/$(APP_NAME)

build-image:
	@echo ">> building docker with tag ${APP_NAME}:latest"
	@docker build . -t $(APP_NAME)
