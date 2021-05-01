.PHONY: clean test security install build release

clean:
	rm -rf ./tmp coverage.out

test: clean
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

lint:
	golangci-lint run

security:
	gosec -quiet ./...

install: security lint test
	CGO_ENABLED=0 go build -ldflags="-s -w" -o $(GOPATH)/bin/cgapp ./cmd/cgapp/main.go

build: security test
	goreleaser --snapshot --skip-publish --rm-dist

release: security test
	git tag -a $(VERSION) -m "$(VERSION)"
	goreleaser --snapshot --skip-publish --rm-dist

delete-tag:
	git tag --delete $(VERSION)

update-pkg-cache:
	curl -i https://proxy.golang.org/github.com/create-go-app/cli/@v/$(VERSION).info
