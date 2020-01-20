# Makefile by Vic Shóstak <truewebartisans@gmail.com> (https://1wa.co)
# Define colors
GREEN=\033[0;32m
NOCOLOR=\033[0m

# Define app variables
BACKEND=./backend
FRONTEND=./frontend
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
	@cd webserver/scripts
	@chmod +x ./register_ssl_for_domain.sh
	@./register_ssl_for_domain.sh \
	--domains $(DOMAINS) --email $(EMAIL) --data-path $(DATA_PATH) --staging 1

certbot-prod:
	@cd webserver/scripts
	@chmod +x ./register_ssl_for_domain.sh
	@./register_ssl_for_domain.sh \
	--domains $(DOMAINS) --email $(EMAIL) --data-path $(DATA_PATH) --staging 0
