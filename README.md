# Get Quote

A simple Go tool for extracting random lines from text files.

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

The project follows Hexagonal Architecture (Ports and Adapters) and SOLID principles:

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

## Installation

```bash
# Clone the repository
git clone https://github.com/eubide/get-quote.git
cd get-quote

# Build the executable
make build
```

## Usage

```bash
./bin/get-quote filename
```

Where `filename` is the name of a file in the configured directory (default `src/files/`).
The `.lst` extension will be added automatically if not provided.

## Examples

Extract a random line from the quotes.lst file:
```bash
./bin/get-quote quotes
```

Extract a random line from the citas.lst file:
```bash
./bin/get-quote citas
```

Use a specific configuration file:
```bash
./bin/get-quote -c ~/my-custom-config.yaml quotes
```

## Configuration

The application uses a YAML configuration file. The configuration file is searched in the following order of preference:

1. Path specified with the `-c` flag (e.g., `./bin/get-quote -c /path/to/config.yaml quotes`)
2. In the user's config directory: `$HOME/.config/get-quote/get-quote.yaml`
3. In the user's home directory: `$HOME/.get-quote.yaml`
4. In the current working directory: `get-quote.yaml`

If no configuration file is found in any of these locations, default values are used.

### Command-line Arguments

```bash
./bin/get-quote [options] <filename>
```

Options:
- `-c <file>`: Specify a custom configuration file path

### Configuration Example

```yaml
# Random Sentence Configuration

filesBaseDir: src/files
defaultExtension: .lst
errorMessages:
  fileNotFound: "The file %s does not exist"
  fileOpenError: "Error opening the file: %v"
  missingParameter: "Usage: %s <filename>\nYou must provide a filename %s"
```

The `filesBaseDir` setting can use the tilde character (`~`) to represent the user's home directory. For example:

```yaml
# Will use the user's home directory directly
filesBaseDir: ~

# Will use $HOME/quotes directory
filesBaseDir: ~/quotes  

# Will use $HOME/.config/get-quote/quotes directory
filesBaseDir: ~/.config/get-quote/quotes
```

Note that only the exact tilde (`~`) or a tilde followed by a slash (`~/`) will be expanded. Paths like `~user` will be left as-is.
