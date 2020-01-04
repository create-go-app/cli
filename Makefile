CGAPP=./build/macOS/cgapp

.PHONY: run

run:
	rm -rf ./app \
	&& $(CGAPP) -b echo -f preact -p ./app

run-external:
	rm -rf ~/Downloads/app \
	&& $(CGAPP) -b echo -f github.com/koddr/sweetconfirm.js -p ~/Downloads/app

test:
	go test ./cmd/cgapp/*.go
	@echo "[✔️] Project was tested!"

clean:
	rm -rf ./build ./app
	@echo "[✔️] Project was cleaned!"

build-macosx:
	rm -rf ./build ./app ./configs/**/.DS_Store \
	&& pkger \
	&& GOOS=darwin GOARCH=amd64 go build -o $(CGAPP) ./*.go
	@echo "[✔️] Build for macOS (amd64) complete!"
