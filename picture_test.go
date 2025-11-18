package thorvg_test

import (
	"image"
	"image/color"
	"io"
	"os"
	"testing"

	"github.com/dnsoa/go/assert"
	"github.com/millken/thorvg"
)

func TestPicture2(t *testing.T) {
	r := assert.New(t)
	t.Run("Creation", func(t *testing.T) {
		picture := thorvg.NewPicture()
		r.NotNil(picture)
		pt, err := picture.GetType()
		r.NoError(err)
		r.Equal(thorvg.TypePicture, pt)
		defer picture.Release()
	})
	t.Run("Load RAW Data", func(t *testing.T) {
		file, err := os.Open(TEST_DIR + "/rawimage_200x300.raw")
		r.NoError(err)
		defer file.Close()
		data := make([]byte, 200*300*4) // 4 bytes per pixel for RGBA
		n, err := io.ReadFull(file, data)
		r.NoError(err)
		r.Equal(200*300*4, n)
		pic := thorvg.NewPicture()
		r.NotNil(pic)

		// Convert byte data to uint32
		uint32Data := make([]uint32, len(data)/4)
		for i := 0; i < len(uint32Data); i++ {
			uint32Data[i] = uint32(data[i*4]) | uint32(data[i*4+1])<<8 | uint32(data[i*4+2])<<16 | uint32(data[i*4+3])<<24
		}

		// Test loading raw data
		err = pic.LoadRaw(uint32Data, 200, 300, thorvg.ColorspaceABGR8888, false)
		r.NoError(err)

		w, h, err := pic.GetSize()
		r.NoError(err)
		r.Equal(float32(200), w)
		r.Equal(float32(300), h)
	})

	t.Run("Load RAW file and render", func(t *testing.T) {
		err := thorvg.Init(0)
		r.NoError(err)
		defer thorvg.Term()

		canvas := thorvg.NewSwCanvas(thorvg.EngineOptionDefault)
		r.NotNil(canvas)
		defer canvas.Destroy()

		buffer := make([]uint32, 100*100)
		err = canvas.SwCanvasSetTarget(buffer, 100, 100, 100, thorvg.ColorspaceABGR8888)
		r.NoError(err)

		file, err := os.Open(TEST_DIR + "/rawimage_200x300.raw")
		r.NoError(err)
		defer file.Close()
		data := make([]byte, 200*300*4)
		n, err := io.ReadFull(file, data)
		r.NoError(err)
		r.Equal(200*300*4, n)

		// Convert byte data to uint32
		uint32Data := make([]uint32, len(data)/4)
		for i := 0; i < len(uint32Data); i++ {
			uint32Data[i] = uint32(data[i*4]) | uint32(data[i*4+1])<<8 | uint32(data[i*4+2])<<16 | uint32(data[i*4+3])<<24
		}

		pic := thorvg.NewPicture()
		r.NotNil(pic)
		defer pic.Release()

		err = pic.LoadRaw(uint32Data, 200, 300, thorvg.ColorspaceABGR8888, false)
		r.NoError(err)

		err = pic.SetSize(100, 150)
		r.NoError(err)

		err = canvas.Push(&pic.Paint)
		r.NoError(err)
	})
}

func TestPictureBasic(t *testing.T) {
	// 创建 Picture
	picture := thorvg.NewPicture()
	if picture == nil {
		t.Fatal("Failed to create picture")
	}
	defer picture.Release()

	// 测试设置大小
	if err := picture.SetSize(100, 200); err != nil {
		t.Errorf("SetSize failed: %v", err)
	}

	// 测试获取大小（对于未加载图片的Picture，这可能会失败）
	w, h, err := picture.GetSize()
	if err != nil {
		t.Logf("GetSize returned error for unloaded picture (expected): %v", err)
	} else {
		if w != 100 || h != 200 {
			t.Errorf("GetSize returned wrong values: got (%f, %f), expected (100, 200)", w, h)
		}
	}

	// 测试设置原点
	if err := picture.SetOrigin(0.5, 0.5); err != nil {
		t.Errorf("SetOrigin failed: %v", err)
	}

	// 测试获取原点
	x, y, err := picture.GetOrigin()
	if err != nil {
		t.Errorf("GetOrigin failed: %v", err)
	} else {
		if x != 0.5 || y != 0.5 {
			t.Errorf("GetOrigin returned wrong values: got (%f, %f), expected (0.5, 0.5)", x, y)
		}
	}
}

