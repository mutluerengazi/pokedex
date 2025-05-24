# Pokédex CLI

A fully-featured command-line Pokédex that lets you **explore the Pokémon world, catch Pokémon, inspect their stats, and list everything you’ve caught** – all from your terminal.
Data comes live from the public [PokéAPI](https://pokeapi.co/), with an in-memory cache so repeated queries are instant.

---

## ✨ Feature Overview

| Category         | What you can do                                                                                 |
| ---------------- | ----------------------------------------------------------------------------------------------- |
| **World map**    | Step through *location-areas* page-by-page (`map`, `mapb`).                                     |
| **Exploration**  | `explore <area>` – list every wild Pokémon that appears in that area.                           |
| **Catching**     | `catch <pokemon>` – throw a Pokéball; difficulty scales with the Pokémon’s base-experience.     |
| **Inspection**   | `inspect <pokemon>` – see height, weight, full stat block and types (only after you catch it!). |
| **Your Pokédex** | `pokedex` – alphabetic list of everything you’ve caught so far.                                 |
| **Help / Quit**  | `help`, `exit`.                                                                                 |
| **Performance**  | Transparent 5-second cache layer for all API calls.                                             |
| **Tests**        | Unit tests for cache logic and more (`go test ./...`).                                          |

---

## 🚀 Installation

### Prerequisites

* Go 1.20 or newer (the program relies on the new default seeding behaviour of `math/rand`).

### Steps

```bash
git clone https://github.com/mutluerengazi/pokedex.git
cd pokedex
go build          # produces ./pokedex
./pokedex
```

---

## 🕹️ Usage

When you start the binary you’ll see the prompt:

```
Pokedex >
```

### Command reference

| Command   | Arguments         | Description                              |
| --------- | ----------------- | ---------------------------------------- |
| `help`    | –                 | Show every command.                      |
| `map`     | –                 | Next page of location-areas.             |
| `mapb`    | –                 | Previous page of location-areas.         |
| `explore` | `<location-area>` | List wild Pokémon in that area.          |
| `catch`   | `<pokemon>`       | Attempt to catch a Pokémon.              |
| `inspect` | `<pokemon>`       | Show stats of a Pokémon you have caught. |
| `pokedex` | –                 | List all caught Pokémon.                 |
| `exit`    | –                 | Quit the application.                    |

### Quick demo

```text
Pokedex > map
pastoria-city-area
route-212-south
great-marsh-area-1
…

Pokedex > explore pastoria-city-area
Exploring pastoria-city-area...
Found Pokémon:
 - tentacool
 - remoraid
 - shellder
 - …

Pokedex > catch shellder
Throwing a Pokéball at shellder...
shellder was caught!
You may now inspect it with the inspect command.

Pokedex > inspect shellder
Name: shellder
Height: 3
Weight: 40
Stats:
  -hp: 30
  -attack: 65
  -defense: 100
  -special-attack: 45
  -special-defense: 25
  -speed: 40
Types:
  - water

Pokedex > pokedex
Your Pokedex:
 - shellder
```

---

## 🗂 Project Structure

```
.
├── cmd/ (optional)        # main.go entry if you use a cmd folder
├── repl.go                # REPL loop & command dispatch
├── command_<x>.go         # one file per CLI command
├── internal/
│   ├── pokeapi/           # thin, cache-aware HTTP client
│   │   ├── client.go
│   │   ├── pokemon.go
│   │   ├── location.go
│   │   └── …
│   └── pokecache/         # thread-safe TTL cache (map+mutex+ticker)
└── README.md
```

---

## 🧪 Testing

```
go test ./...
```

The test suite covers cache insertion, lookup and automatic reaping.

---

## 🔧 Tech Stack

* **Go** – standard library only (HTTP, JSON, sync, time, math/rand).
* **PokéAPI** – community-maintained REST API for Pokémon data.

---

## 🙏 Acknowledgements

* Thanks to [PokéAPI](https://pokeapi.co/) for the fantastic public dataset.

