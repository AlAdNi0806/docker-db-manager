# Docker Database Management CLI

A powerful, interactive command-line tool for managing Docker-based database containers. Built with Go, this CLI simplifies the creation, removal, and management of MySQL and MariaDB databases running in Docker containers.

## Features

- **Interactive Interface**: User-friendly terminal-based forms for easy database management
- **Database Support**: Full support for MySQL and MariaDB
- **Container Management**: Create, remove, and list database containers, images, and volumes
- **Flexible Configuration**: Customize root passwords, database names, and image versions
- **Docker Integration**: Seamlessly works with Docker to handle container lifecycle

## Requirements

- Go 1.25.1 or later
- Docker installed and running
- Linux/macOS/Windows (with appropriate shell support)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/docker-db-management.git
   cd docker-db-management
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Build and install globally:
   ```bash
   make install
   ```

   This will install the `manage-db` binary to `/usr/local/bin`, allowing you to run it from anywhere.

## Usage

After installation, simply run:
```bash
manage-db
```

The CLI will guide you through an interactive process:

1. **Select Action**: Choose from:
   - Create database
   - Remove database
   - Manage db containers/image/volumes

2. **Select Database Type**: Choose between MySQL or MariaDB

3. **Configure Options** (for creation):
   - Pull latest Docker image (recommended for updates)
   - Set root password (defaults to "12345678" if left blank)
   - Specify database name (optional)

## Build from Source

To build without installing globally:
```bash
make build
```

This creates a `manage-db` executable in the current directory and runs it immediately.

## Uninstall

To remove the globally installed binary:
```bash
make uninstall
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Disclaimer

This tool is designed for development and testing environments. Use appropriate security measures when deploying databases in production.