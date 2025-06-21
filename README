# Pokedex

A command-line Pokedex application written in Go. This project allows users to explore Pokémon locations, catch Pokémon, and manage their own Pokedex using data from the PokeAPI.

## Features

- List and navigate Pokémon location areas
- Explore specific location areas to discover wild Pokémon
- Catch Pokémon and add them to your personal Pokedex
- Inspect details about caught Pokémon
- View your collection of caught Pokémon
- Helpful command-line interface with built-in help

## Getting Started

### Prerequisites

- Go 1.18 or newer

### Installation

Clone the repository:

```sh
git clone <repository-url>
cd <repository-directory>
```

Build the project:

```sh
go build
```

### Usage

Run the application:

```sh
./pokedex
```

You will see a prompt.  
Type commands to interact with the Pokedex.

## Commands

| Command   | Description                          | Usage                       |
|-----------|--------------------------------------|-----------------------------|
| help      | Display help message                 | `help`                      |
| exit      | Exit the Pokedex                     | `exit`                      |
| map       | Show next page of location areas     | `map`                       |
| mapb      | Show previous page of location areas | `mapb`                      |
| explore   | Explore a specific location area     | `explore <location-area>`   |
| catch     | Attempt to catch a Pokémon           | `catch <pokemon-name>`      |
| inspect   | Show details about a caught Pokémon  | `inspect <pokemon-name>`    |
| pokedex   | List all caught Pokémon              | `pokedex`                   |

## Example

```
> map
Showing locations 1-20
> explore viridian-forest
Wild Pokémon: caterpie, weedle, pikachu
> catch pikachu
You caught pikachu!
> pokedex
Your Pokémon: pikachu
```

## Project Structure

- `main.go`: Entry point and REPL loop
- `commands.go`: Command implementations
- `repl.go`: Input cleaning and CLI types
- `pokeapi`: PokeAPI client and data models
- `pokecache`: Caching utilities

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

_Inspired by Pokémon and powered by [PokeAPI](https://pokeapi.co/). Not affiliated with Nintendo or The Pokémon Company._