docker:
	docker rm -v -f $$(docker ps -qa) || echo "No containers found. Skipping removal."
	docker-compose up

server:
	go run cmd/main.go

watch: 
	air

test:
	go vet ./...
	go test  -v -coverpkg ./internal/... -coverprofile coverage.out -covermode count ./internal/...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html

proto:
	go get github.com/macgeargear/kokoro-go-proto@latest

swagger:
	swag init -d ./internal/file -g ../../cmd/main.go -o ./docs -md ./docs/markdown --parseDependency --parseInternal