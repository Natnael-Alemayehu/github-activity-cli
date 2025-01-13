# GitHub Activity CLI

A command-line interface tool that displays a user's recent GitHub activities, including push events, event types, and repository interactions.

## Description

This CLI tool fetches and displays GitHub user activity data, providing information about:
- Recent push events and commit counts
- Types of GitHub events performed
- Repositories the user has interacted with

## Installation

### Prerequisites
- Go 1.23.1 or higher
- GitHub account

### Building from Source
```bash
git clone https://github.com/natnael-alemayehu/github-activity-cli.git
cd github-activity-cli
go build ./cmd/github-activity
```

## Usage

Run the CLI tool by providing a GitHub username as an argument:
```sh
./github-activity <username>
```

## Features

- Fetches real-time GitHub activity data using the GitHub API
- Displays push events with commit counts
- Shows a summary of different event types
- Lists all repositories the user has interacted with
- Simple and intuitive command-line interface

## Project Structure

- `cmd/github-activity/`: Contains the main application entry point
- `internal/`: Internal package code
  - `data/`: Data structures for GitHub API responses
  - `consumeBody.go`: GitHub API interaction logic
  - `readFile.go`: File handling utilities

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit

Project idea https://roadmap.sh/projects/github-user-activity