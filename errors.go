package thorvg

import "fmt"

// 定义统一的错误类型
var (
	// 基础错误
	ErrNilPointer   = fmt.Errorf("nil pointer")
	ErrNilPaint     = fmt.Errorf("paint is nil")
	ErrNilCanvas    = fmt.Errorf("canvas is nil")
	ErrNilGradient  = fmt.Errorf("gradient is nil")
	ErrNilScene     = fmt.Errorf("scene is nil")
	ErrNilText      = fmt.Errorf("text is nil")
	ErrNilShape     = fmt.Errorf("shape is nil")
	ErrNilPicture   = fmt.Errorf("picture is nil")
	ErrNilAnimation = fmt.Errorf("animation is nil")
	ErrNilSaver     = fmt.Errorf("saver is nil")
	ErrNilAccessor  = fmt.Errorf("accessor is nil")
	ErrNilMatrix    = fmt.Errorf("matrix is nil")

	// 参数错误
	ErrInvalidArgument = fmt.Errorf("invalid argument")
	ErrEmptyString     = fmt.Errorf("empty string")
	ErrEmptySlice      = fmt.Errorf("empty slice")
	ErrBufferTooSmall  = fmt.Errorf("buffer too small")
	ErrInvalidSize     = fmt.Errorf("invalid size")
	ErrInvalidIndex    = fmt.Errorf("invalid index")

	// 操作错误
	ErrOperationFailed = fmt.Errorf("operation failed")
	ErrNotSupported    = fmt.Errorf("not supported")
	ErrNotImplemented  = fmt.Errorf("not implemented")

	// 资源错误
	ErrResourceLeak     = fmt.Errorf("resource leak")
	ErrMemoryCorruption = fmt.Errorf("memory corruption")
	ErrFailedAllocation = fmt.Errorf("failed allocation")
)

// checkNilPointer 统一的 nil 指针检查
func checkNilPointer(ptr interface{}, err error) error {
	if ptr == nil {
		if err != nil {
			return err
		}
		return ErrNilPointer
	}
	return nil
}

// checkNilPaint 检查 Paint 是否为 nil
func checkNilPaint(p *Paint) error {
	if p == nil || p.ptr == nil {
		return ErrNilPaint
	}
	return nil
}

// checkNilCanvas 检查 Canvas 是否为 nil
func checkNilCanvas(c *Canvas) error {
	if c == nil || c.ptr == nil {
		return ErrNilCanvas
	}
	return nil
}

// checkNilGradient 检查 Gradient 是否为 nil
func checkNilGradient(g *Gradient) error {
	if g == nil || g.ptr == nil {
		return ErrNilGradient
	}
	return nil
}

// checkNilScene 检查 Scene 是否为 nil
func checkNilScene(s *Scene) error {
	if s == nil || s.ptr == nil {
		return ErrNilScene
	}
	return nil
}

// checkNilText 检查 Text 是否为 nil
func checkNilText(t *Text) error {
	if t == nil || t.ptr == nil {
		return ErrNilText
	}
	return nil
}

// checkNilShape 检查 Shape 是否为 nil
func checkNilShape(s *Shape) error {
	if s == nil || s.ptr == nil {
		return ErrNilShape
	}
	return nil
}

// checkNilPicture 检查 Picture 是否为 nil
func checkNilPicture(p *Picture) error {
	if p == nil || p.ptr == nil {
		return ErrNilPicture
	}
	return nil
}

// checkNilAnimation 检查 Animation 是否为 nil
func checkNilAnimation(a *Animation) error {
	if a == nil || a.ptr == nil {
		return ErrNilAnimation
	}
	return nil
}

// checkNilSaver 检查 Saver 是否为 nil
func checkNilSaver(s *Saver) error {
	if s == nil || s.ptr == nil {
		return ErrNilSaver
	}
	return nil
}

// checkNilAccessor 检查 Accessor 是否为 nil
func checkNilAccessor(a *Accessor) error {
	if a == nil || a.ptr == nil {
		return ErrNilAccessor
	}
	return nil
}

// checkNilMatrix 检查 Matrix 是否为 nil
func checkNilMatrix(m *Matrix) error {
	if m == nil {
		return ErrNilMatrix
	}
	return nil
}

// checkEmptyString 检查字符串是否为空
func checkEmptyString(s string) error {
	if s == "" {
		return ErrEmptyString
	}
	return nil
}

// checkEmptySlice 检查切片是否为空
func checkEmptySlice[T any](slice []T) error {
	if len(slice) == 0 {
		return ErrEmptySlice
	}
	return nil
}

// checkBufferSize 检查缓冲区大小是否足够
func checkBufferSize(actual, required int) error {
	if actual < required {
		return fmt.Errorf("%w: required %d, got %d", ErrBufferTooSmall, required, actual)
	}
	return nil
}

// checkPositiveValue 检查值是否为正数
func checkPositiveValue[T int | int32 | uint32 | float32 | float64](value T, name string) error {
	switch v := any(value).(type) {
	case int:
		if v <= 0 {
			return fmt.Errorf("%w: %s must be positive, got %d", ErrInvalidArgument, name, v)
		}
	case int32:
		if v <= 0 {
			return fmt.Errorf("%w: %s must be positive, got %d", ErrInvalidArgument, name, v)
		}
	case uint32:
		if v == 0 {
			return fmt.Errorf("%w: %s must be positive, got %d", ErrInvalidArgument, name, v)
		}
	case float32:
		if v <= 0 {
			return fmt.Errorf("%w: %s must be positive, got %f", ErrInvalidArgument, name, v)
		}
	case float64:
		if v <= 0 {
			return fmt.Errorf("%w: %s must be positive, got %f", ErrInvalidArgument, name, v)
		}
	}
	return nil
}

// checkRangeInt 检查整数是否在指定范围内
func checkRangeInt(value, min, max int, name string) error {
	if value < min || value > max {
		return fmt.Errorf("%w: %s must be in range [%d, %d], got %d", ErrInvalidArgument, name, min, max, value)
	}
	return nil
}

// checkRangeInt32 检查 int32 是否在指定范围内
func checkRangeInt32(value, min, max int32, name string) error {
	if value < min || value > max {
		return fmt.Errorf("%w: %s must be in range [%d, %d], got %d", ErrInvalidArgument, name, min, max, value)
	}
	return nil
}

// checkRangeUint32 检查 uint32 是否在指定范围内
func checkRangeUint32(value, min, max uint32, name string) error {
	if value < min || value > max {
		return fmt.Errorf("%w: %s must be in range [%d, %d], got %d", ErrInvalidArgument, name, min, max, value)
	}
	return nil
}

// checkRangeFloat32 检查 float32 是否在指定范围内
func checkRangeFloat32(value, min, max float32, name string) error {
	if value < min || value > max {
		return fmt.Errorf("%w: %s must be in range [%f, %f], got %f", ErrInvalidArgument, name, min, max, value)
	}
	return nil
}

// checkRangeFloat64 检查 float64 是否在指定范围内
func checkRangeFloat64(value, min, max float64, name string) error {
	if value < min || value > max {
		return fmt.Errorf("%w: %s must be in range [%f, %f], got %f", ErrInvalidArgument, name, min, max, value)
	}
	return nil
}
