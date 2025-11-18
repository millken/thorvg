package thorvg_test

import (
	"testing"

	"github.com/millken/thorvg"
)

func TestShapeBasic(t *testing.T) {
	// 创建 Shape
	shape := thorvg.NewShape()
	if shape == nil {
		t.Fatal("Failed to create shape")
	}
	defer shape.Release()

	// 测试路径构建
	if err := shape.MoveTo(10, 10); err != nil {
		t.Errorf("MoveTo failed: %v", err)
	}

	if err := shape.LineTo(100, 10); err != nil {
		t.Errorf("LineTo failed: %v", err)
	}

	if err := shape.LineTo(100, 100); err != nil {
		t.Errorf("LineTo failed: %v", err)
	}

	if err := shape.Close(); err != nil {
		t.Errorf("Close failed: %v", err)
	}

	// 测试填充颜色
	if err := shape.SetFillColor(255, 0, 0, 255); err != nil {
		t.Errorf("SetFillColor failed: %v", err)
	}

	r, g, b, a, err := shape.GetFillColor()
	if err != nil {
		t.Errorf("GetFillColor failed: %v", err)
	}
	if r != 255 || g != 0 || b != 0 || a != 255 {
		t.Errorf("GetFillColor returned wrong values: %d, %d, %d, %d", r, g, b, a)
	}

	// 测试笔触
	if err := shape.SetStrokeWidth(2.0); err != nil {
		t.Errorf("SetStrokeWidth failed: %v", err)
	}

	width, err := shape.GetStrokeWidth()
	if err != nil {
		t.Errorf("GetStrokeWidth failed: %v", err)
	}
	if width != 2.0 {
		t.Errorf("GetStrokeWidth returned wrong value: %f", width)
	}
}

func TestGradientBasic(t *testing.T) {
	// 创建线性渐变
	grad := thorvg.NewLinearGradient()
	if grad == nil {
		t.Fatal("Failed to create linear gradient")
	}
	defer grad.Delete()

	// 设置渐变参数
	if err := grad.SetLinear(0, 0, 100, 100); err != nil {
		t.Errorf("SetLinear failed: %v", err)
	}

	x1, y1, x2, y2, err := grad.GetLinear()
	if err != nil {
		t.Errorf("GetLinear failed: %v", err)
	}
	if x1 != 0 || y1 != 0 || x2 != 100 || y2 != 100 {
		t.Errorf("GetLinear returned wrong values: %f, %f, %f, %f", x1, y1, x2, y2)
	}

	// 设置颜色点
	stops := []thorvg.ColorStop{
		{Offset: 0.0, R: 255, G: 0, B: 0, A: 255},
		{Offset: 1.0, R: 0, G: 0, B: 255, A: 255},
	}
	if err := grad.SetColorStops(stops); err != nil {
		t.Errorf("SetColorStops failed: %v", err)
	}

	getStops, err := grad.GetColorStops()
	if err != nil {
		t.Errorf("GetColorStops failed: %v", err)
	}
	if len(getStops) != 2 {
		t.Errorf("GetColorStops returned wrong count: %d", len(getStops))
	}
}

func TestSceneBasic(t *testing.T) {
	// 创建 Scene
	scene := thorvg.NewScene()
	if scene == nil {
		t.Fatal("Failed to create scene")
	}
	defer scene.Release()

	// 创建两个 Shape 用于测试
	shape1 := thorvg.NewShape()
	if shape1 == nil {
		t.Fatal("Failed to create shape1")
	}
	// 注意：不要释放shape1，因为所有权已转移给scene

	shape2 := thorvg.NewShape()
	if shape2 == nil {
		t.Fatal("Failed to create shape2")
	}
	// 注意：不要释放shape2，因为所有权已转移给scene

	// 设置 shape1 的路径和颜色
	if err := shape1.MoveTo(10, 10); err != nil {
		t.Errorf("Shape1 MoveTo failed: %v", err)
	}
	if err := shape1.LineTo(50, 10); err != nil {
		t.Errorf("Shape1 LineTo failed: %v", err)
	}
	if err := shape1.LineTo(50, 50); err != nil {
		t.Errorf("Shape1 LineTo failed: %v", err)
	}
	if err := shape1.Close(); err != nil {
		t.Errorf("Shape1 Close failed: %v", err)
	}
	if err := shape1.SetFillColor(255, 0, 0, 255); err != nil {
		t.Errorf("Shape1 SetFillColor failed: %v", err)
	}

	// 设置 shape2 的路径和颜色
	if err := shape2.MoveTo(60, 60); err != nil {
		t.Errorf("Shape2 MoveTo failed: %v", err)
	}
	if err := shape2.LineTo(100, 60); err != nil {
		t.Errorf("Shape2 LineTo failed: %v", err)
	}
	if err := shape2.LineTo(100, 100); err != nil {
		t.Errorf("Shape2 LineTo failed: %v", err)
	}
	if err := shape2.Close(); err != nil {
		t.Errorf("Shape2 Close failed: %v", err)
	}
	if err := shape2.SetFillColor(0, 255, 0, 255); err != nil {
		t.Errorf("Shape2 SetFillColor failed: %v", err)
	}

	// 测试 Push
	if err := scene.Push(shape1); err != nil {
		t.Errorf("Push shape1 failed: %v", err)
	}
	if err := scene.Push(shape2); err != nil {
		t.Errorf("Push shape2 failed: %v", err)
	}

	// 测试 Remove
	if err := scene.Remove(shape1); err != nil {
		t.Errorf("Remove shape1 failed: %v", err)
	}
	// 注意：Remove后不要释放shape1，因为可能所有权没有返回

	// 测试效果
	if err := scene.PushEffectGaussianBlur(2.0, 0, 0, 1); err != nil {
		t.Errorf("PushEffectGaussianBlur failed: %v", err)
	}

	if err := scene.ResetEffects(); err != nil {
		t.Errorf("ResetEffects failed: %v", err)
	}

	// 测试其他效果
	if err := scene.PushEffectDropShadow(0, 0, 0, 128, 45.0, 5.0, 2.0, 1); err != nil {
		t.Errorf("PushEffectDropShadow failed: %v", err)
	}

	if err := scene.PushEffectFill(255, 255, 255, 128); err != nil {
		t.Errorf("PushEffectFill failed: %v", err)
	}

	if err := scene.PushEffectTint(0, 0, 0, 255, 255, 255, 0.5); err != nil {
		t.Errorf("PushEffectTint failed: %v", err)
	}

	if err := scene.PushEffectTritone(0, 0, 0, 128, 128, 128, 255, 255, 255, 0); err != nil {
		t.Errorf("PushEffectTritone failed: %v", err)
	}
}
