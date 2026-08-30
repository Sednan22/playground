.PHONY: all build run clean

all: build run

build:
	@go build -o test ./cmd/

run:
	@./test

clean:
	@rm -f test