package thorvg

/*
#include <stdlib.h>
#include <string.h>
#include "thorvg_capi.h"
*/
import "C"

type Paint struct {
	ptr C.Tvg_Paint
}

// 释放Paint资源
func (p *Paint) Release() error {
	if p.ptr == nil {
		return nil
	}
	res := C.tvg_paint_rel(p.ptr)
	p.ptr = nil
	return resultToError(res)
}

// 增加引用计数
func (p *Paint) Ref() uint16 {
	if p.ptr == nil {
		return 0
	}
	return uint16(C.tvg_paint_ref(p.ptr))
}

// 减少引用计数，如果free为true且计数为0则删除
func (p *Paint) Unref(free bool) uint16 {
	if p.ptr == nil {
		return 0
	}
	cFree := C.bool(free)
	return uint16(C.tvg_paint_unref(p.ptr, cFree))
}

// 获取当前引用计数
func (p *Paint) GetRef() uint16 {
	if p.ptr == nil {
		return 0
	}
	return uint16(C.tvg_paint_get_ref(p.ptr))
}

// 设置可见性
func (p *Paint) SetVisible(visible bool) error {
	if err := checkNilPaint(p); err != nil {
		return err
	}
	return resultToError(C.tvg_paint_set_visible(p.ptr, C.bool(visible)))
}

// 获取可见性
func (p *Paint) GetVisible() (bool, error) {
	if err := checkNilPaint(p); err != nil {
		return false, err
	}
	res := C.tvg_paint_get_visible(p.ptr)
	return bool(res), nil
}

// 缩放
func (p *Paint) Scale(factor float32) error {
	if err := checkNilPaint(p); err != nil {
		return err
	}
	return resultToError(C.tvg_paint_scale(p.ptr, C.float(factor)))
}

// 旋转
func (p *Paint) Rotate(degree float32) error {
	if err := checkNilPaint(p); err != nil {
		return err
	}
	return resultToError(C.tvg_paint_rotate(p.ptr, C.float(degree)))
}

// 平移
func (p *Paint) Translate(x, y float32) error {
	if err := checkNilPaint(p); err != nil {
		return err
	}
	return resultToError(C.tvg_paint_translate(p.ptr, C.float(x), C.float(y)))
}

// 设置变换矩阵
func (p *Paint) SetTransform(m *Matrix) error {
	if err := checkNilPaint(p); err != nil {
		return err
	}
	if err := checkNilMatrix(m); err != nil {
		return err
	}
	cMatrix := C.Tvg_Matrix{
		e11: C.float(m.E11), e12: C.float(m.E12), e13: C.float(m.E13),
		e21: C.float(m.E21), e22: C.float(m.E22), e23: C.float(m.E23),
		e31: C.float(m.E31), e32: C.float(m.E32), e33: C.float(m.E33),
	}
	return resultToError(C.tvg_paint_set_transform(p.ptr, &cMatrix))
}

// 获取变换矩阵
func (p *Paint) GetTransform() (*Matrix, error) {
	if err := checkNilPaint(p); err != nil {
		return nil, err
	}
	var cMatrix C.Tvg_Matrix
	res := C.tvg_paint_get_transform(p.ptr, &cMatrix)
	if res != C.TVG_RESULT_SUCCESS {
		return nil, resultToError(res)
	}
	return &Matrix{
		E11: float32(cMatrix.e11), E12: float32(cMatrix.e12), E13: float32(cMatrix.e13),
		E21: float32(cMatrix.e21), E22: float32(cMatrix.e22), E23: float32(cMatrix.e23),
		E31: float32(cMatrix.e31), E32: float32(cMatrix.e32), E33: float32(cMatrix.e33),
	}, nil
}

// 设置不透明度
func (p *Paint) SetOpacity(opacity uint8) error {
	if err := checkNilPaint(p); err != nil {
		return err
	}
	return resultToError(C.tvg_paint_set_opacity(p.ptr, C.uint8_t(opacity)))
}

// 获取不透明度
func (p *Paint) GetOpacity() (uint8, error) {
	if err := checkNilPaint(p); err != nil {
		return 0, err
	}
	var cOpacity C.uint8_t
	res := C.tvg_paint_get_opacity(p.ptr, &cOpacity)
	return uint8(cOpacity), resultToError(res)
}

