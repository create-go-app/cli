.PHONY: security test

security:
	@gosec ./...

test:
	@go test -v -cover ./...

run:
	@./dist/cgapp_darwin_amd64/cgapp create -p ./app -f react-js -b echo -d postgres

install: generate
	@go install -i ./...

generate:
	@go generate ./...

release:
	@git tag -a $(TAG) -m "Bump $(TAG)" && git push origin $(TAG)
	@goreleaser --rm-dist

release-test:
	@goreleaser --snapshot --skip-publish --rm-dist

delete-tag:
	@git tag --delete $(TAG)