# Pokedex CLI

## Overview

The Pokedex CLI is a command-line application developed in Go, leveraging the PokéAPI to provide users with information about Pokémon. This project is aimed at simulating the functionality of the Pokedex device from the Pokémon series, allowing users to explore locations, discover Pokémon, and manage a collection of caught Pokémon, all from the convenience of their command line.

## Features

- **Interactive REPL**: An interactive read-eval-print loop where users can input commands and receive responses.
- **PokéAPI Integration**: Fetches real-time data about Pokémon, including names, types, and stats.
- **Caching**: Implements caching to store API responses, significantly improving response times for previously fetched data.
- **Pokémon Collection Management**: Users can catch Pokémon and store them in their personal Pokedex, viewing details about each at any time.

## Getting Started

### Prerequisites

- Go programming language installed on your system.
- Basic familiarity with command line interfaces and Go development environment.

### Installation

1. Clone the repository to your local machine.
2. Navigate to the project directory.
3. Compile the project using Go:

    ```bash
    go build
    ```

4. Run the compiled application:

    ```bash
    ./pokedexcli
    ```

## Usage

Upon running the application, you'll enter an interactive REPL mode. Here are the available commands:

- `help`: Displays information on how to use the CLI.
- `exit`: Exits the application.
- `map`: Lists 20 location areas within the Pokémon world. Successive calls cycle through additional locations.
- `mapb`: Lists the previous 20 location areas, allowing for navigation backwards.
- `explore <area_name>`: Shows all Pokémon available in a specified location area.
- `catch <pokemon_name>`: Attempts to catch a Pokémon and adds it to your Pokedex if successful.
- `inspect <pokemon_name>`: Displays detailed information about a caught Pokémon.
- `pokedex`: Lists all Pokémon that have been caught and stored in your Pokedex.

### Example

```plaintext
./pokedexcli
Pokedex > help

Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
...

Pokedex > exit
```

## Development

During the development of the Pokedex CLI, we focused on creating an intuitive and efficient user experience while ensuring the application could handle real-time data from the PokéAPI robustly. Here are some key aspects and learnings from the development process:

### Parsing JSON and Making HTTP Requests in Go

One of the core functionalities of this project involved fetching and parsing JSON data from the PokéAPI. This required a solid understanding of Go's `net/http` package for making HTTP requests and the `encoding/json` package for parsing the JSON response into Go structs. Handling dynamic JSON responses and efficiently mapping them to Go's static types was an interesting challenge that was overcome by carefully designing our data structures and utilizing interface types where necessary.

### Building an Interactive CLI

Creating a responsive and user-friendly command-line interface (CLI) was paramount. We achieved this by implementing a Read-Eval-Print Loop (REPL) that continuously listens for user input, processes the input as commands, and returns the output. This involved structuring commands in a way that they're easily extendable and maintaining a clean separation of concerns between the user interface and the application logic.

### Caching for Performance Optimization

To minimize the latency experienced due to network requests to the PokéAPI and to reduce the load on their servers, we implemented a caching mechanism. This cache stores recently fetched data in memory, making subsequent requests for the same data instantaneous. The challenge here was to design a thread-safe caching mechanism that could handle concurrent access patterns without compromising on performance. We used Go's concurrency primitives, like mutexes, to protect shared resources and ensure our caching system was both efficient and safe.

### Concurrency and Go Routines

Leveraging Go's powerful concurrency model, we were able to make non-blocking HTTP requests to the PokéAPI, enhancing the responsiveness of our CLI. Implementing concurrency required careful management of go routines and channels to avoid common pitfalls such as race conditions and deadlocks.

### Error Handling and User Feedback

Providing clear and helpful feedback for user errors (e.g., invalid commands or parameters) was essential for usability. We developed a comprehensive error handling system that not only catches and logs internal errors for debugging purposes but also guides users towards correct usage.

### Continuous Learning

This project was an excellent opportunity to deepen our understanding of Go and its ecosystem. From effective package management with Go Modules to exploring best practices for CLI design, each aspect of the project brought its learning curve. Moreover, writing idiomatic Go code and following community guidelines helped in making the codebase more maintainable and understandable for new contributors.

The development of the Pokedex CLI was a journey full of challenges, learning, and fun. It stands as a testament to the capabilities of Go as a language for building performant and reliable CLI tools.
