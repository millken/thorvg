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

// NewSwCanvas 创建一个新的软件渲染画布
//
// 参数:
//   - option: 引擎选项，用于配置渲染行为
//
// 返回:
//   - *Canvas: 新创建的画布对象，失败时返回 nil
func NewSwCanvas(option EngineOption) *Canvas {
	canvas := C.tvg_swcanvas_create(C.Tvg_Engine_Option(option))
	if canvas == nil {
		return nil
	}
	return &Canvas{ptr: canvas}
}

// NewGlCanvas 创建一个新的 OpenGL 渲染画布
//
// 返回:
//   - *Canvas: 新创建的画布对象，失败时返回 nil
func NewGlCanvas() *Canvas {
	canvas := C.tvg_glcanvas_create()
	if canvas == nil {
		return nil
	}
	return &Canvas{ptr: canvas}
}

// NewWgCanvas 创建一个新的 WebGPU 渲染画布
//
// 返回:
//   - *Canvas: 新创建的画布对象，失败时返回 nil
func NewWgCanvas() *Canvas {
	canvas := C.tvg_wgcanvas_create()
	if canvas == nil {
		return nil
	}
	return &Canvas{ptr: canvas}
}

// Destroy 销毁画布对象
//
// 返回:
//   - error: 销毁失败时返回错误
func (c *Canvas) Destroy() error {
	if err := checkNilCanvas(c); err != nil {
		return err
	}
	return resultToError(C.tvg_canvas_destroy(c.ptr))
}

// Push 将绘制元素添加到画布中
//
// 参数:
//   - paint: 要添加到画布的绘制对象
//
// 返回:
//   - error: 添加失败时返回错误
//
// 注意: 只有添加到画布中的绘制对象才会成为渲染目标
func (c *Canvas) Push(paint *Paint) error {
	if err := checkNilCanvas(c); err != nil {
		return err
	}
	if err := checkNilPaint(paint); err != nil {
		return err
	}
	return resultToError(C.tvg_canvas_push(c.ptr, paint.ptr))
}

// PushAt 在指定位置之前添加绘制元素到画布中
//
// 参数:
//   - target: 要添加的绘制对象
//   - at: 目标位置之前的现有绘制对象，如果为 nil 则添加到末尾
//
// 返回:
//   - error: 添加失败时返回错误
func (c *Canvas) PushAt(target *Paint, at *Paint) error {
	if err := checkNilCanvas(c); err != nil {
		return err
	}
	if err := checkNilPaint(target); err != nil {
		return err
	}
	var atPtr C.Tvg_Paint
	if at != nil {
		atPtr = at.ptr
	}
	return resultToError(C.tvg_canvas_push_at(c.ptr, target.ptr, atPtr))
}

// Remove 从画布中移除绘制元素
//
// 参数:
//   - paint: 要移除的绘制对象，如果为 nil 则移除所有绘制对象
//
// 返回:
//   - error: 移除失败时返回错误
func (c *Canvas) Remove(paint *Paint) error {
	if err := checkNilCanvas(c); err != nil {
		return err
	}
	var paintPtr C.Tvg_Paint
	if paint != nil {
		paintPtr = paint.ptr
	}
	return resultToError(C.tvg_canvas_remove(c.ptr, paintPtr))
}

// Update 更新画布中的修改过的绘制对象，为渲染做准备
//
// 返回:
//   - error: 更新失败时返回错误
//
// 注意: 只有修改过的绘制对象会被处理
func (c *Canvas) Update() error {
	if err := checkNilCanvas(c); err != nil {
		return err
	}
	return resultToError(C.tvg_canvas_update(c.ptr))
}

// Draw 渲染画布中的绘制对象
//
// 参数:
//   - clear: 如果为 true，在绘制前将目标缓冲区清零
//
// 返回:
//   - error: 绘制失败时返回错误
//
// 注意: 如果画布未更新，可能会隐式调用 Update()
func (c *Canvas) Draw(clear bool) error {
	if err := checkNilCanvas(c); err != nil {
		return err
	}
	return resultToError(C.tvg_canvas_draw(c.ptr, C.bool(clear)))
}

