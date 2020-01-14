package configs

// EditorConfig is .editorconfig file
var EditorConfig string = (`# .editorconfig by Create Go App authors (https://cgapp.1wa.co)
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
var GitIgnore string = (`# .gitignore by Create Go App authors (https://cgapp.1wa.co)
# macOS
.DS_store

# Dev builds
**/build/

# Node.js dependencies
**/node_modules/
`)

// MakeFile is Makefile with run/build/install instructions
var MakeFile string = (`# Makefile by Create Go App authors (https://cgapp.1wa.co)
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

build-backend:
	@cd $(BACKEND)
	@make build

build-frontend:
	@cd $(FRONTEND)
	@npm run build
	@echo "$(GREEN)[OK]$(NOCOLOR) Project frontend was builded!"
`)
