test:
	go test -race ./...

coverage:
	go test -coverprofile=cover.out ./...
	go tool cover -html  cover.out

lint:
	gometalinter --disable-all --deadline=300s \
		--enable=vet         --enable=vetshadow   --enable=golint \
		--enable=staticcheck --enable=ineffassign --enable=errcheck \
		--enable=unconvert   --enable=deadcode \
		--exclude=go/pkg/mod \
		./...
