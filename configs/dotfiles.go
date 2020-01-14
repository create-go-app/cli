package configs

// EditorConfig is .editorconfig file
var EditorConfig string = (`# .editorconfig by Create Go App Authors
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
`)

// GitIgnore is .gitignore file
var GitIgnore string = (`# .gitignore by Create Go App Authors
# macOS
.DS_store

# Dev builds
**/build/

# Node.js dependencies
**/node_modules/
`)

// MakeFile is Makefile with run/build/install instructions
var MakeFile string = (`# Makefile by Create Go App Authors
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
	@echo "$(GREEN)[OK]$(NOCOLOR) App backend for GNU/Linux x64 was builded!"

build-frontend:
	@cd $(FRONTEND)
	@npm run build
	@echo "$(GREEN)[OK]$(NOCOLOR) App frontend was builded!"
`)
