/*
Package thorvg provides Go bindings for the ThorVG vector graphics library.

ThorVG is a platform-independent portable library for drawing vector-based
scenes and animations. It provides high-performance rendering with support
for SVG, shapes, and various graphics operations.

# Installation

Before using this package, you need to install the ThorVG C library:

	# Ubuntu/Debian
	sudo apt-get install meson ninja-build pkg-config
	git clone https://github.com/thorvg/thorvg.git
	cd thorvg
	meson build
	ninja -C build install

Then install the Go bindings:

	go get github.com/millken/thorvg

# Basic Usage

Initialize ThorVG, create a canvas, and render shapes:

	package main

	import (
		"github.com/millken/thorvg"
	)

	func main() {
		// Initialize ThorVG
		init, err := thorvg.Init(thorvg.EngineTypeSW, 4)
		if err != nil {
			panic(err)
		}
		defer init.Term()

		// Create canvas
		canvas, err := thorvg.NewCanvas(thorvg.EngineTypeSW)
		if err != nil {
			panic(err)
		}
		defer canvas.Destroy()

		// Set up buffer
		width, height := uint(800), uint(600)
		buffer := make([]byte, width*height*4)
		canvas.SetTarget(buffer, width, width, height)
		canvas.Clear()

		// Create a shape
		rect, _ := thorvg.NewShape()
		rect.AppendRect(100, 100, 200, 150)
		rect.SetFillColor(255, 0, 0, 255)
		canvas.Push(&rect.Paint)

		// Render
		canvas.Draw()
		canvas.Sync()
	}

# Drawing Shapes

ThorVG supports various shapes:

	// Rectangle
	rect, _ := thorvg.NewShape()
	rect.AppendRect(x, y, width, height)
	rect.SetFillColor(r, g, b, a)

	// Circle
	circle, _ := thorvg.NewShape()
	circle.AppendCircle(cx, cy, rx, ry)
	circle.SetFillColor(r, g, b, a)

	// With stroke
	rect.SetStrokeColor(0, 0, 0, 255)
	rect.SetStrokeWidth(2.0)

# Loading Images and SVG

Load vector graphics and images:

	picture, _ := thorvg.NewPicture()

	// From file
	picture.Load("image.svg")

	// From memory
	picture.LoadData(data, "svg")

	// Set size
	picture.SetSize(200, 200)

	// Add to canvas
	canvas.Push(&picture.Paint)

# Engine Types

ThorVG supports multiple rendering engines:

  - EngineTypeSW: Software rendering (CPU-based)
  - EngineTypeGL: OpenGL rendering (GPU-accelerated)

# Error Handling

All operations that can fail return an error. Always check errors:

	if err := canvas.Draw(); err != nil {
		log.Fatalf("Failed to draw: %v", err)
	}

# Thread Safety

ThorVG is not thread-safe. Ensure that all ThorVG operations are
performed from the same goroutine, or use proper synchronization.

# Memory Management

The Canvas and Paint objects should be properly cleaned up:

	canvas, _ := thorvg.NewCanvas(thorvg.EngineTypeSW)
	defer canvas.Destroy()

The Initializer should also be terminated:

	init, _ := thorvg.Init(thorvg.EngineTypeSW, 4)
	defer init.Term()
*/
package thorvg
