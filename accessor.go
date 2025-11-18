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

// NewAccessor 创建一个新的 Accessor 对象
func NewAccessor() *Accessor {
	paint := C.tvg_accessor_new()
	if paint == nil {
		return nil
	}
	return &Accessor{ptr: paint}
}

// Set 设置 paint 并遍历其后代，调用回调函数
// 注意：这个方法需要特殊的实现来处理Go函数指针到C函数指针的转换
// 暂时未实现
func (a *Accessor) Set(paint *Paint, callback func(*Paint) bool) error {
	if err := checkNilAccessor(a); err != nil {
		return err
	}
	if err := checkNilPaint(paint); err != nil {
		return err
	}
	if callback == nil {
		return ErrInvalidArgument
	}

	// TODO: 实现Go函数指针到C函数指针的转换
	// 这需要一个复杂的机制来注册回调函数
	return ErrNotImplemented
}

// GenerateID 从名称生成唯一ID
func (a *Accessor) GenerateID(name string) uint32 {
	if a == nil || a.ptr == nil {
		return 0
	}
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return uint32(C.tvg_accessor_generate_id(cName))
}

// Release 释放 Accessor 资源
func (a *Accessor) Release() error {
	if a == nil || a.ptr == nil {
		return nil
	}
	res := C.tvg_accessor_del(a.ptr)
	a.ptr = nil
	return resultToError(res)
}

// GenerateID 包级函数，从名称生成唯一ID
func GenerateID(name string) uint32 {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return uint32(C.tvg_accessor_generate_id(cName))
}
