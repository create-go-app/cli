# A default Makefile for Create Go App project.
# Author: Vic Shóstak <truewebartisans@gmail.com> (https://1wa.co)

.PHONY: test run build

FRONTEND_PATH = $(PWD)/frontend
BACKEND_PATH = $(PWD)/backend

test:
	@if [ -d "$(FRONTEND_PATH)" ]; then cd $(FRONTEND_PATH) && npm run test; fi
	@if [ -d "$(BACKEND_PATH)" ]; then cd $(BACKEND_PATH) && go test ./...; fi

run: test
	@if [ -d "$(FRONTEND_PATH)" ]; then cd $(FRONTEND_PATH) && $(MAKE) run; fi
	@if [ -d "$(BACKEND_PATH)" ]; then cd $(BACKEND_PATH) && $(MAKE) run; fi

build: test
	@if [ -d "$(FRONTEND_PATH)" ]; then cd $(FRONTEND_PATH) && $(MAKE) build; fi
	@if [ -d "$(BACKEND_PATH)" ]; then cd $(BACKEND_PATH) && $(MAKE) build; fi
