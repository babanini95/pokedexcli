# Pokedexcli

A command-line interface (CLI) Pokedex application written in Go.

## Description

This application allows users to interact with Pokemon data through a terminal-based interface. Users can explore Pokemon information, catch Pokemon, and manage their collection. All the data fetched from [Poke API](https://pokeapi.co/)

## Installation

```bash
git clone https://github.com/babanini95/pokedexcli.git
cd pokedexcli
go build
```

## Usage

Run the application:
```bash
./pokedexcli
```

### Available Commands
- `help`    : Displays a list of available commands
- `exit`    : Exits the program
- `map`     : Displays the next 20 location areas in Pokemon world
- `mapb`    : Displays the previous 20 location areas in Pokemon world
- `explore` : Lists all the pokemon located in a given area. Accept area's name or ID as argument
- `catch`   : Try catching a pokemon. Accept pokemon's name or ID as argument
- `inspect` : Shows Pokemon's name, height, weight, stats and type(s). Accept pokemon's name as argument
- `pokedex` : Lists all Pokemon you have been caught

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)