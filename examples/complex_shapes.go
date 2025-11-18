// +build ignore

// This example demonstrates creating multiple shapes with different colors and styles
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
	// Initialize ThorVG with software engine and 4 threads
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
	width, height := uint(1000), uint(800)
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

	// Create a background rectangle (light gray)
	bg, _ := thorvg.NewShape()
	bg.AppendRect(0, 0, float32(width), float32(height))
	bg.SetFillColor(240, 240, 240, 255)
	canvas.Push(&bg.Paint)

	// Create multiple colored rectangles
	colors := []struct {
		r, g, b uint8
		name    string
	}{
		{255, 100, 100, "Red"},
		{100, 255, 100, "Green"},
		{100, 100, 255, "Blue"},
		{255, 255, 100, "Yellow"},
		{255, 100, 255, "Magenta"},
		{100, 255, 255, "Cyan"},
	}

	x := float32(50)
	for i, col := range colors {
		rect, _ := thorvg.NewShape()
		y := float32(50 + i*120)
		rect.AppendRect(x, y, 200, 100)
		rect.SetFillColor(col.r, col.g, col.b, 255)
		
		// Add a black stroke
		rect.SetStrokeColor(0, 0, 0, 255)
		rect.SetStrokeWidth(3.0)
		
		canvas.Push(&rect.Paint)
	}

	// Create circles with gradual transparency
	for i := 0; i < 5; i++ {
		circle, _ := thorvg.NewShape()
		cx := float32(400 + i*100)
		cy := float32(300)
		radius := float32(60)
		
		circle.AppendCircle(cx, cy, radius, radius)
		
		// Gradual transparency
		alpha := uint8(255 - i*40)
		circle.SetFillColor(200, 100, 50, alpha)
		circle.SetStrokeColor(50, 50, 50, 255)
		circle.SetStrokeWidth(2.0)
		
		canvas.Push(&circle.Paint)
	}

	// Create overlapping shapes demonstrating transparency
	for i := 0; i < 4; i++ {
		shape, _ := thorvg.NewShape()
		x := float32(350 + i*60)
		y := float32(550)
		
		shape.AppendRect(x, y, 120, 120)
		
		// Semi-transparent colors
		colors := [][]uint8{
			{255, 0, 0, 180},
			{0, 255, 0, 180},
			{0, 0, 255, 180},
			{255, 255, 0, 180},
		}
		
		col := colors[i%len(colors)]
		shape.SetFillColor(col[0], col[1], col[2], col[3])
		
		canvas.Push(&shape.Paint)
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

	outputPath := "complex_shapes.png"
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

	fmt.Printf("Successfully rendered complex shapes to %s\n", outputPath)
	fmt.Println("The output demonstrates:")
	fmt.Println("- Multiple colored rectangles with strokes")
	fmt.Println("- Circles with gradual transparency")
	fmt.Println("- Overlapping semi-transparent shapes")
}
