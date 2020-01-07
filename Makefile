CGAPP=./build/macOS/cgapp

.PHONY: clean

run-init:
	rm -rf ./app \
	&& $(CGAPP) init -p ./app -b net/http -f preact

run-docker:
	rm -rf ./app \
	&& $(CGAPP) docker -p ./app nginx

run-ex-init:
	rm -rf ~/Downloads/app \
	&& $(CGAPP) init -p ~/Downloads/app -b net/http -f preact

run-ex-docker:
	rm -rf ~/Downloads/app \
	&& $(CGAPP) docker -p ~/Downloads/app nginx

test:
	go test ./internal/cgapp/*.go
	@echo "[✔️] Project was tested!"

clean:
	rm -rf ./build ./app
	@echo "[✔️] Project was cleaned!"

build-macosx:
	rm -rf ./build ./app ./configs/**/.DS_Store \
	&& pkger \
	&& GOOS=darwin GOARCH=amd64 go build -o $(CGAPP) ./cmd/cgapp/*.go
	@echo "[✔️] Build for macOS (amd64) complete!"
