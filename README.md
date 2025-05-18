# Pokedex CLI

A command-line interface Pokedex application that allows users to explore the world of Pokemon by accessing data from the [PokeAPI](https://pokeapi.co/).

## Features

- Interactive command-line interface with a REPL (Read-Eval-Print Loop)
- Browse through different Pokemon locations
- Paginated navigation through location areas
- Simple and intuitive command system

## Installation

### Prerequisites

- Go 1.18 or higher

### Steps

1. Clone the repository:
   ```
   git clone https://github.com/mutluerengazi/pokedex.git
   ```

2. Navigate to the project directory:
   ```
   cd pokedex
   ```

3. Build the application:
   ```
   go build
   ```

4. Run the application:
   ```
   ./pokedex
   ```

## Usage

Once you start the application, you'll be presented with a prompt:

```
Pokedex >
```

### Available Commands

| Command | Description |
|---------|-------------|
| `help`  | Displays a help message with all available commands |
| `map`   | Get the next page of Pokemon locations |
| `mapb`  | Get the previous page of Pokemon locations |
| `exit`  | Exit the Pokedex application |

### Examples

View available commands:
```
Pokedex > help

Welcome to the Pokedex!
Usage:

help: Displays a help message
map: Get the next page of locations
mapb: Get the previous page of locations
exit: Exit the Pokedex
```

Browse locations:
```
Pokedex > map
canalave-city
eterna-city
pastoria-city
sunyshore-city
sinnoh-pokemon-league
...
```

Navigate to previous locations:
```
Pokedex > mapb
```

Exit the application:
```
Pokedex > exit
Closing the Pokedex... Goodbye!
```

## Project Structure

- `main.go` - Entry point for the application
- `repl.go` - REPL implementation for command-line interface
- `config.go` - Configuration for API URLs
- Command files:
  - `command_help.go` - Implementation of the help command
  - `command_map.go` - Implementation of map and mapb commands
  - `command_exit.go` - Implementation of the exit command
- `internal/pokeapi/` - Client implementation for interfacing with the PokeAPI

## Technologies

- Go programming language
- [PokeAPI](https://pokeapi.co/) - RESTful Pokemon data API

## Testing

Run the tests with:

```
go test ./...
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgements

- [PokeAPI](https://pokeapi.co/) for providing the Pokemon data