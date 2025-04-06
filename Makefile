.PHONY: build install

build:
	go build -o galaxygoat ./cmd/galaxygoat

install:
	go install ./cmd/galaxygoat
