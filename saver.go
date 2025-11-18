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

// NewSaver 创建一个新的 Saver 对象
func NewSaver() *Saver {
	paint := C.tvg_saver_new()
	if paint == nil {
		return nil
	}
	return &Saver{ptr: paint}
}

// Save 将画布保存到文件
func (s *Saver) SavePaint(paint interface{}, path string, quality uint32) error {
	if err := checkNilSaver(s); err != nil {
		return err
	}
	if err := checkEmptyString(path); err != nil {
		return err
	}
	if err := checkRangeUint32(quality, 0, 100, "quality"); err != nil {
		return err
	}

	var p *Paint
	switch v := paint.(type) {
	case *Paint:
		p = v
	case *Shape:
		p = &v.Paint
	case *Scene:
		p = &v.Paint
	case *Picture:
		p = &v.Paint
	case *Animation:
		p = &v.GetPicture().Paint
	default:
		return ErrInvalidArgument
	}
	if p == nil || p.ptr == nil {
		return ErrInvalidArgument
	}

	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	return resultToError(C.tvg_saver_save_paint(s.ptr, p.ptr, cPath, C.uint32_t(quality)))
}

// SaveAnimation 将动画保存到文件
func (s *Saver) SaveAnimation(animation *Animation, path string, quality, fps uint32) error {
	if err := checkNilSaver(s); err != nil {
		return err
	}
	if err := checkNilAnimation(animation); err != nil {
		return err
	}
	if err := checkEmptyString(path); err != nil {
		return err
	}
	if err := checkRangeUint32(quality, 0, 100, "quality"); err != nil {
		return err
	}
	if err := checkPositiveValue(fps, "fps"); err != nil {
		return err
	}

	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	return resultToError(C.tvg_saver_save_animation(s.ptr, animation.ptr, cPath, C.uint32_t(quality), C.uint32_t(fps)))
}

// Sync 同步保存操作，确保保存完成
func (s *Saver) Sync() error {
	if err := checkNilSaver(s); err != nil {
		return err
	}
	return resultToError(C.tvg_saver_sync(s.ptr))
}

// Release 释放 Saver 资源
func (s *Saver) Release() error {
	if s == nil || s.ptr == nil {
		return nil
	}
	res := C.tvg_saver_del(s.ptr)
	s.ptr = nil
	return resultToError(res)
}
