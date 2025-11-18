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
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <svg-file>\n", os.Args[0])
		os.Exit(1)
	}

	svgPath := os.Args[1]

	// Initialize ThorVG
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

	// Load SVG
	picture, err := thorvg.NewPicture()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create picture: %v\n", err)
		os.Exit(1)
	}

	if err := picture.Load(svgPath); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load SVG: %v\n", err)
		os.Exit(1)
	}

	// Get original size
	origW, origH, err := picture.GetSize()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get size: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Original SVG size: %.0fx%.0f\n", origW, origH)

	// Set desired output size
	width, height := uint(800), uint(600)
	if err := picture.SetSize(float32(width), float32(height)); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to set size: %v\n", err)
		os.Exit(1)
	}

	// Set up the target buffer
	buffer := make([]byte, width*height*4)

	if err := canvas.SetTarget(buffer, width, width, height); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to set target: %v\n", err)
		os.Exit(1)
	}

	// Clear the canvas with white background
	if err := canvas.Clear(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to clear canvas: %v\n", err)
		os.Exit(1)
	}

	// Add picture to canvas
	if err := canvas.Push(&picture.Paint); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to push picture: %v\n", err)
		os.Exit(1)
	}

	// Render
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

	outputPath := "svg_output.png"
	f, err := os.Create(outputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create output file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to encode PNG: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully rendered SVG to %s\n", outputPath)
}
