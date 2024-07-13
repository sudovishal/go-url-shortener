.PHONY: build run dev docker-build docker-run

build:
    go build -o server ./cmd/server

run: build
    ./server

dev:
    air