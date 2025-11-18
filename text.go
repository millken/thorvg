package thorvg

/*

#include <stdlib.h>
#include <string.h>
#include "thorvg_capi.h"
*/
import "C"
import (
	"unsafe"
)

// NewText 创建一个新的 Text 对象
func NewText() *Text {
	paint := C.tvg_text_new()
	if paint == nil {
		return nil
	}
	return &Text{Paint{ptr: paint}}
}

// SetFont 设置文本的字体
func (t *Text) SetFont(name string) error {
	if err := checkNilText(t); err != nil {
		return err
	}
	if err := checkEmptyString(name); err != nil {
		return err
	}
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return resultToError(C.tvg_text_set_font(t.ptr, cName))
}

// SetSize 设置文本的字体大小
func (t *Text) SetSize(size float32) error {
	if err := checkNilText(t); err != nil {
		return err
	}
	if err := checkPositiveValue(size, "size"); err != nil {
		return err
	}
	return resultToError(C.tvg_text_set_size(t.ptr, C.float(size)))
}

// SetText 设置要渲染的文本内容
func (t *Text) SetText(text string) error {
	if err := checkNilText(t); err != nil {
		return err
	}
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))
	return resultToError(C.tvg_text_set_text(t.ptr, cText))
}

// SetAlign 设置文本的对齐方式
func (t *Text) SetAlign(x, y float32) error {
	if err := checkNilText(t); err != nil {
		return err
	}
	return resultToError(C.tvg_text_align(t.ptr, C.float(x), C.float(y)))
}

// SetLayout 设置文本的布局框
func (t *Text) SetLayout(w, h float32) error {
	if err := checkNilText(t); err != nil {
		return err
	}
	if err := checkPositiveValue(w, "width"); err != nil {
		return err
	}
	if err := checkPositiveValue(h, "height"); err != nil {
		return err
	}
	return resultToError(C.tvg_text_layout(t.ptr, C.float(w), C.float(h)))
}

// SetWrapMode 设置文本的换行模式
func (t *Text) SetWrapMode(mode TextWrap) error {
	if err := checkNilText(t); err != nil {
		return err
	}
	return resultToError(C.tvg_text_wrap_mode(t.ptr, C.Tvg_Text_Wrap(mode)))
}

// SetItalic 设置文本的斜体效果
func (t *Text) SetItalic(shear float32) error {
	if err := checkNilText(t); err != nil {
		return err
	}
	return resultToError(C.tvg_text_set_italic(t.ptr, C.float(shear)))
}

// SetOutline 设置文本的轮廓
func (t *Text) SetOutline(width float32, r, g, b uint8) error {
	if err := checkNilText(t); err != nil {
		return err
	}
	if err := checkPositiveValue(width, "width"); err != nil {
		return err
	}
	return resultToError(C.tvg_text_set_outline(t.ptr, C.float(width), C.uint8_t(r), C.uint8_t(g), C.uint8_t(b)))
}

// SetColor 设置文本的填充颜色
func (t *Text) SetColor(r, g, b uint8) error {
	if err := checkNilText(t); err != nil {
		return err
	}
	return resultToError(C.tvg_text_set_color(t.ptr, C.uint8_t(r), C.uint8_t(g), C.uint8_t(b)))
}

// SetGradient 设置文本的渐变填充
func (t *Text) SetGradient(gradient *Gradient) error {
	if err := checkNilText(t); err != nil {
		return err
	}
	if err := checkNilGradient(gradient); err != nil {
		return err
	}
	return resultToError(C.tvg_text_set_gradient(t.ptr, gradient.ptr))
}
