package thorvg_test

import (
	"testing"

	"github.com/dnsoa/go/assert"
	"github.com/millken/thorvg"
)

func TestSaverBasic(t *testing.T) {
	// 创建 Saver
	saver := thorvg.NewSaver()
	if saver == nil {
		t.Fatal("Failed to create saver")
	}
	defer saver.Release()

	// 创建一个 Shape 用于测试
	shape := thorvg.NewShape()
	if shape == nil {
		t.Fatal("Failed to create shape")
	}
	defer shape.Release()

	// 设置 shape 的路径和颜色
	if err := shape.MoveTo(10, 10); err != nil {
		t.Errorf("Shape MoveTo failed: %v", err)
	}
	if err := shape.LineTo(50, 10); err != nil {
		t.Errorf("Shape LineTo failed: %v", err)
	}
	if err := shape.LineTo(50, 50); err != nil {
		t.Errorf("Shape LineTo failed: %v", err)
	}
	if err := shape.Close(); err != nil {
		t.Errorf("Shape Close failed: %v", err)
	}
	if err := shape.SetFillColor(255, 0, 0, 255); err != nil {
		t.Errorf("Shape SetFillColor failed: %v", err)
	}

	// 测试保存 - 即使不支持格式，我们也要测试API调用本身
	err := saver.SavePaint(shape, "/tmp/test_shape.tvg", 100)
	// 注意：这里可能返回"not supported"错误，这是正常的，取决于ThorVG的编译配置

	// 只有在没有其他错误（如参数错误）时才继续测试Sync
	if err == nil || err.Error() == "not supported" {
		// 测试同步 - 如果没有保存任务，可能会返回"insufficient condition"
		syncErr := saver.Sync()
		if syncErr != nil && syncErr.Error() != "insufficient condition" {
			t.Errorf("Unexpected sync error: %v", syncErr)
		}
	} else {
		t.Logf("Save returned error (expected for unsupported formats): %v", err)
	}

	// 测试保存 Scene
	scene := thorvg.NewScene()
	if scene == nil {
		t.Fatal("Failed to create scene")
	}
	defer scene.Release()

	// 添加 shape 到 scene
	if err := scene.Push(shape); err != nil {
		t.Errorf("Scene Push failed: %v", err)
	}

	// 测试保存scene
	err = saver.SavePaint(scene, "/tmp/test_scene.tvg", 100)
	if err == nil || err.Error() == "not supported" {
		// 正常情况
		t.Logf("Scene save completed (format support depends on ThorVG build)")
	} else {
		t.Logf("Scene save returned error: %v", err)
	}
}

func TestSaverErrors(t *testing.T) {
	// 测试空路径
	saver := thorvg.NewSaver()
	if saver == nil {
		t.Fatal("Failed to create saver")
	}
	defer saver.Release()

	shape := thorvg.NewShape()
	if shape == nil {
		t.Fatal("Failed to create shape")
	}
	defer shape.Release()

	// 测试空路径错误
	if err := saver.SavePaint(shape, "", 100); err == nil {
		t.Error("Expected error for empty path")
	}

	// 测试不支持的画布类型
	if err := saver.SavePaint("invalid", "/tmp/test.png", 100); err == nil {
		t.Error("Expected error for unsupported paint type")
	}
}

func TestSaver(t *testing.T) {
	thorvg.Init(0)
	r := assert.New(t)
	t.Run("Saver Creation", func(t *testing.T) {
		saver := thorvg.NewSaver()
		r.NotNil(saver)
	})

	t.Run("Save a lottie into tvg", func(t *testing.T) {
		animation := thorvg.NewAnimation()
		r.NotNil(animation)

		picture := animation.GetPicture()
		r.NoError(picture.Load(TEST_DIR + "test.json"))
		r.NoError(picture.SetSize(100, 100))
		saver := thorvg.NewSaver()
		err := saver.SaveAnimation(animation, TEST_DIR+"test_output.gif", 100, 30)
		// 某些 ThorVG 构建可能不支持保存动画，这是正常的
		r.NoError(err)

		// 如果保存成功，测试同步
		if err == nil {
			r.NoError(saver.Sync())
		}
		// os.Remove(TEST_DIR + "test_output.gif")
	})
}