func TestPictureLoad(t *testing.T) {
	picture := thorvg.NewPicture()
	if picture == nil {
		t.Fatal("Failed to create picture")
	}
	defer picture.Release()

	// 测试加载不存在的文件（应该失败）
	if err := picture.Load("/nonexistent/file.png"); err == nil {
		t.Error("Expected error for loading nonexistent file")
	}

	// 测试加载空路径（应该失败）
	if err := picture.Load(""); err == nil {
		t.Error("Expected error for loading empty path")
	}
}

func TestPictureLoadRaw(t *testing.T) {
	picture := thorvg.NewPicture()
	if picture == nil {
		t.Fatal("Failed to create picture")
	}
	defer picture.Release()

	// 创建一些测试数据
	data := []uint32{0xFF0000FF, 0x00FF00FF, 0x0000FFFF, 0xFFFFFF00} // RGBA数据

	// 测试加载原始数据
	if err := picture.LoadRaw(data, 2, 2, thorvg.ColorspaceABGR8888, true); err != nil {
		t.Errorf("LoadRaw failed: %v", err)
	}

	// 测试加载空数据（应该失败）
	if err := picture.LoadRaw([]uint32{}, 2, 2, thorvg.ColorspaceABGR8888, true); err == nil {
		t.Error("Expected error for loading empty raw data")
	}

	// 测试加载零宽度（应该失败）
	if err := picture.LoadRaw(data, 0, 2, thorvg.ColorspaceABGR8888, true); err == nil {
		t.Error("Expected error for loading with zero width")
	}

	// 测试加载零高度（应该失败）
	if err := picture.LoadRaw(data, 2, 0, thorvg.ColorspaceABGR8888, true); err == nil {
		t.Error("Expected error for loading with zero height")
	}
}

func TestPictureLoadData(t *testing.T) {
	picture := thorvg.NewPicture()
	if picture == nil {
		t.Fatal("Failed to create picture")
	}
	defer picture.Release()

	// 创建一些测试SVG数据
	svgData := []byte(`<svg width="100" height="100" xmlns="http://www.w3.org/2000/svg">
		<rect width="100" height="100" fill="red"/>
	</svg>`)

	// 测试加载SVG数据
	if err := picture.LoadData(svgData, "svg", "", true); err != nil {
		t.Errorf("LoadData failed: %v", err)
	}

	// 测试加载空数据（应该失败）
	if err := picture.LoadData([]byte{}, "svg", "", true); err == nil {
		t.Error("Expected error for loading empty data")
	}
}

func TestPictureGetPaint(t *testing.T) {
	picture := thorvg.NewPicture()
	if picture == nil {
		t.Fatal("Failed to create picture")
	}
	defer picture.Release()

	// 测试获取不存在的画笔（应该返回nil）
	paint := picture.GetPaint(999)
	if paint != nil {
		t.Error("Expected nil for nonexistent paint ID")
	}
}

