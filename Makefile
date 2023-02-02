.PHONY: up
up:
	@docker-compose  up -d
run: up
	@go run main.go