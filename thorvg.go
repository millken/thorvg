package thorvg

/*
#cgo CFLAGS: -I${SRCDIR}/include
#cgo darwin LDFLAGS: -L${SRCDIR} -lc++
#cgo darwin,arm64 LDFLAGS: -lthorvg_darwin_arm64 -L/opt/homebrew/opt/libomp/lib -lomp
#cgo linux,amd64 LDFLAGS: -L${SRCDIR}/_libs/ -lthorvg_linux_amd64 -ldl -lm -lstdc++ -lomp -lGL

#include <stdlib.h>
#include <string.h>
#include "thorvg_capi.h"
*/
import "C"

// 错误处理辅助函数，将Tvg_Result转换为Go error
// resultToError 将Tvg_Result转换为Go error（向后兼容）
func resultToError(res C.Tvg_Result) error {
	result := resultToResult(res)
	if result.IsSuccess() {
		return nil
	}
	return result
}

// resultToResult 将Tvg_Result转换为Go Result类型
func resultToResult(res C.Tvg_Result) Result {
	switch byte(res) {
	case 0:
		return ResultSuccess
	case 1:
		return ResultInvalidArgument
	case 2:
		return ResultInsufficientCondition
	case 3:
		return ResultFailedAllocation
	case 4:
		return ResultMemoryCorruption
	case 5:
		return ResultNotSupported
	case 255:
		return ResultUnknown
	default:
		return ResultUnknown
	}
}
