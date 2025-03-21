.PHONY: build run clean test test-verbose test-coverage test-html test-race test-pkg test-randomline test-quotereader

# Build the application
build:
	go build -ldflags="-s -w" -o bin/get-quote cmd/get-quote/main.go

# Run the application with quotes file
run: build
	./bin/get-quote quotes

# Clean build artifacts
clean:
	rm -rf bin/
	rm -rf coverage.out
	rm -rf coverage.html
	rm -rf *.log


deploy:build
	cp bin/get-quote ~/bin/get-quote

deploy-typinator:deploy
	ln -sf ~/bin/get-quote /Users/eubide/Library/Application\ Support/Typinator/Sets/Includes/Scripts/get-quote