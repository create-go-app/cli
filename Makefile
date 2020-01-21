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
	@go test ./...
	@echo "$(GREEN)[OK]$(NOCOLOR) Project was tested!"

install:
	@go install ./...
	@echo "$(GREEN)[OK]$(NOCOLOR) Project was installed to GOPATH/bin folder!"

run:
	@$(BUILD)/darwin/$(NAME) start -p $(APP)

generate:
	@go generate ./configs/defaults.go
	@echo "$(GREEN)[OK]$(NOCOLOR) Embed configs was generated!"

build: clean generate
	@CGO_ENABLED=0 GOARCH=amd64
	@GOOS=darwin go build -o $(BUILD)/darwin/$(NAME) ./cmd/$(NAME)/*.go
	@echo "$(GREEN)[OK]$(NOCOLOR) App backend for macOS x64 was builded!"
	@GOOS=linux go build -o $(BUILD)/linux/$(NAME) ./cmd/$(NAME)/*.go
	@echo "$(GREEN)[OK]$(NOCOLOR) App backend for GNU/Linux x64 was builded!"
	@GOOS=windows go build -ldflags="-H windowsgui" -o $(BUILD)/windows/$(NAME).exe ./cmd/$(NAME)/*.go
	@echo "$(GREEN)[OK]$(NOCOLOR) App backend for MS Windows x64 was builded!"
