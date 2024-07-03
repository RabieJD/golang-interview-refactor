VERSION=1.0.0

# Go parameters
GOCMD=go
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOINSTALL=$(GOCMD) install
GORUN=$(GOCMD) run
GOVET=$(GOCMD) vet
GOFMT=$(GOCMD) fmt
GOMOD=$(GOCMD) mod
SERVER_PATH= cmd/web-api/main.go
ENV_VARS_TEST := $(shell cat .env.test | grep -v '^#' | xargs)
ENV_VARS_PROD := $(shell cat .env.prod | grep -v '^#' | xargs)


init:
	$(GOMOD) download

run-prod-env:
	@$(ENV_VARS_PROD) $(GORUN) ${SERVER_PATH}

run-test-env:
	@$(ENV_VARS_TEST)  $(GORUN) ${SERVER_PATH}

tidy:
	$(GOMOD) tidy

fmt:
	$(GOFMT) ./...

vet:
	$(GOVET) ./...

tool-moq:
	$(GOINSTALL) github.com/matryer/moq@latest

moq:
	moq -skip-ensure -out pkg/testdata/service/calculator/zmoq_dependencies.go -pkg calculator pkg/service/calculator CartRepo PriceRepo
	moq -skip-ensure -out pkg/testdata/controllers/zmoq_dependencies.go -pkg contyrollers pkg/controllers CalculatorHandler
test:
	$(GOTEST) -v ./...

test-race:
	$(GOTEST) --race -v ./...

clean-test-cache:
	$(GOCLEAN) -testcache

clean-mod-cache:
	$(GOCLEAN) -modcache

docker-compose-run:
	docker compose -f docker/docker-compose.yml up -d --build

docker-compose-down:
	docker compose -f docker/docker-compose.yml down