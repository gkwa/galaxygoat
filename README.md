# GalaxyGoat

A utility to remove specific HTML elements (like `path` elements) from HTML content while preserving the overall structure.

## Features

- Uses Cobra CLI for a modern command-line interface
- Follows SOLID principles with clean separation of concerns
- Handles HTML properly using the Go standard HTML parser
- Designed for use in Unix pipelines

## Usage

```bash
# Remove all path elements (default)
cat input.html | galaxygoat > output.html

# Remove multiple element types
cat input.html | galaxygoat --remove="path,svg,script" > output.html
# or
cat input.html | galaxygoat -r "path,svg,script" > output.html
```

