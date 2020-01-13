CGAPP=./build/macOS/cgapp

.PHONY: clean

clean:
	@rm -rf ./build ./app **/.DS_Store
	@echo "[OK] Project was cleaned!"

test:
	@go test ./internal/cgapp/*.go
	@echo "[OK] Project was tested!"

install:
	@go install ./cmd/cgapp/*.go
	@echo "[OK] Project was installed to GOPATH/bin folder!"

build-macosx: clean
	@GOOS=darwin GOARCH=amd64 go build -o $(CGAPP) ./cmd/cgapp/*.go
	@echo "[OK] Build for macOS (amd64) complete!"

run-init:
	@rm -rf ./app
	@$(CGAPP) init -p ./app -b net/http -f preact

run-docker:
	@rm -rf ./app
	@$(CGAPP) docker -p ./app nginx

run-ex-init:
	@rm -rf ~/Downloads/app
	@$(CGAPP) init -p ~/Downloads/app -b net/http -f preact

run-ex-docker:
	@$(CGAPP) docker -p ~/Downloads/app nginx