func TestPictureSetAssetResolver(t *testing.T) {
	picture := thorvg.NewPicture()
	if picture == nil {
		t.Fatal("Failed to create picture")
	}
	defer picture.Release()

	// 测试设置资源解析器（应该返回错误，因为未实现）
	err := picture.SetAssetResolver(func(paint *thorvg.Paint, src string) bool {
		return true
	})

	if err == nil {
		t.Error("Expected error for unimplemented SetAssetResolver method")
	}

	expectedMsg := "asset resolver not implemented - requires complex callback handling"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestPictureErrors(t *testing.T) {
	// 测试空图片对象
	var picture *thorvg.Picture

	// 测试 nil picture 的错误处理
	// 注意：现在我们的错误处理会检查 nil 指针并返回错误，而不是 panic

	// 测试 SetSize 方法
	if err := picture.SetSize(100, 200); err == nil {
		t.Error("Expected error for nil picture SetSize")
	} else if err.Error() != "picture is nil" {
		t.Errorf("Expected 'picture is nil' error, got: %v", err)
	}

	// 测试 GetSize 方法
	if _, _, err := picture.GetSize(); err == nil {
		t.Error("Expected error for nil picture GetSize")
	} else if err.Error() != "picture is nil" {
		t.Errorf("Expected 'picture is nil' error, got: %v", err)
	}

	// 测试 SetOrigin 方法
	if err := picture.SetOrigin(0, 0); err == nil {
		t.Error("Expected error for nil picture SetOrigin")
	} else if err.Error() != "picture is nil" {
		t.Errorf("Expected 'picture is nil' error, got: %v", err)
	}

	// 测试 GetOrigin 方法
	if _, _, err := picture.GetOrigin(); err == nil {
		t.Error("Expected error for nil picture GetOrigin")
	} else if err.Error() != "picture is nil" {
		t.Errorf("Expected 'picture is nil' error, got: %v", err)
	}

	// 测试 Load 方法
	if err := picture.Load("test.png"); err == nil {
		t.Error("Expected error for nil picture Load")
	} else if err.Error() != "picture is nil" {
		t.Errorf("Expected 'picture is nil' error, got: %v", err)
	}

	// 测试 LoadRaw 方法
	if err := picture.LoadRaw([]uint32{0xFF0000FF}, 1, 1, thorvg.ColorspaceABGR8888, true); err == nil {
		t.Error("Expected error for nil picture LoadRaw")
	} else if err.Error() != "picture is nil" {
		t.Errorf("Expected 'picture is nil' error, got: %v", err)
	}

	// 测试 LoadData 方法
	if err := picture.LoadData([]byte("test"), "png", "", true); err == nil {
		t.Error("Expected error for nil picture LoadData")
	} else if err.Error() != "picture is nil" {
		t.Errorf("Expected 'picture is nil' error, got: %v", err)
	}

	// 测试 LoadImage 方法
	if err := picture.LoadImage(image.NewRGBA(image.Rect(0, 0, 10, 10))); err == nil {
		t.Error("Expected error for nil picture LoadImage")
	} else if err.Error() != "picture is nil" {
		t.Errorf("Expected 'picture is nil' error, got: %v", err)
	}

	// 测试 SetAssetResolver 方法
	if err := picture.SetAssetResolver(func(*thorvg.Paint, string) bool { return true }); err == nil {
		t.Error("Expected error for nil picture SetAssetResolver")
	} else if err.Error() != "picture is nil" {
		t.Errorf("Expected 'picture is nil' error, got: %v", err)
	}

	t.Log("Nil picture tests completed - unified error handling works correctly")
}

func TestPictureLoadImage(t *testing.T) {
	// 创建 Picture
	picture := thorvg.NewPicture()
	if picture == nil {
		t.Fatal("Failed to create picture")
	}
	defer picture.Release()

	// 测试加载 nil 图片（应该失败）
	if err := picture.LoadImage(nil); err == nil {
		t.Error("Expected error for loading nil image")
	}

	// 创建一个简单的测试图片
	img := image.NewRGBA(image.Rect(0, 0, 2, 2))
	img.Set(0, 0, color.RGBA{255, 0, 0, 255})     // 红色
	img.Set(1, 0, color.RGBA{0, 255, 0, 255})     // 绿色
	img.Set(0, 1, color.RGBA{0, 0, 255, 255})     // 蓝色
	img.Set(1, 1, color.RGBA{255, 255, 255, 255}) // 白色

	// 测试加载图片
	if err := picture.LoadImage(img); err != nil {
		t.Errorf("LoadImage failed: %v", err)
	}

	// 验证图片大小
	w, h, err := picture.GetSize()
	if err != nil {
		t.Errorf("GetSize failed after LoadImage: %v", err)
	} else {
		if w != 2 || h != 2 {
			t.Errorf("Expected size (2, 2), got (%f, %f)", w, h)
		}
	}

	// 测试加载空图片（应该失败）
	emptyImg := image.NewRGBA(image.Rect(0, 0, 0, 0))
	if err := picture.LoadImage(emptyImg); err == nil {
		t.Error("Expected error for loading empty image")
	}
}

func TestPictureLoadImageWithDifferentFormats(t *testing.T) {
	picture := thorvg.NewPicture()
	if picture == nil {
		t.Fatal("Failed to create picture")
	}
	defer picture.Release()

	// 测试不同类型的图片
	testCases := []struct {
		name string
		img  image.Image
	}{
		{
			name: "RGBA",
			img:  image.NewRGBA(image.Rect(0, 0, 10, 10)),
		},
		{
			name: "NRGBA",
			img:  image.NewNRGBA(image.Rect(0, 0, 10, 10)),
		},
		{
			name: "YCbCr",
			img:  image.NewYCbCr(image.Rect(0, 0, 10, 10), image.YCbCrSubsampleRatio444),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 填充一些测试数据
			if rgba, ok := tc.img.(*image.RGBA); ok {
				for y := 0; y < 10; y++ {
					for x := 0; x < 10; x++ {
						rgba.Set(x, y, color.RGBA{
							R: uint8(x * 25),
							G: uint8(y * 25),
							B: 128,
							A: 255,
						})
					}
				}
			}

			if err := picture.LoadImage(tc.img); err != nil {
				t.Errorf("LoadImage failed for %s: %v", tc.name, err)
			}
		})
	}
}
