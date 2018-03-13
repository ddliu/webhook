COMMIT = $(shell git describe --always)
run:
	go run webhook.go
test:
	go test -v -cover ./...
vendor:
	dep ensure -v
release:
	goreleaser --rm-dist