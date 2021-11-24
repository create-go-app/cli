.PHONY: clean lint security critic test install build release

clean:
	rm -rf ./tmp coverage.out

lint:
	golangci-lint run ./...

security:
	gosec -quiet ./...

critic:
	gocritic check -enableAll ./...

test: clean lint security critic
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

install: test
	CGO_ENABLED=0 go build -ldflags="-s -w" -o $(GOPATH)/bin/cgapp ./cmd/cgapp/main.go

build: test
	goreleaser --snapshot --skip-publish --rm-dist

release: test
	git tag -a $(VERSION) -m "$(VERSION)"
	goreleaser --snapshot --skip-publish --rm-dist

delete-tag:
	git tag --delete $(VERSION)

update-pkg-cache:
	curl -i https://proxy.golang.org/github.com/create-go-app/cli/v3/@v/$(VERSION).info
