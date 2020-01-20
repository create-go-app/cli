package configs

// EditorConfig is .editorconfig file
var EditorConfig string = (`# .editorconfig by Vic Shóstak <truewebartisans@gmail.com> (https://1wa.co)
root = true

[*]
indent_style = space
indent_size = 2
charset = utf-8
trim_trailing_whitespace = true
insert_final_newline = true

[{Dockerfile,Makefile,go.mod,go.sum,*.go}]
indent_style = tab
indent_size = 4

[{*.yml,*.yaml}]
indent_style = space
indent_size = 4
`)

// GitIgnore is .gitignore file
var GitIgnore string = (`# .gitignore by Vic Shóstak <truewebartisans@gmail.com> (https://1wa.co)
# macOS
.DS_store

# Dev builds
**/app/
**/build/

# Node.js dependencies
**/node_modules/
`)

// MakeFile is Makefile with run/build/install instructions
var MakeFile string = (`# Makefile by Vic Shóstak <truewebartisans@gmail.com> (https://1wa.co)
# Define colors
GREEN=\033[0;32m
NOCOLOR=\033[0m

# Define app variables
BACKEND=./backend
FRONTEND=./frontend
WEBSERVER=./webserver
NAME=apiserver

.PHONY: deploy

deploy:
	@docker-compose up --build --force-recreate
	@echo "$(GREEN)[OK]$(NOCOLOR) App (dev) was deployed!"

deploy-prod:
	@docker-compose -f docker-compose.yml -f docker-compose.prod.yml up --build --force-recreate -d
	@echo "$(GREEN)[OK]$(NOCOLOR) App (prod) was deployed!"

backend-run:
	@cd $(BACKEND)
	@go run ./cmd/$(NAME)/*.go

backend-test:
	@cd $(BACKEND)
	@go test ./internal/$(NAME)/*.go
	@echo "$(GREEN)[OK]$(NOCOLOR) Project was tested!"

backend-build-darwin:
	@GOOS=darwin GOARCH=amd64
	@go build -o $(BACKEND)/build/backend $(BACKEND)/cmd/$(NAME)/*.go
	@echo "$(GREEN)[OK]$(NOCOLOR) App backend for macOS x64 was builded!"

backend-build-linux:
	@GOOS=linux GOARCH=amd64
	@go build -o $(BACKEND)/build/backend $(BACKEND)/cmd/$(NAME)/*.go
	@echo "$(GREEN)[OK]$(NOCOLOR) App backend for GNU/Linux x64 was builded!"

backend-build-windows:
	@GOOS=windows GOARCH=amd64
	@go build -ldflags="-H windowsgui" -o $(BACKEND)/build/backend.exe $(BACKEND)/cmd/$(NAME)/*.go
	@echo "$(GREEN)[OK]$(NOCOLOR) App backend for MS Windows x64 was builded!"

frontend-run:
	@cd $(FRONTEND)
	@npm start

frontend-build:
	@cd $(FRONTEND)
	@npm run build
	@echo "$(GREEN)[OK]$(NOCOLOR) App frontend was builded!"

certbot:
	@chmod +x $(WEBSERVER)/scripts/register_ssl_for_domain.sh
	@sudo $(WEBSERVER)/scripts/register_ssl_for_domain.sh \
	--domains $(DOMAINS) --email $(EMAIL) --data-path $(DATA_PATH) --staging 1

certbot-prod:
	@chmod +x $(WEBSERVER)/scripts/register_ssl_for_domain.sh
	@sudo $(WEBSERVER)/scripts/register_ssl_for_domain.sh \
	--domains $(DOMAINS) --email $(EMAIL) --data-path $(DATA_PATH) --staging 0
`)
