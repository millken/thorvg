// Package thorvg provides Go bindings for ThorVG vector graphics library.
// ThorVG is a platform-independent portable library for drawing vector-based scenes and animations.
package thorvg

/*
#cgo pkg-config: thorvg
#include <thorvg_capi.h>
#include <stdlib.h>
*/
import "C"
import (
	"errors"
	"unsafe"
)

// Result represents the result code from ThorVG operations
type Result int

const (
	ResultSuccess               Result = C.TVG_RESULT_SUCCESS
	ResultInvalidArgument       Result = C.TVG_RESULT_INVALID_ARGUMENT
	ResultInsufficientCondition Result = C.TVG_RESULT_INSUFFICIENT_CONDITION
	ResultFailedAllocation      Result = C.TVG_RESULT_FAILED_ALLOCATION
	ResultMemoryCorruption      Result = C.TVG_RESULT_MEMORY_CORRUPTION
	ResultNotSupported          Result = C.TVG_RESULT_NOT_SUPPORTED
	ResultUnknown               Result = C.TVG_RESULT_UNKNOWN
)

// EngineType represents the rendering engine type
type EngineType int

const (
	EngineTypeSW EngineType = C.TVG_ENGINE_SW
	EngineTypeGL EngineType = C.TVG_ENGINE_GL
)

// ColorSpace represents the color space
type ColorSpace int

const (
	ColorSpaceABGR8888 ColorSpace = C.TVG_COLORSPACE_ABGR8888
	ColorSpaceARGB8888 ColorSpace = C.TVG_COLORSPACE_ARGB8888
)

// Error converts a Result to a Go error
func (r Result) Error() error {
	if r == ResultSuccess {
		return nil
	}
	switch r {
	case ResultInvalidArgument:
		return errors.New("thorvg: invalid argument")
	case ResultInsufficientCondition:
		return errors.New("thorvg: insufficient condition")
	case ResultFailedAllocation:
		return errors.New("thorvg: failed allocation")
	case ResultMemoryCorruption:
		return errors.New("thorvg: memory corruption")
	case ResultNotSupported:
		return errors.New("thorvg: not supported")
	default:
		return errors.New("thorvg: unknown error")
	}
}

// Initializer manages ThorVG initialization and cleanup
type Initializer struct {
	engineType EngineType
	threads    uint
}

// Init initializes the ThorVG library
func Init(engineType EngineType, threads uint) (*Initializer, error) {
	result := C.tvg_engine_init(C.Tvg_Engine(engineType), C.uint(threads))
	if err := Result(result).Error(); err != nil {
		return nil, err
	}
	return &Initializer{
		engineType: engineType,
		threads:    threads,
	}, nil
}

// Term terminates the ThorVG library
func (i *Initializer) Term() error {
	result := C.tvg_engine_term(C.Tvg_Engine(i.engineType))
	return Result(result).Error()
}

// Canvas represents a ThorVG canvas
type Canvas struct {
	ptr *C.Tvg_Canvas
}

// NewCanvas creates a new canvas
func NewCanvas(engineType EngineType) (*Canvas, error) {
	ptr := C.tvg_swcanvas_create()
	if ptr == nil {
		return nil, errors.New("thorvg: failed to create canvas")
	}
	return &Canvas{ptr: ptr}, nil
}

// SetTarget sets the target buffer for the canvas
func (c *Canvas) SetTarget(buffer []byte, stride, w, h uint) error {
	if len(buffer) == 0 {
		return errors.New("thorvg: buffer is empty")
	}
	result := C.tvg_swcanvas_set_target(
		c.ptr,
		(*C.uint32_t)(unsafe.Pointer(&buffer[0])),
		C.uint(stride),
		C.uint(w),
		C.uint(h),
		C.Tvg_Colorspace(ColorSpaceARGB8888),
	)
	return Result(result).Error()
}

// Clear clears the canvas with full transparency
func (c *Canvas) Clear() error {
	result := C.tvg_canvas_clear(c.ptr, C.bool(true))
	return Result(result).Error()
}

// Update updates the canvas
func (c *Canvas) Update() error {
	result := C.tvg_canvas_update(c.ptr)
	return Result(result).Error()
}

