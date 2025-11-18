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

// NewLinearGradient 创建一个新的线性渐变对象
func NewLinearGradient() *Gradient {
	grad := C.tvg_linear_gradient_new()
	if grad == nil {
		return nil
	}
	return &Gradient{ptr: grad}
}

// NewRadialGradient 创建一个新的径向渐变对象
func NewRadialGradient() *Gradient {
	grad := C.tvg_radial_gradient_new()
	if grad == nil {
		return nil
	}
	return &Gradient{ptr: grad}
}

// SetLinear 设置线性渐变参数
func (g *Gradient) SetLinear(x1, y1, x2, y2 float32) error {
	if err := checkNilGradient(g); err != nil {
		return err
	}
	return resultToError(C.tvg_linear_gradient_set(g.ptr, C.float(x1), C.float(y1), C.float(x2), C.float(y2)))
}

// GetLinear 获取线性渐变参数
func (g *Gradient) GetLinear() (x1, y1, x2, y2 float32, err error) {
	if err = checkNilGradient(g); err != nil {
		return
	}
	var cx1, cy1, cx2, cy2 C.float
	res := C.tvg_linear_gradient_get(g.ptr, &cx1, &cy1, &cx2, &cy2)
	err = resultToError(res)
	x1, y1, x2, y2 = float32(cx1), float32(cy1), float32(cx2), float32(cy2)
	return
}

// SetRadial 设置径向渐变参数
func (g *Gradient) SetRadial(cx, cy, r, fx, fy, fr float32) error {
	if err := checkNilGradient(g); err != nil {
		return err
	}
	return resultToError(C.tvg_radial_gradient_set(g.ptr, C.float(cx), C.float(cy), C.float(r), C.float(fx), C.float(fy), C.float(fr)))
}

// GetRadial 获取径向渐变参数
func (g *Gradient) GetRadial() (cx, cy, r, fx, fy, fr float32, err error) {
	if err = checkNilGradient(g); err != nil {
		return
	}
	var ccx, ccy, cr, cfx, cfy, cfr C.float
	res := C.tvg_radial_gradient_get(g.ptr, &ccx, &ccy, &cr, &cfx, &cfy, &cfr)
	err = resultToError(res)
	cx, cy, r, fx, fy, fr = float32(ccx), float32(ccy), float32(cr), float32(cfx), float32(cfy), float32(cfr)
	return
}

// SetColorStops 设置渐变颜色点
func (g *Gradient) SetColorStops(colorStops []ColorStop) error {
	if err := checkNilGradient(g); err != nil {
		return err
	}
	if err := checkEmptySlice(colorStops); err != nil {
		return err
	}

	cStops := make([]C.Tvg_Color_Stop, len(colorStops))
	for i, stop := range colorStops {
		cStops[i] = *stop.Cptr()
	}

	return resultToError(C.tvg_gradient_set_color_stops(g.ptr, &cStops[0], C.uint32_t(len(colorStops))))
}

// GetColorStops 获取渐变颜色点
func (g *Gradient) GetColorStops() ([]ColorStop, error) {
	if err := checkNilGradient(g); err != nil {
		return nil, err
	}

	var cStops *C.Tvg_Color_Stop
	var cnt C.uint32_t

	res := C.tvg_gradient_get_color_stops(g.ptr, &cStops, &cnt)
	if res != C.TVG_RESULT_SUCCESS {
		return nil, resultToError(res)
	}

	if cnt == 0 {
		return nil, nil
	}

	// 复制颜色点
	stops := make([]ColorStop, cnt)
	cStopsSlice := (*[1 << 30]C.Tvg_Color_Stop)(unsafe.Pointer(cStops))[:cnt:cnt]
	for i, stop := range cStopsSlice {
		stops[i] = ColorStop{
			Offset: float32(stop.offset),
			R:      uint8(stop.r),
			G:      uint8(stop.g),
			B:      uint8(stop.b),
			A:      uint8(stop.a),
		}
	}

	return stops, nil
}

// SetSpread 设置渐变扩展模式
func (g *Gradient) SetSpread(spread FillSpread) error {
	if err := checkNilGradient(g); err != nil {
		return err
	}
	return resultToError(C.tvg_gradient_set_spread(g.ptr, C.Tvg_Stroke_Fill(spread)))
}

// GetSpread 获取渐变扩展模式
func (g *Gradient) GetSpread() (FillSpread, error) {
	if err := checkNilGradient(g); err != nil {
		return 0, err
	}
	var spread C.Tvg_Stroke_Fill
	res := C.tvg_gradient_get_spread(g.ptr, &spread)
	return FillSpread(spread), resultToError(res)
}

// SetTransform 设置渐变变换矩阵
func (g *Gradient) SetTransform(m *Matrix) error {
	if err := checkNilGradient(g); err != nil {
		return err
	}
	if err := checkNilMatrix(m); err != nil {
		return err
	}
	cMatrix := m.CPtr()
	return resultToError(C.tvg_gradient_set_transform(g.ptr, cMatrix))
}

// GetTransform 获取渐变变换矩阵
func (g *Gradient) GetTransform() (*Matrix, error) {
	if err := checkNilGradient(g); err != nil {
		return nil, err
	}
	var cMatrix C.Tvg_Matrix
	res := C.tvg_gradient_get_transform(g.ptr, &cMatrix)
	if res != C.TVG_RESULT_SUCCESS {
		return nil, resultToError(res)
	}
	return &Matrix{
		E11: float32(cMatrix.e11), E12: float32(cMatrix.e12), E13: float32(cMatrix.e13),
		E21: float32(cMatrix.e21), E22: float32(cMatrix.e22), E23: float32(cMatrix.e23),
		E31: float32(cMatrix.e31), E32: float32(cMatrix.e32), E33: float32(cMatrix.e33),
	}, nil
}

// GetType 获取渐变类型
func (g *Gradient) GetType() (Type, error) {
	if err := checkNilGradient(g); err != nil {
		return 0, err
	}
	var cType C.Tvg_Type
	res := C.tvg_gradient_get_type(g.ptr, &cType)
	return Type(cType), resultToError(res)
}

// Duplicate 复制渐变对象
func (g *Gradient) Duplicate() (*Gradient, error) {
	if err := checkNilGradient(g); err != nil {
		return nil, err
	}
	cGrad := C.tvg_gradient_duplicate(g.ptr)
	if cGrad == nil {
		return nil, ErrOperationFailed
	}
	return &Gradient{ptr: cGrad}, nil
}

// Delete 删除渐变对象
func (g *Gradient) Delete() error {
	if g.ptr == nil {
		return nil
	}
	res := C.tvg_gradient_del(g.ptr)
	g.ptr = nil
	return resultToError(res)
}
