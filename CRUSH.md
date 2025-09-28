# CRUSH.md

## Build


## Lint


## Test


## Code Style Guidelines

### Imports

Grouped by standard library, then third-party, then project packages.

### Formatting

Use `go fmt` to format code.

### Types

Prefer concrete types over interfaces where appropriate.

### Naming Conventions

Use camelCase for variables and functions.  Use PascalCase for structs and interfaces.

### Error Handling

Handle errors explicitly.
Avoid using `panic` unless for unrecoverable errors during initialization.
