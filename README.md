# ThorVG Go Bindings

Go bindings for [ThorVG](https://github.com/thorvg/thorvg), a platform-independent portable library for drawing vector-based scenes and animations.

## Features

- 🎨 Vector graphics rendering
- 📐 Shape drawing (rectangles, circles, paths)
- 🖼️ SVG and image loading support
- 🎬 Animation support
- 🚀 Hardware-accelerated rendering
- 💾 Memory-efficient canvas operations

## Installation

### Prerequisites

First, install the ThorVG C library:

#### Ubuntu/Debian
```bash
# Install dependencies
sudo apt-get install meson ninja-build pkg-config

# Build and install ThorVG
git clone https://github.com/thorvg/thorvg.git
cd thorvg
meson build
ninja -C build install
```

#### macOS
```bash
brew install meson ninja pkg-config
git clone https://github.com/thorvg/thorvg.git
cd thorvg
meson build
ninja -C build install
```

### Install Go Bindings

```bash
go get github.com/millken/thorvg
```

## Usage

### Basic Example

```go
package main

import (
    "fmt"
    "github.com/millken/thorvg"
)

func main() {
    // Initialize ThorVG
    init, err := thorvg.Init(thorvg.EngineTypeSW, 4)
    if err != nil {
        panic(err)
    }
    defer init.Term()

    // Create a canvas
    canvas, err := thorvg.NewCanvas(thorvg.EngineTypeSW)
    if err != nil {
        panic(err)
    }
    defer canvas.Destroy()

    // Set up rendering target
    width, height := uint(800), uint(600)
    buffer := make([]byte, width*height*4)
    canvas.SetTarget(buffer, width, width, height)
    canvas.Clear()

    // Create and draw a red rectangle
    rect, _ := thorvg.NewShape()
    rect.AppendRect(100, 100, 200, 150)
    rect.SetFillColor(255, 0, 0, 255)
    canvas.Push(&rect.Paint)

    // Render
    canvas.Draw()
    canvas.Sync()
    
    fmt.Println("Rendered successfully!")
}
```

### Drawing Shapes

```go
// Rectangle
shape, _ := thorvg.NewShape()
shape.AppendRect(x, y, width, height)
shape.SetFillColor(255, 0, 0, 255) // Red

// Circle
circle, _ := thorvg.NewShape()
circle.AppendCircle(cx, cy, rx, ry)
circle.SetFillColor(0, 0, 255, 255) // Blue

// With stroke
shape.SetStrokeColor(0, 0, 0, 255)
shape.SetStrokeWidth(2.0)
```

### Loading Images

```go
picture, _ := thorvg.NewPicture()

// Load from file
err := picture.Load("image.svg")

// Or load from memory
data := []byte{...} // Your image data
err := picture.LoadData(data, "svg")

// Set size
picture.SetSize(200, 200)

// Add to canvas
canvas.Push(&picture.Paint)
```

## API Documentation

### Core Types

- **Canvas**: The main drawing surface
- **Shape**: Vector shapes (rectangles, circles, paths)
- **Picture**: Images and SVG files
- **Paint**: Base type for drawable objects

### Engine Types

- `EngineTypeSW`: Software rendering (CPU-based)
- `EngineTypeGL`: OpenGL rendering (GPU-accelerated)

### Color Spaces

- `ColorSpaceARGB8888`: ARGB 32-bit color
- `ColorSpaceABGR8888`: ABGR 32-bit color

## Examples

See the [examples](./examples) directory for complete working examples:

- `simple.go`: Basic shapes and rendering
- More examples coming soon!

## Building

```bash
# Build the library
go build

# Run examples (requires ThorVG C library installed)
go run examples/simple.go
```

## Requirements

- Go 1.16 or later
- ThorVG C library (0.11.0 or later)
- CGo enabled
- pkg-config

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [ThorVG](https://github.com/thorvg/thorvg) - The underlying vector graphics library
- The ThorVG team for creating an excellent C API

## Links

- [ThorVG Official Site](https://www.thorvg.org/)
- [ThorVG GitHub](https://github.com/thorvg/thorvg)
- [ThorVG Documentation](https://github.com/thorvg/thorvg/wiki)