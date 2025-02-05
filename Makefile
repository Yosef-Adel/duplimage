build:
	@go build -o  bin/duplImage

run: build
	@./bin/duplImage -d ~/Pictures/Terminal
