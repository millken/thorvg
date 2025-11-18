package thorvg_test

import (
	"os"
	"testing"

	"github.com/dnsoa/go/assert"
	"github.com/millken/thorvg"
)

const (
	TEST_DIR = "./resources/"
)

func TestText(t *testing.T) {
	r := assert.New(t)
	t.Run("Load SVG Data", func(t *testing.T) {
		text := thorvg.NewText()
		r.NotNil(text)
		defer text.Release()

		tt, err := text.GetType()
		r.NoError(err)
		r.Equal(thorvg.TypeText, tt)
	})
	t.Run("Load TTF Data from a file", func(t *testing.T) {
		err := thorvg.Init(0)
		r.NoError(err)
		defer thorvg.Term()

		text := thorvg.NewText()
		r.NotNil(text)
		defer text.Release()

		// Test unloading invalid font
		err = thorvg.UnloadFont(TEST_DIR + "/invalid.ttf")
		r.Error(err) // Should fail for invalid font

		// Test loading font
		err = thorvg.LoadFont(TEST_DIR + "/Arial.ttf")
		r.NoError(err)

		// Test loading invalid font
		err = thorvg.LoadFont(TEST_DIR + "/invalid.ttf")
		r.Error(err)

		// Test loading with empty path
		err = thorvg.LoadFont("")
		r.Error(err)

		// Test unloading font
		err = thorvg.UnloadFont(TEST_DIR + "/Arial.ttf")
		r.NoError(err)
	})

	t.Run("Load TTF Data from a buffer", func(t *testing.T) {
		data, err := os.ReadFile(TEST_DIR + "/Arial.ttf")
		r.NoError(err)

		svg := "<svg height=\"1000\" viewBox=\"0 0 600 600\" ></svg>"

		// Test loading with empty name
		err = thorvg.LoadFontData("", data, "", false)
		r.Error(err)

		// Test loading with unknown mimetype
		err = thorvg.LoadFontData("ArialSvg", []byte(svg), "unknown", false)
		r.Error(err) // SVG data is not supported as font

		// Test loading font data
		err = thorvg.LoadFontData("ArialUnknown", data, "unknown", false)
		r.NoError(err)

		err = thorvg.LoadFontData("ArialTtf", data, "ttf", true)
		r.NoError(err)

		err = thorvg.LoadFontData("Arial", data, "", false)
		r.NoError(err)

		// Test unloading loaded fonts
		err = thorvg.UnloadFont("ArialUnknown")
		r.Error(err) // Should fail as it's loaded by name, not path

		// Note: UnloadFont is for path-based fonts, not name-based
	})

	t.Run("Text Font", func(t *testing.T) {
		r := assert.New(t)
		err := thorvg.Init(0)
		r.NoError(err)
		defer thorvg.Term()

		text := thorvg.NewText()
		r.NotNil(text)
		defer text.Release()

		err = thorvg.LoadFont(TEST_DIR + "/Arial.ttf")
		r.NoError(err)

		err = text.SetFont("Arial")
		r.NoError(err)

		err = text.SetSize(80)
		r.NoError(err)

		err = text.SetSize(1)
		r.NoError(err)

		err = text.SetSize(50)
		r.NoError(err)

		err = text.SetFont("")
		r.Error(err) // Setting empty font name fails with "insufficient condition"

		err = text.SetFont("InvalidFont")
		r.Error(err)
	})

	t.Run("Text Basic", func(t *testing.T) {
		err := thorvg.Init(0)
		r.NoError(err)
		defer thorvg.Term()

		text := thorvg.NewText()
		r.NotNil(text)
		// Note: Don't release text after pushing to canvas, canvas manages lifecycle

		canvas := thorvg.NewSwCanvas(thorvg.EngineOptionDefault)
		r.NotNil(canvas)
		defer canvas.Destroy()

		buffer := make([]uint32, 100*100)
		err = canvas.SwCanvasSetTarget(buffer, 100, 100, 100, thorvg.ColorspaceABGR8888)
		r.NoError(err)

		err = thorvg.LoadFont(TEST_DIR + "/Arial.ttf")
		r.NoError(err)

		err = text.SetFont("Arial")
		r.NoError(err)

		err = text.SetSize(80)
		r.NoError(err)

		err = text.SetText("")
		r.NoError(err)

		err = text.SetText("ABCDEFGHIJIKLMOPQRSTUVWXYZ")
		r.NoError(err)

		err = text.SetText("THORVG Text")
		r.NoError(err)

		err = text.SetColor(255, 255, 255)
		r.NoError(err)

		err = canvas.Push(&text.Paint)
		r.NoError(err)
		// Canvas now owns the text object, don't release it manually
	})
	t.Run("Text with composite glyphs", func(t *testing.T) {
		err := thorvg.Init(0)
		r.NoError(err)
		defer thorvg.Term()

		text := thorvg.NewText()
		r.NotNil(text)
		// Note: Don't release text after pushing to canvas, canvas manages lifecycle

		canvas := thorvg.NewSwCanvas(thorvg.EngineOptionDefault)
		r.NotNil(canvas)
		defer canvas.Destroy()

		buffer := make([]uint32, 100*100)
		err = canvas.SwCanvasSetTarget(buffer, 100, 100, 100, thorvg.ColorspaceABGR8888)
		r.NoError(err)

		err = thorvg.LoadFont(TEST_DIR + "/Arial.ttf")
		r.NoError(err)

		err = text.SetFont("Arial")
		r.NoError(err)

		err = text.SetSize(80)
		r.NoError(err)

		err = text.SetText("\xc5\xbb\x6f\xc5\x82\xc4\x85\x64\xc5\xba \xc8\xab")
		r.NoError(err)

		err = text.SetColor(255, 255, 255)
		r.NoError(err)

		err = canvas.Push(&text.Paint)
		r.NoError(err)
		// Canvas now owns the text object, don't release it manually
	})
}

// func TestSaver(t *testing.T) {
// 	r := assert.New(t)
// 	t.Run("Saver Creation", func(t *testing.T) {
// 		saver := thorvg.NewSaver()
// 		r.NotNil(saver)
// 	})

// 	t.Run("Save a lottie into gif", func(t *testing.T) {
// 		r := assert.New(t)
// 		err := thorvg.Init(0)
// 		r.NoError(err)
// 		defer thorvg.Term()

// 		animation := thorvg.NewAnimation()
// 		r.NotNil(animation)
// 		defer animation.Release()

// 		picture := animation.GetPicture()
// 		r.NotNil(picture)

// 		// Load a simple SVG instead of JSON for saver test
// 		data := []byte(`<svg width="100" height="100"><rect width="100" height="100" fill="red"/></svg>`)
// 		err = picture.LoadData(data, "svg", "", false)
// 		r.NoError(err)

// 		err = picture.SetSize(100, 100)
// 		r.NoError(err)

// 		saver := thorvg.NewSaver()
// 		r.NotNil(saver)
// 		defer saver.Release()

// 		// Save the picture
// 		err = saver.Save(picture, TEST_DIR+"test_output.tvg", 100)
// 		if err != nil && err.Error() == "not supported" {
// 			// TVG format may not be supported in this build, skip test
// 			t.Skip("TVG format not supported in this ThorVG build")
// 		}
// 		r.NoError(err)

// 		err = saver.Sync()
// 		r.NoError(err)
// 	})
// }
