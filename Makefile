prep: dependencies generate format lint test
	@echo "\033[1;32mREADY TO MERGE\033[0m"

format:
	@echo "\033[1;36mFormatting code...\033[0m"
	@go fmt ./...

lint:
	@echo "\033[1;36mLinting code...\033[0m"
	@golangci-lint run internal/...

test:
	@echo "\033[1;36mRunning tests...\033[0m"
	@go test ./... "-race" -coverprofile=coverage.txt -covermode=atomic

dependencies:
	@echo "\033[1;36mGet dependencies...\033[0m"
	@scripts/dev-dependencies.sh

generate: generate_openapi generate_protobuf
	@go mod tidy

generate_openapi:
	@echo "\033[1;36mGenerating OpenAPI routes, clients and types...\033[0m"
	@scripts/openapi.sh images internal/edge/ports/images ports false
	@scripts/openapi.sh image-builder internal/clients imagebuilder true

generate_protobuf:
	@echo "\033[1;36mGenerating protobuf files...\033[0m"
	@buf generate --path api --template api/protobuf/edge/buf.gen.yaml

generate_unit_tests:
	@echo "\033[1;36mGenerating unit tests...\033[0m"
	@find . -name '*.go' -not -name '*_test.go' -not -name 'openapi*.go' -not -path './internal/commomn/genproto/*' \
	| xargs gotests -all -w 

clean: clean_openapi clean_protobuf clean_unit_tests
	@go mod tidy

clean_openapi:
	@echo "\033[1;36mCleaning OpenAPI files...\033[0m"
	@find . -name 'openapi*.go' -delete

clean_protobuf:
	@echo "\033[1;36mCleaning protobuf files...\033[0m"
	@find . -name '*.pb.go' -delete

clean_unit_tests:
	@echo "\033[1;36mCleaning unit tests...\033[0m"
	@find . -name '*_test.go' -delete

clean_env:
	find . -name '*.db' -delete
	docker rm postgresql_database redis -f
	docker run --rm -d --name redis -p 6379:6379 redis
	docker run --rm -d --name postgresql_database \
	-e POSTGRESQL_USER=user -e POSTGRESQL_PASSWORD=pass \
	-e POSTGRESQL_DATABASE=db -p 5432:5432 registry.redhat.io/rhel8/postgresql-10
	go run ./internal/edge