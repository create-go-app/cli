.PHONY: clean lint security critic test install build release build-and-push-images delete-tag update-pkg-cache

clean:
	rm -rf ./tmp coverage.out

lint:
	$(GOPATH)/bin/golangci-lint run ./...

security:
	$(GOPATH)/bin/gosec -quiet ./...

critic:
	$(GOPATH)/bin/gocritic check -enableAll ./...

test: clean lint security critic
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

install: test
	CGO_ENABLED=0 go build -ldflags="-s -w" -o $(GOPATH)/bin/cgapp ./cmd/cgapp/main.go

build: test
	$(GOPATH)/bin/goreleaser --snapshot --skip-publish --rm-dist

release: test
	git tag -a v$(VERSION) -m "$(VERSION)"
	$(GOPATH)/bin/goreleaser --snapshot --skip-publish --rm-dist

build-and-push-images: test
	podman build -t docker.io/koddr/cgapp:latest .
	podman push docker.io/koddr/cgapp:latest
	podman build -t docker.io/koddr/cgapp:$(VERSION) .
	podman push docker.io/koddr/cgapp:$(VERSION)
	podman image rm docker.io/koddr/cgapp:$(VERSION)

update-pkg-cache:
	curl -i https://proxy.golang.org/github.com/create-go-app/cli/v3/@v/v$(VERSION).info

delete-tag:
	git tag --delete v$(VERSION)
	podman image rm docker.io/koddr/cgapp:latest
	podman image rm docker.io/koddr/cgapp:$(VERSION)