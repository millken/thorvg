# Quick Start Guide

This guide will help you get started with ThorVG Go bindings in just a few minutes.

## Prerequisites

- Go 1.21 or later
- C compiler (gcc or clang)
- pkg-config
- Meson and Ninja build tools

## Step 1: Install ThorVG C Library

### On Ubuntu/Debian

```bash
# Install build dependencies
sudo apt-get update
sudo apt-get install -y meson ninja-build pkg-config gcc

# Clone and build ThorVG
git clone https://github.com/thorvg/thorvg.git
cd thorvg
meson build
sudo ninja -C build install
sudo ldconfig
```

### On macOS

```bash
# Install dependencies via Homebrew
brew install meson ninja pkg-config

# Clone and build ThorVG
git clone https://github.com/thorvg/thorvg.git
cd thorvg
meson build
sudo ninja -C build install
```

### Verify Installation

```bash
pkg-config --modversion thorvg
```

This should output the ThorVG version (e.g., 0.11.0).

## Step 2: Create Your First Go Project

```bash
# Create a new directory for your project
mkdir my-thorvg-project
cd my-thorvg-project

# Initialize a Go module
go mod init example.com/my-thorvg-project

# Get ThorVG Go bindings
go get github.com/millken/thorvg
```

## Step 3: Write Your First Program

Create a file named `main.go`:

```go
package main

import (
    "fmt"
    "os"
    "github.com/millken/thorvg"
)

func main() {
    // Initialize ThorVG
    init, err := thorvg.Init(thorvg.EngineTypeSW, 4)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }
    defer init.Term()

    // Create a canvas
    canvas, err := thorvg.NewCanvas(thorvg.EngineTypeSW)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }
    defer canvas.Destroy()

    // Set up rendering buffer
    width, height := uint(400), uint(300)
    buffer := make([]byte, width*height*4)
    
    canvas.SetTarget(buffer, width, width, height)
    canvas.Clear()

    // Create a red circle
    circle, _ := thorvg.NewShape()
    circle.AppendCircle(200, 150, 100, 100)
    circle.SetFillColor(255, 0, 0, 255)
    
    canvas.Push(&circle.Paint)

    // Render
    canvas.Draw()
    canvas.Sync()

    fmt.Println("Successfully rendered a red circle!")
}
```

## Step 4: Build and Run

```bash
go build -o my-app
./my-app
```

You should see: `Successfully rendered a red circle!`

## Next Steps

1. **Explore Examples**: Check out the [examples](./examples) directory for more complex use cases
2. **Read the Docs**: See the [API documentation](https://godoc.org/github.com/millken/thorvg)
3. **Save Output**: Modify the example to save the buffer as a PNG file (see `examples/simple.go`)
4. **Load SVG**: Try loading and rendering SVG files (see `examples/svg_loader.go`)

## Common Issues

### "Package thorvg was not found"

This means pkg-config cannot find ThorVG. Make sure:
- ThorVG is installed (`sudo ninja -C build install`)
- You ran `sudo ldconfig` (Linux) after installation
- pkg-config can find it: `pkg-config --libs thorvg`

### CGo Errors

Make sure you have a C compiler installed:
- Linux: `sudo apt-get install gcc`
- macOS: `xcode-select --install`

### Import Errors

Run `go mod tidy` to ensure all dependencies are properly resolved.

## Getting Help

- [ThorVG Documentation](https://github.com/thorvg/thorvg/wiki)
- [Open an Issue](https://github.com/millken/thorvg/issues)
- [ThorVG Official Site](https://www.thorvg.org/)

Happy coding!