// 复制Paint
func (p *Paint) Duplicate() (*Paint, error) {
	if err := checkNilPaint(p); err != nil {
		return nil, err
	}
	cPaint := C.tvg_paint_duplicate(p.ptr)
	if cPaint == nil {
		return nil, ErrOperationFailed
	}
	return &Paint{ptr: cPaint}, nil
}

// 检查区域是否与填充区域相交
func (p *Paint) Intersects(x, y, w, h int32) (bool, error) {
	if err := checkNilPaint(p); err != nil {
		return false, err
	}
	return bool(C.tvg_paint_intersects(p.ptr, C.int32_t(x), C.int32_t(y), C.int32_t(w), C.int32_t(h))), nil
}

// 获取轴对齐包围盒
func (p *Paint) GetAABB() (x, y, w, h float32, err error) {
	if err = checkNilPaint(p); err != nil {
		return
	}
	var cx, cy, cw, ch C.float
	res := C.tvg_paint_get_aabb(p.ptr, &cx, &cy, &cw, &ch)
	err = resultToError(res)
	x, y, w, h = float32(cx), float32(cy), float32(cw), float32(ch)
	return
}

// 获取对象导向包围盒
func (p *Paint) GetOBB() ([4]Point, error) {
	if err := checkNilPaint(p); err != nil {
		return [4]Point{}, err
	}
	var cPoints [4]C.Tvg_Point
	res := C.tvg_paint_get_obb(p.ptr, &cPoints[0])
	if res != C.TVG_RESULT_SUCCESS {
		return [4]Point{}, resultToError(res)
	}
	var points [4]Point
	for i := 0; i < 4; i++ {
		points[i] = Point{X: float32(cPoints[i].x), Y: float32(cPoints[i].y)}
	}
	return points, nil
}

// 设置遮罩方法
func (p *Paint) SetMaskMethod(target *Paint, method MaskMethod) error {
	if err := checkNilPaint(p); err != nil {
		return err
	}
	if err := checkNilPaint(target); err != nil {
		return err
	}
	return resultToError(C.tvg_paint_set_mask_method(p.ptr, target.ptr, C.Tvg_Mask_Method(method)))
}

// 获取遮罩方法
func (p *Paint) GetMaskMethod(target *Paint) (MaskMethod, error) {
	if err := checkNilPaint(p); err != nil {
		return 0, err
	}
	if err := checkNilPaint(target); err != nil {
		return 0, err
	}
	var cMethod C.Tvg_Mask_Method
	res := C.tvg_paint_get_mask_method(p.ptr, target.ptr, &cMethod)
	return MaskMethod(cMethod), resultToError(res)
}

// 设置剪裁
func (p *Paint) SetClip(clipper *Paint) error {
	if err := checkNilPaint(p); err != nil {
		return err
	}
	if err := checkNilPaint(clipper); err != nil {
		return err
	}
	return resultToError(C.tvg_paint_set_clip(p.ptr, clipper.ptr))
}

// 获取剪裁
func (p *Paint) GetClip() (*Paint, error) {
	if err := checkNilPaint(p); err != nil {
		return nil, err
	}
	cClipper := C.tvg_paint_get_clip(p.ptr)
	if cClipper == nil {
		return nil, nil
	}
	return &Paint{ptr: cClipper}, nil
}

// 获取父Paint
func (p *Paint) GetParent() (*Paint, error) {
	if err := checkNilPaint(p); err != nil {
		return nil, err
	}
	cParent := C.tvg_paint_get_parent(p.ptr)
	if cParent == nil {
		return nil, nil
	}
	return &Paint{ptr: cParent}, nil
}

// 获取类型
func (p *Paint) GetType() (Type, error) {
	if err := checkNilPaint(p); err != nil {
		return 0, err
	}
	var cType C.Tvg_Type
	res := C.tvg_paint_get_type(p.ptr, &cType)
	return Type(cType), resultToError(res)
}

// 设置混合方法
func (p *Paint) SetBlendMethod(method BlendMethod) error {
	if err := checkNilPaint(p); err != nil {
		return err
	}
	return resultToError(C.tvg_paint_set_blend_method(p.ptr, C.Tvg_Blend_Method(method)))
}
