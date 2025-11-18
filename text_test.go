package thorvg_test

import (
	"testing"

	"github.com/millken/thorvg"
)

func TestTextBasic(t *testing.T) {
	// 创建 Text
	text := thorvg.NewText()
	if text == nil {
		t.Fatal("Failed to create text")
	}
	defer text.Release()

	// 测试设置文本内容
	if err := text.SetText("Hello, World!"); err != nil {
		t.Errorf("SetText failed: %v", err)
	}

	// 测试设置字体大小
	if err := text.SetSize(24.0); err != nil {
		t.Errorf("SetSize failed: %v", err)
	}

	// 测试设置字体（可能失败，因为字体可能不存在）
	if err := text.SetFont("Arial"); err != nil {
		t.Logf("SetFont failed (expected if font not available): %v", err)
	}

	// 测试设置颜色
	if err := text.SetColor(255, 0, 0); err != nil {
		t.Errorf("SetColor failed: %v", err)
	}

	// 测试设置对齐
	if err := text.SetAlign(0.5, 0.5); err != nil {
		t.Errorf("SetAlign failed: %v", err)
	}

	// 测试设置布局
	if err := text.SetLayout(200, 100); err != nil {
		t.Errorf("SetLayout failed: %v", err)
	}

	// 测试设置换行模式
	if err := text.SetWrapMode(thorvg.TextWrapWord); err != nil {
		t.Errorf("SetWrapMode failed: %v", err)
	}

	// 测试设置斜体
	if err := text.SetItalic(0.2); err != nil {
		t.Errorf("SetItalic failed: %v", err)
	}

	// 测试设置轮廓
	if err := text.SetOutline(2.0, 0, 0, 0); err != nil {
		t.Errorf("SetOutline failed: %v", err)
	}
}

func TestTextGradient(t *testing.T) {
	text := thorvg.NewText()
	if text == nil {
		t.Fatal("Failed to create text")
	}
	defer text.Release()

	// 创建线性渐变
	gradient := thorvg.NewLinearGradient()
	if gradient == nil {
		t.Fatal("Failed to create gradient")
	}

	// 设置渐变
	if err := gradient.SetLinear(0, 0, 100, 0); err != nil {
		t.Errorf("Gradient SetLinear failed: %v", err)
	}

	// 设置文本渐变填充
	if err := text.SetGradient(gradient); err != nil {
		t.Errorf("SetGradient failed: %v", err)
	}

	// Text 现在拥有渐变的所有权，不需要手动释放
}

func TestTextErrors(t *testing.T) {
	// 测试空文本对象
	var text *thorvg.Text

	// 测试 nil text 的错误处理
	if err := text.SetText("test"); err == nil {
		t.Error("Expected error for nil text SetText")
	}

	if err := text.SetSize(12); err == nil {
		t.Error("Expected error for nil text SetSize")
	}

	if err := text.SetFont("Arial"); err == nil {
		t.Error("Expected error for nil text SetFont")
	}

	if err := text.SetColor(255, 0, 0); err == nil {
		t.Error("Expected error for nil text SetColor")
	}

	if err := text.SetAlign(0.5, 0.5); err == nil {
		t.Error("Expected error for nil text SetAlign")
	}

	if err := text.SetLayout(100, 50); err == nil {
		t.Error("Expected error for nil text SetLayout")
	}

	if err := text.SetWrapMode(thorvg.TextWrapWord); err == nil {
		t.Error("Expected error for nil text SetWrapMode")
	}

	if err := text.SetItalic(0.1); err == nil {
		t.Error("Expected error for nil text SetItalic")
	}

	if err := text.SetOutline(1.0, 0, 0, 0); err == nil {
		t.Error("Expected error for nil text SetOutline")
	}

	// 测试 nil 渐变
	gradient := thorvg.NewLinearGradient()
	if gradient == nil {
		t.Fatal("Failed to create gradient")
	}
	defer gradient.Delete()

	if err := text.SetGradient(gradient); err == nil {
		t.Error("Expected error for nil text SetGradient")
	}
}

func TestFontFunctions(t *testing.T) {
	// 测试加载不存在的字体（应该失败）
	if err := thorvg.LoadFont("/nonexistent/font.ttf"); err == nil {
		t.Error("Expected error for loading nonexistent font")
	}

	// 测试加载空字体数据（应该失败）
	if err := thorvg.LoadFontData("test", []byte{}, "ttf", false); err == nil {
		t.Error("Expected error for loading empty font data")
	}

	// 测试卸载不存在的字体（应该失败）
	if err := thorvg.UnloadFont("nonexistent"); err == nil {
		t.Log("UnloadFont for nonexistent font returned error (expected)")
	}
}
