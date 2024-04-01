A CLI tool written in Go which pulls information from PokeAPI and displays to the user. 

This is my first project using Go and is part of the [Boot.dev](https://www.boot.dev/assignments/dff17f87-1ce8-43ce-a43b-2cb611ce76f1) curriculum. The initial commit to this repository was created by following [wagslane's wonderful guide](https://www.youtube.com/watch?v=8yrmAGcCnKg).

| Command | Description |
|---------|-------------|
|`help`| Displays the help menu|
|`exit`| Exits the program|
|`map` | Displays the next 20 location areas |
|`mapb`| Displays the previous 20 location areas |
|`explore <location-area>`| Lists Pokemon available in a given location area|
|`catch <pokemon>`| Attempts to catch a Pokemon |
|`inspect <pokemon>`| Displays stats for a caught Pokemon|
|`pokedex`| Lists all caught Pokemon|

# Entensibility 
Here are features I may go back to implement: 
- Locations (instead of location-areas) with sub-areas
- Restrict catchable Pokemon to explored location-area
- Evolve caught Pokemon
- Settings for Pokedex to include seen vs caught

Backend improvements:
- Create submenu map navigation instead of `map` and `mapb`
- Cache full list of locations rather than 20 at a time
- Output a save file

Organization:
- Put commands into their own folder

Stretch goal: 
- Create ASCII map tiles to create a simple 2D experience
