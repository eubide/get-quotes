# Get Quote

A simple Go tool for extracting random lines from text files.

## Project Structure

```
get-quote/
├── bin/                    # Compilation output directory
├── cmd/                    # Executable commands
│   └── get-quote/          # Main command
├── pkg/                    # Package code
│   ├── config/             # Configuration handling
│   └── randomline/         # Random line functionality
├── src/                    # Source files
│   └── files/              # Data files
│       ├── quotes.lst      # Quotes in English
│       └── citas.lst       # Quotes in Spanish
├── go.mod                  # Go module definition
├── go.sum                  # Dependency checksums
├── Makefile                # Build automation
├── .get-quote.yaml         # Configuration file
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

## Configuration

The application uses a YAML configuration file called `.get-quote.yaml`. The configuration file can be placed in:
1. The current directory
2. `$HOME/.config/.get-quote/`
3. The user's home directory

If no configuration file is found, default values are used.

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