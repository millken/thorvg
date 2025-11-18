package thorvg

/*
#include <stdlib.h>
#include <string.h>
#include "thorvg_capi.h"
*/
import "C"
import "unsafe"

// Version 获取 TVG 版本号
func Version() (major, minor, micro uint32, version string, res Result) {
	var cMajor, cMinor, cMicro C.uint32_t
	var cVersion *C.char
	r := C.tvg_engine_version(&cMajor, &cMinor, &cMicro, (**C.char)(unsafe.Pointer(&cVersion)))
	major = uint32(cMajor)
	minor = uint32(cMinor)
	micro = uint32(cMicro)
	if cVersion != nil {
		version = C.GoString(cVersion)
	}
	res = resultToResult(r)
	return
}

// Init 初始化 ThorVG 引擎
func Init(threads uint) error {
	return resultToError(C.tvg_engine_init(C.uint(threads)))
}

// Term 终止 ThorVG 引擎
func Term() error {
	return resultToError(C.tvg_engine_term())
}
