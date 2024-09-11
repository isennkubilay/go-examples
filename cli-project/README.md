# DevOps Daily CLI

DevOps Daily CLI is a command-line interface tool designed to simplify common Docker and Kubernetes operations for DevOps engineers and developers.

## Features

### Docker Commands
- Docker Login
- View Docker Logs
- Build Docker Images
- Push Docker Images
- Pull Docker Images
- Save Docker Images (gzip)
- Load Docker Images

### Kubernetes Commands
- Get Pods
- Get Logs from Pod
- Delete Pod
- Get Services (placeholder)
- Apply Manifest (placeholder)

## Installation

1. Ensure you have Go installed on your system. If not, download and install it from [golang.org](https://golang.org/).

2. Clone this repository:
   ```
   git clone https://github.com/isennkubilay/cli-project.git
   ```

3. Navigate to the project directory:
   ```
   cd cli-project
   ```

4. Build the project:
   ```
   go build
   ```

## Usage

Run the compiled binary: `./devops-cli`


Follow the on-screen prompts to navigate through the available commands and options.

## Requirements

- Go 1.16 or later
- Docker (for Docker commands)
- kubectl configured with appropriate permissions (for Kubernetes commands)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
