package thorvg

/*
#include <stdlib.h>
#include <string.h>
#include "thorvg_capi.h"
*/
import "C"
import (
	"errors"
	"image"
	"unsafe"
)

// NewPicture 创建一个新的 Picture 对象
func NewPicture() *Picture {
	paint := C.tvg_picture_new()
	if paint == nil {
		return nil
	}
	return &Picture{Paint{ptr: paint}}
}

// Load 从文件加载图片
func (p *Picture) Load(path string) error {
	if err := checkNilPicture(p); err != nil {
		return err
	}
	if err := checkEmptyString(path); err != nil {
		return err
	}
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	return resultToError(C.tvg_picture_load(p.ptr, cPath))
}

// LoadRaw 从内存加载原始图片数据
func (p *Picture) LoadRaw(data []uint32, w, h uint32, cs Colorspace, copy bool) error {
	if err := checkNilPicture(p); err != nil {
		return err
	}
	if err := checkEmptySlice(data); err != nil {
		return err
	}
	if err := checkPositiveValue(w, "width"); err != nil {
		return err
	}
	if err := checkPositiveValue(h, "height"); err != nil {
		return err
	}
	return resultToError(C.tvg_picture_load_raw(p.ptr, (*C.uint32_t)(unsafe.Pointer(&data[0])), C.uint32_t(w), C.uint32_t(h), C.Tvg_Colorspace(cs), C.bool(copy)))
}

// LoadData 从内存数据加载图片
func (p *Picture) LoadData(data []byte, mimetype, rpath string, copy bool) error {
	if err := checkNilPicture(p); err != nil {
		return err
	}
	if err := checkEmptySlice(data); err != nil {
		return err
	}

	cMimetype := C.CString(mimetype)
	cRpath := C.CString(rpath)
	defer func() {
		C.free(unsafe.Pointer(cMimetype))
		C.free(unsafe.Pointer(cRpath))
	}()

	return resultToError(C.tvg_picture_load_data(p.ptr, (*C.char)(unsafe.Pointer(&data[0])), C.uint32_t(len(data)), cMimetype, cRpath, C.bool(copy)))
}

// SetAssetResolver 设置资源解析器回调
// 注意：这个方法需要特殊的实现来处理Go函数指针到C函数指针的转换
// 暂时未实现
func (p *Picture) SetAssetResolver(resolver func(*Paint, string) bool) error {
	if p == nil || p.ptr == nil {
		return errors.New("picture is nil")
	}
	if resolver == nil {
		return errors.New("resolver is nil")
	}

	// TODO: 实现Go函数指针到C函数指针的转换
	// 这需要一个复杂的机制来注册回调函数
	return errors.New("asset resolver not implemented - requires complex callback handling")
}

// SetSize 设置图片大小
func (p *Picture) SetSize(w, h float32) error {
	if err := checkNilPicture(p); err != nil {
		return err
	}
	if err := checkPositiveValue(w, "width"); err != nil {
		return err
	}
	if err := checkPositiveValue(h, "height"); err != nil {
		return err
	}
	return resultToError(C.tvg_picture_set_size(p.ptr, C.float(w), C.float(h)))
}

// GetSize 获取图片大小
func (p *Picture) GetSize() (w, h float32, err error) {
	if err = checkNilPicture(p); err != nil {
		return
	}
	var cW, cH C.float
	res := C.tvg_picture_get_size(p.ptr, &cW, &cH)
	err = resultToError(res)
	w = float32(cW)
	h = float32(cH)
	return
}

// SetOrigin 设置图片的原点
func (p *Picture) SetOrigin(x, y float32) error {
	if err := checkNilPicture(p); err != nil {
		return err
	}
	return resultToError(C.tvg_picture_set_origin(p.ptr, C.float(x), C.float(y)))
}

// GetOrigin 获取图片的原点
func (p *Picture) GetOrigin() (x, y float32, err error) {
	if err = checkNilPicture(p); err != nil {
		return
	}
	var cX, cY C.float
	res := C.tvg_picture_get_origin(p.ptr, &cX, &cY)
	err = resultToError(res)
	x = float32(cX)
	y = float32(cY)
	return
}

// GetPaint 通过ID获取子画笔对象
func (p *Picture) GetPaint(id uint32) *Paint {
	if p.ptr == nil {
		return nil
	}
	cPaint := C.tvg_picture_get_paint(p.ptr, C.uint32_t(id))
	if cPaint == nil {
		return nil
	}
	return &Paint{ptr: cPaint}
}

// LoadImage 从标准的 Go image.Image 加载图片
// 这是一个易用的接口，自动处理图片格式转换和像素数据提取
func (p *Picture) LoadImage(img image.Image) error {
	if err := checkNilPicture(p); err != nil {
		return err
	}
	if img == nil {
		return ErrInvalidArgument
	}

	// 获取图片边界
	bounds := img.Bounds()
	if bounds.Dx() == 0 || bounds.Dy() == 0 {
		return ErrInvalidSize
	}

	width := uint32(bounds.Dx())
	height := uint32(bounds.Dy())

	// 创建像素数据数组，使用 ABGR8888 格式（ThorVG 的默认格式）
	pixels := make([]uint32, width*height)

	// 遍历图片的每个像素
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// 获取像素颜色
			r, g, b, a := img.At(x, y).RGBA()

			// 转换为 8 位颜色值
			r8 := uint8(r >> 8)
			g8 := uint8(g >> 8)
			b8 := uint8(b >> 8)
			a8 := uint8(a >> 8)

			// 计算在像素数组中的索引
			idx := (y-bounds.Min.Y)*int(width) + (x - bounds.Min.X)

			// 转换为 ABGR8888 格式
			// ABGR8888: A(8) B(8) G(8) R(8)
			pixels[idx] = uint32(a8)<<24 | uint32(b8)<<16 | uint32(g8)<<8 | uint32(r8)
		}
	}

	// 使用 LoadRaw 方法加载转换后的像素数据
	return p.LoadRaw(pixels, width, height, ColorspaceABGR8888, true)
}
