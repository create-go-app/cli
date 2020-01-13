CGAPP=./build/macOS/cgapp

.PHONY: clean

clean:
	@rm -rf ./build ./app **/.DS_Store
	@echo "[OK] Project was cleaned!"

test:
	@go test ./...
	@echo "[OK] Project was tested!"

install:
	@go install ./...
	@echo "[OK] Project was installed to GOPATH/bin folder!"

run:
	@rm -rf ./app
	@$(CGAPP) start -p ./app

run-ex:
	@rm -rf ~/Downloads/app
	@$(CGAPP) start -p ~/Downloads/app -b net/http -f preact -w nginx

build-macosx: clean
	@GOOS=darwin GOARCH=amd64 go build -o $(CGAPP) ./cmd/cgapp/...
	@echo "[OK] Build for macOS (amd64) complete!"
