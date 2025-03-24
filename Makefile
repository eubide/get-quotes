.PHONY: build run clean deploy deploy-typinator

# Build the application
build:
	go build -ldflags="-s -w" -o bin/get-quote cmd/get-quote/main.go

# Run the application with quotes file
run: build
	./bin/get-quote -c get-quote.yaml quotes
	./bin/get-quote -c get-quote.yaml citas

# Clean build artifacts
clean:
	rm -rf bin/
	rm -rf *.log


deploy:build
	cp bin/get-quote ~/bin/get-quote

deploy-typinator:deploy
	ln -sf ~/bin/get-quote /Users/eubide/Library/Application\ Support/Typinator/Sets/Includes/Scripts/get-quote