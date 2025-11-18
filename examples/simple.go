// +build ignore

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/millken/thorvg"
)

func main() {
	// Initialize ThorVG with software engine
	init, err := thorvg.Init(thorvg.EngineTypeSW, 4)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize ThorVG: %v\n", err)
		os.Exit(1)
	}
	defer init.Term()

	// Create a canvas
	canvas, err := thorvg.NewCanvas(thorvg.EngineTypeSW)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create canvas: %v\n", err)
		os.Exit(1)
	}
	defer canvas.Destroy()

	// Set up the target buffer
	width, height := uint(800), uint(600)
	buffer := make([]byte, width*height*4)

	if err := canvas.SetTarget(buffer, width, width, height); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to set target: %v\n", err)
		os.Exit(1)
	}

	// Clear the canvas
	if err := canvas.Clear(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to clear canvas: %v\n", err)
		os.Exit(1)
	}

	// Create a red rectangle
	rect, err := thorvg.NewShape()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create shape: %v\n", err)
		os.Exit(1)
	}

	if err := rect.AppendRect(100, 100, 200, 150); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to append rect: %v\n", err)
		os.Exit(1)
	}

	if err := rect.SetFillColor(255, 0, 0, 255); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to set fill color: %v\n", err)
		os.Exit(1)
	}

	if err := canvas.Push(&rect.Paint); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to push shape: %v\n", err)
		os.Exit(1)
	}

	// Create a blue circle
	circle, err := thorvg.NewShape()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create circle: %v\n", err)
		os.Exit(1)
	}

	if err := circle.AppendCircle(500, 300, 100, 100); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to append circle: %v\n", err)
		os.Exit(1)
	}

	if err := circle.SetFillColor(0, 0, 255, 255); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to set fill color: %v\n", err)
		os.Exit(1)
	}

	if err := canvas.Push(&circle.Paint); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to push circle: %v\n", err)
		os.Exit(1)
	}

	// Draw the canvas
	if err := canvas.Draw(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to draw: %v\n", err)
		os.Exit(1)
	}

	if err := canvas.Sync(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to sync: %v\n", err)
		os.Exit(1)
	}

	// Convert buffer to image and save as PNG
	img := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	for y := 0; y < int(height); y++ {
		for x := 0; x < int(width); x++ {
			idx := (y*int(width) + x) * 4
			img.Set(x, y, color.RGBA{
				R: buffer[idx+2],
				G: buffer[idx+1],
				B: buffer[idx+0],
				A: buffer[idx+3],
			})
		}
	}

	f, err := os.Create("output.png")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create output file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to encode PNG: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Successfully rendered to output.png")
}
