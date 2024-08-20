.PHONY: clean security critic test install build release build-and-push-images delete-tag update-pkg-cache

clean:
	rm -rf ./tmp ./tests

security:
	go run github.com/securego/gosec/v2/cmd/gosec@latest -quiet ./...

critic:
	go run github.com/go-critic/go-critic/cmd/gocritic@latest check -enableAll ./...

test: clean security critic
	mkdir ./tests
	go test -coverprofile=./tests/coverage.out ./...
	go tool cover -func=./tests/coverage.out
	rm -rf ./tests

install: test
	CGO_ENABLED=0 go build -ldflags="-s -w" -o $(GOPATH)/bin/cgapp ./cmd/cgapp/main.go

build: test
	goreleaser --snapshot

release: test
	git tag -a v$(VERSION) -m "$(VERSION)"
	goreleaser --snapshot

build-and-push-images: test
	docker build -t docker.io/koddr/cgapp:latest .
	docker push docker.io/koddr/cgapp:latest
	docker build -t docker.io/koddr/cgapp:$(VERSION) .
	docker push docker.io/koddr/cgapp:$(VERSION)
	docker image rm docker.io/koddr/cgapp:$(VERSION)

update-pkg-cache:
	curl -i https://proxy.golang.org/github.com/create-go-app/cli/v4/@v/v$(VERSION).info

delete-tag:
	git tag --delete v$(VERSION)
	docker image rm docker.io/koddr/cgapp:latest
	docker image rm docker.io/koddr/cgapp:$(VERSION)
