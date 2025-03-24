# Get Quote

A simple, elegant Go tool for extracting random quotes from text files. Built using clean architecture principles, it provides a lightweight command-line utility that can be easily integrated into scripts, terminal workflows, or other applications.

![Go Version](https://img.shields.io/badge/Go-1.16+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/license-MIT-blue.svg)

## Features

- 🚀 **Fast and lightweight**: Minimal dependencies, quick startup time
- 🔄 **Random selection**: Get a different quote each time
- 🔧 **Configurable**: Custom file locations, error messages, and more
- 🏠 **User-friendly configuration**: Supports user home directory with `~` notation
- 🌍 **Multilingual support**: Works with any text file format
- 🧩 **Clean architecture**: Built using hexagonal design and SOLID principles
- 🔌 **Extensible**: Easy to add new features or adapters

## Installation

### From Source

```bash
# Clone the repository
git clone https://github.com/eubide/get-quote.git
cd get-quote

# Build the executable
make build

# Optional: Install to your local bin directory
make deploy
```

### Using Go Install

```bash
go install github.com/eubide/get-quote/cmd/get-quote@latest
```

## Quick Start

```bash
# Get a random quote in English
./bin/get-quote quotes

# Get a random quote in Spanish
./bin/get-quote citas

# List available quote files
./bin/get-quote -l

# Show help
./bin/get-quote -h
```

## Usage

```bash
./bin/get-quote [options] <filename>
```

Where `<filename>` is the name of a file in the configured directory (default `src/files/`).
The `.lst` extension will be added automatically if not provided.

### Command-line Options

- `-c <file>`: Specify a custom configuration file path
- `-h, --help`: Show help message
- `-l, --list`: List available quote files

### Examples

```bash
# Extract a random quote from quotes.lst
./bin/get-quote quotes

# Extract a random quote from citas.lst
./bin/get-quote citas

# Use a specific configuration file
./bin/get-quote -c ~/my-custom-config.yaml quotes

# List all available quote files
./bin/get-quote -l
```

## Integration Examples

### Add to your shell prompt

Add to your `.bashrc` or `.zshrc`:

```bash
# Display a random quote every time you open a terminal
get-quote quotes
```

### Use with notification systems

```bash
# Send a random quote notification every hour (for Unix-like systems)
# Add to your crontab (crontab -e):
0 * * * * export DISPLAY=:0 && /usr/bin/notify-send "Quote of the Hour" "$(/path/to/get-quote quotes)"
```

### Use as a library

```go
package main

import (
	"fmt"
	"log"

	"github.com/eubide/get-quote/internal/adapters/secondary/config"
	"github.com/eubide/get-quote/internal/adapters/secondary/repository"
	"github.com/eubide/get-quote/internal/app/services"
)

func main() {
	// Initialize configuration
	cfg, err := config.NewYAMLConfigWithOrder("")
	if err != nil {
		log.Printf("Warning: Using default configuration: %v\n", err)
	}

	// Initialize repository
	repo := repository.NewFileRepository(cfg)

	// Initialize service
	service := services.NewQuoteService(repo, cfg)

	// Get a random quote
	quote, err := service.GetRandomQuote("quotes")
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	// Use the quote
	fmt.Println(quote.Text)
}
```

## Configuration

The application uses a YAML configuration file. It searches for configuration in the following order:

1. Path specified with the `-c` flag
2. In the user's config directory: `$HOME/.config/get-quote/get-quote.yaml`
3. In the user's home directory: `$HOME/.get-quote.yaml`
4. In the current working directory: `get-quote.yaml`

If no configuration file is found in any of these locations, default values are used.

### Configuration Example

```yaml
# Random Quote Configuration

filesBaseDir: src/files
defaultExtension: .lst
errorMessages:
  fileNotFound: "The file %s does not exist"
  fileOpenError: "Error opening the file: %v"
  missingParameter: "Usage: %s <filename>\nYou must provide a filename %s"
```

### Home Directory Support

The `filesBaseDir` setting can use the tilde character (`~`) to represent the user's home directory:

```yaml
# Will use the user's home directory directly
filesBaseDir: ~

# Will use $HOME/quotes directory
filesBaseDir: ~/quotes

# Will use $HOME/.config/get-quote/quotes directory
filesBaseDir: ~/.config/get-quote/quotes
```

Note: Only the exact tilde (`~`) or a tilde followed by a slash (`~/`) will be expanded. Paths like `~user` will be left as-is.

## Creating Your Own Quote Files

Quote files are simple text files with one quote per line. Create your own collections:

1. Create a new `.lst` file in your configured `filesBaseDir`
2. Add one quote per line
3. Access it using `get-quote your-file-name`

## Architecture

This project follows the Hexagonal Architecture (also known as Ports and Adapters) and SOLID principles:

### Hexagonal Architecture

The hexagonal architecture separates the application into layers:

1. **Domain Layer (Core)**: Contains the business logic and domain entities.
   - `internal/core/domain`: Domain entities like Quote
   - `internal/core/ports`: Interfaces that define how the core interacts with the outside world

2. **Application Layer**: Implements use cases using the domain.
   - `internal/app/services`: Services that implement the business logic

3. **Adapters Layer**: Connects the application to external systems.
   - **Primary/Driving Adapters**: Entry points to the application (CLI)
   - **Secondary/Driven Adapters**: Implementations of ports that connect to external systems (file system, configuration)

### SOLID Principles

- **Single Responsibility Principle**: Each class has only one reason to change.
  - The `QuoteRepository` is only responsible for retrieving quotes.
  - The `ConfigProvider` is only responsible for configuration.

- **Open/Closed Principle**: Software entities should be open for extension but closed for modification.
  - New repository implementations can be added without changing existing code.

- **Liskov Substitution Principle**: Objects should be replaceable with instances of their subtypes.
  - Any implementation of `QuoteRepository` can be used interchangeably.

- **Interface Segregation Principle**: Clients should not depend on interfaces they don't use.
  - Interfaces are small and focused on specific functionality.

- **Dependency Inversion Principle**: High-level modules should not depend on low-level modules.
  - The core domain depends on abstractions (ports), not concrete implementations.

## Project Structure

```
get-quote/
├── bin/                    # Compilation output directory
├── cmd/                    # Executable commands
│   └── get-quote/          # Main command
├── internal/               # Internal packages
│   ├── core/               # Domain layer (hexagon center)
│   │   ├── domain/         # Domain entities
│   │   └── ports/          # Ports (interfaces)
│   ├── adapters/           # Adapters layer
│   │   ├── primary/        # Primary/Driving adapters
│   │   │   └── cli/        # CLI command handler
│   │   └── secondary/      # Secondary/Driven adapters
│   │       ├── config/     # Configuration adapter
│   │       └── repository/ # Repository adapter
│   └── app/                # Application layer
│       └── services/       # Use cases implementation
├── src/                    # Source files
│   └── files/              # Data files
│       ├── quotes.lst      # Quotes in English
│       └── citas.lst       # Quotes in Spanish
├── go.mod                  # Go module definition
├── go.sum                  # Dependency checksums
├── Makefile                # Build automation
├── get-quote.yaml          # Configuration file
└── README.md               # Project documentation
```

## Development

### Building from Source

```bash
# Clean the build directory
make clean

# Build the application
make build

# Run examples
make run
```

### Deploy to Your bin Directory

```bash
# Deploy to ~/bin/get-quote
make deploy
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Contributing

Contributions are welcome! Feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request