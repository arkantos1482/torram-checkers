# Checkers Game on Cosmos SDK

This project implements a Checkers game on the Cosmos SDK, allowing players to create games, query their status, and export the game state. Users can create games with basic settings or use the *Torram* option for more detailed game creation.

## Commands

### Create a Checkers Game with Detailed Settings (Torram)
Creates a Checkers game with additional parameters, including start and end times and a description, along with the black and red players.

```bash
minid tx checkers torram create-game [flags]
```

### Create a Basic Checkers Game
Creates a Checkers game at a specified index for the black and red players.

```bash
minid tx checkers normal create-game [flags]
```
Query Game by Index
Retrieve the current state of a game at a given index.

```bash
minid query checkers get-game [index] [flags]
```
Export State to JSON
Export the current state of the chain to a JSON file.

```bash
minid export [flags]
```

### Getting Started
Install Dependencies: Make sure to have the Cosmos SDK installed and configured.
Build the Project: Run make to build the project.
Start a Local Node: Use minid start to start the local blockchain node.
Run Commands: Use the commands listed above to create and query games.
Development
Feel free to explore and modify the code to enhance the Checkers game or add new features!

### License
This project is licensed under the MIT License.