// Sync 确保绘制任务完成
//
// 返回:
//   - error: 同步失败时返回错误
//
// 注意: 画布渲染可能是异步的，必须在 Draw() 后调用 Sync() 确保完成
func (c *Canvas) Sync() error {
	if err := checkNilCanvas(c); err != nil {
		return err
	}
	return resultToError(C.tvg_canvas_sync(c.ptr))
}

// SetViewport 设置画布的绘制区域
//
// 参数:
//   - x, y: 矩形左上角坐标
//   - w, h: 矩形宽度和高度
//
// 返回:
//   - error: 设置失败时返回错误
//
// 注意: 只能在 Sync() 后更改视口
func (c *Canvas) SetViewport(x, y, w, h int32) error {
	if err := checkNilCanvas(c); err != nil {
		return err
	}
	return resultToError(C.tvg_canvas_set_viewport(c.ptr, C.int32_t(x), C.int32_t(y), C.int32_t(w), C.int32_t(h)))
}

// SwCanvasSetTarget 设置软件画布的目标缓冲区
//
// 参数:
//   - buffer: 像素缓冲区
//   - stride: 每行像素数
//   - w, h: 图像宽度和高度
//   - cs: 颜色空间
//
// 返回:
//   - error: 设置失败时返回错误
func (c *Canvas) SwCanvasSetTarget(buffer []uint32, stride, w, h uint32, cs Colorspace) error {
	if err := checkNilCanvas(c); err != nil {
		return err
	}
	if err := checkEmptySlice(buffer); err != nil {
		return err
	}
	if err := checkBufferSize(len(buffer), int(stride*h)); err != nil {
		return err
	}
	return resultToError(C.tvg_swcanvas_set_target(
		c.ptr,
		(*C.uint32_t)(unsafe.Pointer(&buffer[0])),
		C.uint32_t(stride),
		C.uint32_t(w),
		C.uint32_t(h),
		C.Tvg_Colorspace(cs),
	))
}

// GlCanvasSetTarget 设置 OpenGL 画布的目标
//
// 参数:
//   - context: OpenGL 上下文
//   - id: 帧缓冲区对象 ID
//   - w, h: 图像宽度和高度
//   - cs: 颜色空间
//
// 返回:
//   - error: 设置失败时返回错误
func (c *Canvas) GlCanvasSetTarget(context unsafe.Pointer, id int32, w, h uint32, cs Colorspace) error {
	if err := checkNilCanvas(c); err != nil {
		return err
	}
	if err := checkPositiveValue(w, "width"); err != nil {
		return err
	}
	if err := checkPositiveValue(h, "height"); err != nil {
		return err
	}
	return resultToError(C.tvg_glcanvas_set_target(
		c.ptr,
		context,
		C.int32_t(id),
		C.uint32_t(w),
		C.uint32_t(h),
		C.Tvg_Colorspace(cs),
	))
}

// WgCanvasSetTarget 设置 WebGPU 画布的目标
//
// 参数:
//   - device: WebGPU 设备
//   - instance: WebGPU 实例
//   - target: WebGPU 目标纹理
//   - w, h: 图像宽度和高度
//   - cs: 颜色空间
//   - canvasType: 画布类型
//
// 返回:
//   - error: 设置失败时返回错误
func (c *Canvas) WgCanvasSetTarget(device, instance, target unsafe.Pointer, w, h uint32, cs Colorspace, canvasType int) error {
	if err := checkNilCanvas(c); err != nil {
		return err
	}
	if err := checkPositiveValue(w, "width"); err != nil {
		return err
	}
	if err := checkPositiveValue(h, "height"); err != nil {
		return err
	}
	return resultToError(C.tvg_wgcanvas_set_target(
		c.ptr,
		device,
		instance,
		target,
		C.uint32_t(w),
		C.uint32_t(h),
		C.Tvg_Colorspace(cs),
		C.int(canvasType),
	))
}
