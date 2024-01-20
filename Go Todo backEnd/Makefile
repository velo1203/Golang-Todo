FUNCTION_NAME	=superrichquiz-octet
BUILD					=$(CURDIR)/build
BIN						=$(CURDIR)/bin
MAIN					=$(CURDIR)/cmd/main.go
BUCKET 				=AWS_BUCKET_NAME
REGION				=ap-northeast-2
ARCH					=amd64
PROD_ZIPFILE	=$(FUNCTION_NAME)-prod-linux-$(ARCH).zip
PROD_HANDLER 	=bootstrap
DEV_ZIPFILE		=$(FUNCTION_NAME)-dev-linux-$(ARCH).zip
DEV_HANDLER 	=bootstrap

default: help

deps: ## install dependency
	@echo "\033[32mDependency ...\033[0m"
	@go install gorm.io/gorm
.PHONY: deps

env.local: ## copy env for local
	@cp configs/common.local configs/common.go

env.dev: ## copy env for development
	@cp configs/common.dev configs/common.go

env.prod: ## copy env for production
	@cp configs/common.prod configs/common.go

run: env.local ## run local
	@echo "\033[32mRunning ...\033[0m"
	@go run $(MAIN)
.PHONY: run

fmt: ## show formatting
	@echo "\033[32mfmt ...\033[0m"
	@gofmt -s -w . && go mod tidy
.PHONY: fmt

lint: fmt ## linting
	golangci-lint run ./...

build: deps ## build
	@echo "\033[32mBuilding ...\033[0m"
	@go build -o $(BIN) $(MAIN)
.PHONY: build

install: ## install
	@echo "\033[32mInstalling ...\033[0m"
	go install -v
.PHONY: install

test: ## testing
	@echo "\033[32mTesting...\033[0m"
	@go test ./... -v
	@echo "all tests passed"
.PHONY: test

build.local: env.local ## build for local
	@echo "\033[32mBuilding for Local running ...\033[0m"
	@go build -o $(BUILD)/main $(MAIN)

build.dev: env.dev ## build for development
	@echo "\033[32mBuilding for AWS Lamdbda Development ...\033[0m"
	@GOOS=linux GOARCH=$(ARCH) go build -o $(BUILD)/$(DEV_HANDLER) $(MAIN)
	@cd $(BUILD) && \
		zip -9 $(DEV_ZIPFILE) $(DEV_HANDLER)

build.prod: env.prod ## build for production
	@echo "\033[32mBuilding for AWS Lamdbda Production ...\033[0m"
	@GOOS=linux GOARCH=$(ARCH) go build -o $(BUILD)/$(PROD_HANDLER) $(MAIN)
	@cd $(BUILD) && \
		zip -9 $(PROD_ZIPFILE) $(PROD_HANDLER)

deploy.dev: build.dev lambda.deploy.dev ## deploy for development
deploy.prod: build.prod lambda.deploy.prod ## deploy for production

lambda.deploy.dev: ## copy to S3 & update lambda function for development
	@echo "\033[32mDistribution Development : copy to S3 & update lambda function ...\033[0m"
	@aws s3 cp $(BUILD)/$(DEV_ZIPFILE) s3://$(BUCKET)
	@aws lambda update-function-code \
		--function-name $(FUNCTION_NAME)-dev \
		--s3-bucket $(BUCKET) \
		--s3-key $(DEV_ZIPFILE) \
		--region $(REGION) \
	 &2> /dev/null

lambda.deploy.prod: ## copy to S3 & update lambda function for production
	@echo "\033[32mDistribution Prodction : copy to S3 & update lambda function ...\033[0m"
	@aws s3 cp $(BUILD)/$(PROD_ZIPFILE) s3://$(BUCKET)
	@aws lambda update-function-code \
		--function-name $(FUNCTION_NAME) \
		--s3-bucket $(BUCKET) \
		--s3-key $(PROD_ZIPFILE) \
		--region $(REGION) \
	 &2> /dev/null


docker.build.local: ## build docker for local running
	@echo "\033[32mDocker for local ...\033[0m"
	@docker compose -f docker/local/docker-compose.yml build --no-cache
	@docker compose -f docker/local/docker-compose.yml create
.PHONY: docker.build.local

docker.start.local: ## start docker for local
	@echo "\033[32mDocker start ...\033[0m"
	@docker compose -f docker/local/docker-compose.yml up -d
.PHONY: docker.start.local

docker.stop.local:    ## stop docker for local
	@echo "\033[32mDocker stop ...\033[0m"
	@docker compose -f docker/local/docker-compose.yml down
.PHONY: docker.stop.local


docker.build.dev: ## build docker for development
	@echo "\033[32mDocker for development ...\033[0m"
	@docker compose -f docker/dev/docker-compose.yml build --no-cache
	@docker compose -f docker/dev/docker-compose.yml create
.PHONY: docker.build.dev

docker.start.dev: ## start docker for development
	@echo "\033[32mDocker start for development...\033[0m"
	@docker compose -f docker/dev/docker-compose.yml up -d
.PHONY: docker.start.dev

docker.stop.dev: ## stop docker for development
	@echo "\033[32mDocker stop for development...\033[0m"
	@docker compose -f docker/dev/docker-compose.yml down
.PHONY: docker.stop.dev


docker.build.prod: ## build docker for production
	@echo "\033[32mDocker for production ...\033[0m"
	@docker compose -f docker/prod/docker-compose.yml build --no-cache
	@docker compose -f docker/prod/docker-compose.yml create
.PHONY: docker.build.prod

docker.start.prod: ## start docker for production
	@echo "\033[32mDocker start for production...\033[0m"
	@docker compose -f docker/prod/docker-compose.yml up -d
.PHONY: docker.start.prod

docker.stop.prod: ## stop docker for production
	@echo "\033[32mDocker stop for production...\033[0m"
	@docker compose -f docker/prod/docker-compose.yml down
.PHONY: docker.stop.prod


clean: ## clean
	@echo "\033[32mCleaning...\033[0m"
	@go clean
	@rm -rf $(BIN)/*
	@rm -rf $(BUILD)/*
.PHONY: clean

help: ## Show help for each of the Makefile recipes.
	@grep -E '^[a-zA-Z0-9 -.]+:.*#'  $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.PHONY: help