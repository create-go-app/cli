.PHONY: test security install build release

test:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

security:
	gosec -quiet ./...

install: security test
	CGO_ENABLED=0 go build -ldflags="-s -w" -o $(GOPATH)/bin/cgapp main.go

build: security test
	goreleaser --snapshot --skip-publish --rm-dist

release: security test
	git tag -a $(VERSION) -m "$(VERSION)"
	goreleaser --snapshot --skip-publish --rm-dist

delete-tag:
	git tag --delete $(VERSION)

update-pkg-cache:
	curl -i https://proxy.golang.org/github.com/create-go-app/cli/@v/$(VERSION).info
