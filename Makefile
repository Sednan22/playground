.PHONY: all build run clean

all: build run

build:
	@go build -o test

run:
	@./test

clean:
	@rm -f test