// Draw draws the canvas
func (c *Canvas) Draw() error {
	result := C.tvg_canvas_draw(c.ptr)
	return Result(result).Error()
}

// Sync synchronizes the canvas
func (c *Canvas) Sync() error {
	result := C.tvg_canvas_sync(c.ptr)
	return Result(result).Error()
}

// Push adds a paint to the canvas
func (c *Canvas) Push(paint *Paint) error {
	result := C.tvg_canvas_push(c.ptr, paint.ptr)
	return Result(result).Error()
}

// Destroy destroys the canvas
func (c *Canvas) Destroy() error {
	if c.ptr != nil {
		result := C.tvg_canvas_destroy(c.ptr)
		c.ptr = nil
		return Result(result).Error()
	}
	return nil
}

// Paint represents a ThorVG paint object
type Paint struct {
	ptr *C.Tvg_Paint
}

// Shape represents a ThorVG shape
type Shape struct {
	Paint
}

// NewShape creates a new shape
func NewShape() (*Shape, error) {
	ptr := C.tvg_shape_new()
	if ptr == nil {
		return nil, errors.New("thorvg: failed to create shape")
	}
	return &Shape{Paint: Paint{ptr: ptr}}, nil
}

// AppendRect adds a rectangle to the shape
func (s *Shape) AppendRect(x, y, w, h float32) error {
	result := C.tvg_shape_append_rect(
		s.ptr,
		C.float(x),
		C.float(y),
		C.float(w),
		C.float(h),
		C.float(0),
		C.float(0),
	)
	return Result(result).Error()
}

// AppendCircle adds a circle to the shape
func (s *Shape) AppendCircle(cx, cy, rx, ry float32) error {
	result := C.tvg_shape_append_circle(
		s.ptr,
		C.float(cx),
		C.float(cy),
		C.float(rx),
		C.float(ry),
	)
	return Result(result).Error()
}

// SetFillColor sets the fill color for the shape
func (s *Shape) SetFillColor(r, g, b, a uint8) error {
	result := C.tvg_shape_set_fill_color(
		s.ptr,
		C.uint8_t(r),
		C.uint8_t(g),
		C.uint8_t(b),
		C.uint8_t(a),
	)
	return Result(result).Error()
}

// SetStrokeColor sets the stroke color for the shape
func (s *Shape) SetStrokeColor(r, g, b, a uint8) error {
	result := C.tvg_shape_set_stroke_color(
		s.ptr,
		C.uint8_t(r),
		C.uint8_t(g),
		C.uint8_t(b),
		C.uint8_t(a),
	)
	return Result(result).Error()
}

// SetStrokeWidth sets the stroke width for the shape
func (s *Shape) SetStrokeWidth(width float32) error {
	result := C.tvg_shape_set_stroke_width(s.ptr, C.float(width))
	return Result(result).Error()
}

// Picture represents a ThorVG picture
type Picture struct {
	Paint
}

// NewPicture creates a new picture
func NewPicture() (*Picture, error) {
	ptr := C.tvg_picture_new()
	if ptr == nil {
		return nil, errors.New("thorvg: failed to create picture")
	}
	return &Picture{Paint: Paint{ptr: ptr}}, nil
}

// Load loads an image file into the picture
func (p *Picture) Load(path string) error {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	result := C.tvg_picture_load(p.ptr, cPath)
	return Result(result).Error()
}

// LoadData loads image data into the picture
func (p *Picture) LoadData(data []byte, mimeType string) error {
	cMimeType := C.CString(mimeType)
	defer C.free(unsafe.Pointer(cMimeType))
	result := C.tvg_picture_load_data(
		p.ptr,
		(*C.char)(unsafe.Pointer(&data[0])),
		C.uint(len(data)),
		cMimeType,
		C.bool(false),
	)
	return Result(result).Error()
}

// SetSize sets the size of the picture
func (p *Picture) SetSize(w, h float32) error {
	result := C.tvg_picture_set_size(p.ptr, C.float(w), C.float(h))
	return Result(result).Error()
}

// GetSize gets the size of the picture
func (p *Picture) GetSize() (float32, float32, error) {
	var w, h C.float
	result := C.tvg_picture_get_size(p.ptr, &w, &h)
	if err := Result(result).Error(); err != nil {
		return 0, 0, err
	}
	return float32(w), float32(h), nil
}
