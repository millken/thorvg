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

// LoadFont 从文件加载可缩放字体数据
//
// ThorVG 使用指定的路径作为键高效缓存加载的数据。
// 这意味着再次加载同一文件不会导致重复操作；
// 相反，ThorVG 将重用之前加载的字体数据。
//
// 参数:
//   - path: 字体文件的绝对路径
//
// 返回:
//   - error: 如果加载失败返回错误
func LoadFont(path string) error {
	if err := checkEmptyString(path); err != nil {
		return err
	}
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	return resultToError(C.tvg_font_load(cPath))
}

// LoadFontData 从内存块加载可缩放字体数据
//
// ThorVG 使用指定的名称作为键高效缓存加载的字体数据。
// 这意味着再次加载相同的字体不会导致重复操作。
// 相反，ThorVG 将重用之前加载的字体数据。
//
// 参数:
//   - name: 字体存储和访问的名称
//   - data: 字体数据内容的指针
//   - mimetype: 字体数据的 mimetype 或扩展名
//   - copy: 如果为 true，则将数据复制到引擎本地缓冲区，否则不复制
//
// 返回:
//   - error: 如果加载失败返回错误
//
// 注意: 如果 copy 为 true，用户负责释放数据内存。
func LoadFontData(name string, data []byte, mimetype string, copy bool) error {
	if err := checkEmptyString(name); err != nil {
		return err
	}
	if err := checkEmptySlice(data); err != nil {
		return err
	}
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cData := (*C.char)(C.CBytes(data))
	defer C.free(unsafe.Pointer(cData))
	cMimetype := C.CString(mimetype)
	defer C.free(unsafe.Pointer(cMimetype))
	return resultToError(C.tvg_font_load_data(cName, cData, C.uint32_t(len(data)), cMimetype, C.bool(copy)))
}

// UnloadFont 卸载指定的可缩放字体数据
//
// 此函数用于释放与已加载字体文件关联的资源。
//
// 参数:
//   - path: 已加载字体文件的路径
//
// 返回:
//   - error: 如果卸载失败返回错误
//
// 注意: 如果字体数据当前正在使用，它不会立即卸载。
func UnloadFont(path string) error {
	if err := checkEmptyString(path); err != nil {
		return err
	}
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	return resultToError(C.tvg_font_unload(cPath))
}
