# Makefile by Create Go App Authors
# Define colors
GREEN=\033[0;32m
NOCOLOR=\033[0m

# Define app variables
BACKEND=./backend
FRONTEND=./frontend

.PHONY: run-backend

run-backend:
	@go run $(BACKEND)/...

run-frontend:
	@cd $(FRONTEND)
	@npm start

build-backend-darwin:
	@cd ./backend
	@GOOS=darwin GOARCH=amd64
	@go build -o $(BACKEND)/build/backend $(BACKEND)/...
	@echo "$(GREEN)[OK]$(NOCOLOR) App backend for macOS x64 was builded!"

build-backend-linux:
	@cd ./backend
	@GOOS=linux GOARCH=amd64
	@go build -o $(BACKEND)/build/backend $(BACKEND)/...
	@echo "$(GREEN)[OK]$(NOCOLOR) App backend for GNU/Linux x64 was builded!"

build-backend-windows:
	@cd ./backend
	@GOOS=windows GOARCH=amd64
	@go build -ldflags="-H windowsgui" -o $(BACKEND)/build/backend $(BACKEND)/...
	@echo "$(GREEN)[OK]$(NOCOLOR) App backend for MS Windows x64 was builded!"

build-frontend:
	@cd $(FRONTEND)
	@npm run build
	@echo "$(GREEN)[OK]$(NOCOLOR) App frontend was builded!"
