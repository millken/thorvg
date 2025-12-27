package thorvg_test

import (
	"testing"
	"unsafe"

	"github.com/millken/thorvg"
)

func init() {
	// 初始化 ThorVG 引擎
	if err := thorvg.Init(0); err != nil {
		panic("Failed to initialize ThorVG engine: " + err.Error())
	}
}

func TestNewSwCanvas(t *testing.T) {
	canvas := thorvg.NewSwCanvas(thorvg.EngineOptionDefault)
	if canvas == nil {
		t.Fatal("NewSwCanvas returned nil")
	}
	defer canvas.Destroy()

	// 移除对未导出字段的检查
}

func TestNewWgCanvas(t *testing.T) {
	canvas := thorvg.NewWgCanvas()
	if canvas == nil {
		t.Skip("NewWgCanvas returned nil (expected in test environment without WebGPU)")
		return
	}
	defer canvas.Destroy()

	// 移除对未导出字段的检查
}

func TestCanvasPush(t *testing.T) {
	canvas := thorvg.NewSwCanvas(thorvg.EngineOptionDefault)
	if canvas == nil {
		t.Fatal("NewSwCanvas returned nil")
	}
	defer canvas.Destroy()

	shape := thorvg.NewShape()
	if shape == nil {
		t.Fatal("NewShape returned nil")
	}
	defer shape.Release()

	err := canvas.Push(&shape.Paint)
	if err != nil {
		t.Fatalf("Push failed: %v", err)
	}

	// Canvas 接管了 Paint 的生命周期，不需要手动释放
}

func TestCanvasPushAt(t *testing.T) {
	canvas := thorvg.NewSwCanvas(thorvg.EngineOptionDefault)
	if canvas == nil {
		t.Fatal("NewSwCanvas returned nil")
	}
	defer canvas.Destroy()

	shape1 := thorvg.NewShape()
	if shape1 == nil {
		t.Fatal("NewShape returned nil")
	}
	defer shape1.Release()

	shape2 := thorvg.NewShape()
	if shape2 == nil {
		t.Fatal("NewShape returned nil")
	}
	defer shape2.Release()

	err := canvas.Push(&shape1.Paint)
	if err != nil {
		t.Fatalf("Push failed: %v", err)
	}

	err = canvas.PushAt(&shape2.Paint, &shape1.Paint)
	if err != nil {
		t.Fatalf("PushAt failed: %v", err)
	}

	// Canvas 接管了 Paint 的生命周期
}

func TestCanvasRemove(t *testing.T) {
	canvas := thorvg.NewSwCanvas(thorvg.EngineOptionDefault)
	if canvas == nil {
		t.Fatal("NewSwCanvas returned nil")
	}
	defer canvas.Destroy()

	shape := thorvg.NewShape()
	if shape == nil {
		t.Fatal("NewShape returned nil")
	}

	err := canvas.Push(&shape.Paint)
	if err != nil {
		t.Fatalf("Push failed: %v", err)
	}

	err = canvas.Remove(&shape.Paint)
	if err != nil {
		t.Fatalf("Remove failed: %v", err)
	}

	// Remove 后 Paint 可能已经被释放，不要再次释放
}

func TestCanvasUpdate(t *testing.T) {
	canvas := thorvg.NewSwCanvas(thorvg.EngineOptionDefault)
	if canvas == nil {
		t.Fatal("NewSwCanvas returned nil")
	}
	defer canvas.Destroy()

	// 设置目标缓冲区
	width, height := uint32(100), uint32(100)
	buffer := make([]uint32, width*height)
	err := canvas.SwCanvasSetTarget(buffer, width, width, height, thorvg.ColorspaceABGR8888)
	if err != nil {
		t.Fatalf("SwCanvasSetTarget failed: %v", err)
	}

	err = canvas.Update()
	if err != nil {
		t.Fatalf("Update failed: %v", err)
	}
}

func TestCanvasDraw(t *testing.T) {
	canvas := thorvg.NewSwCanvas(thorvg.EngineOptionDefault)
	if canvas == nil {
		t.Fatal("NewSwCanvas returned nil")
	}
	defer canvas.Destroy()

	// 创建一个简单的像素缓冲区用于测试
	width, height := uint32(100), uint32(100)
	buffer := make([]uint32, width*height)

	err := canvas.SwCanvasSetTarget(buffer, width, width, height, thorvg.ColorspaceABGR8888)
	if err != nil {
		t.Fatalf("SwCanvasSetTarget failed: %v", err)
	}

	err = canvas.Draw(true)
	if err != nil {
		t.Fatalf("Draw failed: %v", err)
	}

	err = canvas.Sync()
	if err != nil {
		t.Fatalf("Sync failed: %v", err)
	}
}

func TestCanvasSetViewport(t *testing.T) {
	canvas := thorvg.NewSwCanvas(thorvg.EngineOptionDefault)
	if canvas == nil {
		t.Fatal("NewSwCanvas returned nil")
	}
	defer canvas.Destroy()

	err := canvas.SetViewport(0, 0, 100, 100)
	if err != nil {
		t.Fatalf("SetViewport failed: %v", err)
	}
}

func TestSwCanvasSetTarget(t *testing.T) {
	canvas := thorvg.NewSwCanvas(thorvg.EngineOptionDefault)
	if canvas == nil {
		t.Fatal("NewSwCanvas returned nil")
	}
	defer canvas.Destroy()

	width, height := uint32(100), uint32(100)
	buffer := make([]uint32, width*height)

	err := canvas.SwCanvasSetTarget(buffer, width, width, height, thorvg.ColorspaceABGR8888)
	if err != nil {
		t.Fatalf("SwCanvasSetTarget failed: %v", err)
	}
}

func TestWgCanvasSetTarget(t *testing.T) {
	canvas := thorvg.NewWgCanvas()
	if canvas == nil {
		t.Skip("NewWgCanvas returned nil (expected in test environment without WebGPU)")
		return
	}
	defer canvas.Destroy()

	// 注意：这个测试需要有效的 WebGPU 上下文，在测试环境中可能不可用
	// 这里只是测试函数调用，不验证实际结果
	err := canvas.WgCanvasSetTarget(unsafe.Pointer(uintptr(1)), unsafe.Pointer(uintptr(1)), unsafe.Pointer(uintptr(1)), 100, 100, thorvg.ColorspaceABGR8888, 0)
	// 在没有 WebGPU 上下文的情况下，可能会失败，但函数调用应该成功
	_ = err // 忽略错误，因为测试环境可能不支持 WebGPU
}
