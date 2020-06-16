# Define colors
GREEN=\033[0;32m
NOCOLOR=\033[0m

# Define app variables
NAME=cgapp
APP=./app
BUILD=./build

.PHONY: clean

clean:
	@rm -rf $(BUILD) $(APP) **/.DS_Store
	@echo "$(GREEN)[OK]$(NOCOLOR) Project was cleaned!"

test:
	@go test -v -cover -race ./...
	@echo "$(GREEN)[OK]$(NOCOLOR) Project was tested!"

install: generate
	@go install ./...
	@echo "$(GREEN)[OK]$(NOCOLOR) Project was installed to GOPATH/bin folder!"

run:
	@$(BUILD)/darwin/$(NAME) create -p $(APP)

generate:
	@go generate ./...
	@echo "$(GREEN)[OK]$(NOCOLOR) Embed configs was generated!"

security:
	@gosec ./...
	@echo "$(GREEN)[OK]$(NOCOLOR) Go security check was completed!"

build: clean generate security
	@CGO_ENABLED=0 GOARCH=amd64
	@GOOS=darwin go build -o $(BUILD)/darwin/$(NAME) ./cmd/$(NAME)/*.go
	@echo "$(GREEN)[OK]$(NOCOLOR) App backend for macOS x64 was builded!"
	@GOOS=linux go build -o $(BUILD)/linux/$(NAME) ./cmd/$(NAME)/*.go
	@echo "$(GREEN)[OK]$(NOCOLOR) App backend for GNU/Linux x64 was builded!"
	@GOOS=windows go build -ldflags="-H windowsgui" -o $(BUILD)/windows/$(NAME).exe ./cmd/$(NAME)/*.go
	@echo "$(GREEN)[OK]$(NOCOLOR) App backend for MS Windows x64 was builded!"
