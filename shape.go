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

// Shape 表示一个形状对象，用于绘制路径、填充和笔触
type Shape struct {
	Paint
}

// NewShape 创建一个新的 Shape 对象
func NewShape() *Shape {
	paint := C.tvg_shape_new()
	if paint == nil {
		return nil
	}
	return &Shape{Paint{ptr: paint}}
}

// Reset 重置形状路径属性，保留颜色、填充和笔触属性
func (s *Shape) Reset() error {
	if err := checkNilShape(s); err != nil {
		return err
	}
	return resultToError(C.tvg_shape_reset(s.ptr))
}

// MoveTo 设置子路径的初始点
func (s *Shape) MoveTo(x, y float32) error {
	if err := checkNilShape(s); err != nil {
		return err
	}
	return resultToError(C.tvg_shape_move_to(s.ptr, C.float(x), C.float(y)))
}

// LineTo 添加一条从当前点到指定点的直线
func (s *Shape) LineTo(x, y float32) error {
	if err := checkNilShape(s); err != nil {
		return err
	}
	return resultToError(C.tvg_shape_line_to(s.ptr, C.float(x), C.float(y)))
}

// CubicTo 添加一条三次贝塞尔曲线
func (s *Shape) CubicTo(cx1, cy1, cx2, cy2, x, y float32) error {
	if err := checkNilShape(s); err != nil {
		return err
	}
	return resultToError(C.tvg_shape_cubic_to(s.ptr, C.float(cx1), C.float(cy1), C.float(cx2), C.float(cy2), C.float(x), C.float(y)))
}

// Close 闭合当前子路径
func (s *Shape) Close() error {
	if err := checkNilShape(s); err != nil {
		return err
	}
	return resultToError(C.tvg_shape_close(s.ptr))
}

// AppendRect 向路径添加一个矩形，可以有圆角
func (s *Shape) AppendRect(x, y, w, h, rx, ry float32, cw bool) error {
	if err := checkNilShape(s); err != nil {
		return err
	}
	return resultToError(C.tvg_shape_append_rect(s.ptr, C.float(x), C.float(y), C.float(w), C.float(h), C.float(rx), C.float(ry), C.bool(cw)))
}

// AppendCircle 向路径添加一个椭圆
func (s *Shape) AppendCircle(cx, cy, rx, ry float32, cw bool) error {
	if err := checkNilShape(s); err != nil {
		return err
	}
	return resultToError(C.tvg_shape_append_circle(s.ptr, C.float(cx), C.float(cy), C.float(rx), C.float(ry), C.bool(cw)))
}

// AppendPath 向路径添加一个子路径
func (s *Shape) AppendPath(cmds []PathCommand, pts []Point) error {
	if err := checkNilShape(s); err != nil {
		return err
	}
	if err := checkEmptySlice(cmds); err != nil {
		return err
	}
	if err := checkEmptySlice(pts); err != nil {
		return err
	}

	cmdCnt := len(cmds)
	ptsCnt := len(pts)

	cCmds := make([]C.Tvg_Path_Command, cmdCnt)
	for i, cmd := range cmds {
		cCmds[i] = C.Tvg_Path_Command(cmd)
	}

	cPts := make([]C.Tvg_Point, ptsCnt)
	for i, pt := range pts {
		cPts[i] = *pt.Cptr()
	}

	return resultToError(C.tvg_shape_append_path(s.ptr, &cCmds[0], C.uint32_t(cmdCnt), &cPts[0], C.uint32_t(ptsCnt)))
}

// GetPath 获取形状的当前路径数据
func (s *Shape) GetPath() ([]PathCommand, []Point, error) {
	if err := checkNilShape(s); err != nil {
		return nil, nil, err
	}

	var cmds *C.Tvg_Path_Command
	var cmdsCnt C.uint32_t
	var pts *C.Tvg_Point
	var ptsCnt C.uint32_t

	res := C.tvg_shape_get_path(s.ptr, &cmds, &cmdsCnt, &pts, &ptsCnt)
	if res != C.TVG_RESULT_SUCCESS {
		return nil, nil, resultToError(res)
	}

	if cmdsCnt == 0 || ptsCnt == 0 {
		return nil, nil, nil
	}

	// 复制命令
	goCmds := make([]PathCommand, cmdsCnt)
	cCmdsSlice := (*[1 << 30]C.Tvg_Path_Command)(unsafe.Pointer(cmds))[:cmdsCnt:cmdsCnt]
	for i, cmd := range cCmdsSlice {
		goCmds[i] = PathCommand(cmd)
	}

	// 复制点
	goPts := make([]Point, ptsCnt)
	cPtsSlice := (*[1 << 30]C.Tvg_Point)(unsafe.Pointer(pts))[:ptsCnt:ptsCnt]
	for i, pt := range cPtsSlice {
		goPts[i] = Point{X: float32(pt.x), Y: float32(pt.y)}
	}

	return goCmds, goPts, nil
}

// SetStrokeWidth 设置笔触宽度
func (s *Shape) SetStrokeWidth(width float32) error {
	if err := checkNilShape(s); err != nil {
		return err
	}
	return resultToError(C.tvg_shape_set_stroke_width(s.ptr, C.float(width)))
}

// GetStrokeWidth 获取笔触宽度
func (s *Shape) GetStrokeWidth() (float32, error) {
	if err := checkNilShape(s); err != nil {
		return 0, err
	}
	var width C.float
	res := C.tvg_shape_get_stroke_width(s.ptr, &width)
	return float32(width), resultToError(res)
}

// SetStrokeColor 设置笔触颜色
func (s *Shape) SetStrokeColor(r, g, b, a uint8) error {
	if err := checkNilShape(s); err != nil {
		return err
	}
	return resultToError(C.tvg_shape_set_stroke_color(s.ptr, C.uint8_t(r), C.uint8_t(g), C.uint8_t(b), C.uint8_t(a)))
}

// GetStrokeColor 获取笔触颜色
func (s *Shape) GetStrokeColor() (r, g, b, a uint8, err error) {
	if err = checkNilShape(s); err != nil {
		return 0, 0, 0, 0, err
	}
	var cr, cg, cb, ca C.uint8_t
	res := C.tvg_shape_get_stroke_color(s.ptr, &cr, &cg, &cb, &ca)
	return uint8(cr), uint8(cg), uint8(cb), uint8(ca), resultToError(res)
}

// SetStrokeGradient 设置笔触渐变
func (s *Shape) SetStrokeGradient(grad *Gradient) error {
	if err := checkNilShape(s); err != nil {
		return err
	}
	if err := checkNilGradient(grad); err != nil {
		return err
	}
	return resultToError(C.tvg_shape_set_stroke_gradient(s.ptr, grad.ptr))
}

// GetStrokeGradient 获取笔触渐变
func (s *Shape) GetStrokeGradient() (*Gradient, error) {
	if err := checkNilShape(s); err != nil {
		return nil, err
	}
	var grad C.Tvg_Gradient
	res := C.tvg_shape_get_stroke_gradient(s.ptr, &grad)
	if res != C.TVG_RESULT_SUCCESS {
		return nil, resultToError(res)
	}
	return &Gradient{ptr: grad}, nil
}

// SetStrokeDash 设置笔触虚线模式
func (s *Shape) SetStrokeDash(dashPattern []float32, offset float32) error {
	if err := checkNilShape(s); err != nil {
		return err
	}
	if len(dashPattern) == 0 {
		// 重置虚线模式
		return resultToError(C.tvg_shape_set_stroke_dash(s.ptr, nil, 0, C.float(offset)))
	}

	cDash := make([]C.float, len(dashPattern))
	for i, d := range dashPattern {
		cDash[i] = C.float(d)
	}

	return resultToError(C.tvg_shape_set_stroke_dash(s.ptr, &cDash[0], C.uint32_t(len(dashPattern)), C.float(offset)))
}

// GetStrokeDash 获取笔触虚线模式
func (s *Shape) GetStrokeDash() ([]float32, float32, error) {
	if err := checkNilShape(s); err != nil {
		return nil, 0, err
	}

	var dashPattern *C.float
	var cnt C.uint32_t
	var offset C.float

	res := C.tvg_shape_get_stroke_dash(s.ptr, &dashPattern, &cnt, &offset)
	if res != C.TVG_RESULT_SUCCESS {
		return nil, 0, resultToError(res)
	}

	if cnt == 0 {
		return nil, float32(offset), nil
	}

	// 复制虚线模式
	dash := make([]float32, cnt)
	cDashSlice := (*[1 << 30]C.float)(unsafe.Pointer(dashPattern))[:cnt:cnt]
	for i, d := range cDashSlice {
		dash[i] = float32(d)
	}

	return dash, float32(offset), nil
}

// SetStrokeCap 设置笔触端点样式
func (s *Shape) SetStrokeCap(cap StrokeCap) error {
	if err := checkNilShape(s); err != nil {
		return err
	}
	return resultToError(C.tvg_shape_set_stroke_cap(s.ptr, C.Tvg_Stroke_Cap(cap)))
}

// GetStrokeCap 获取笔触端点样式
func (s *Shape) GetStrokeCap() (StrokeCap, error) {
	if err := checkNilShape(s); err != nil {
		return 0, err
	}
	var cap C.Tvg_Stroke_Cap
	res := C.tvg_shape_get_stroke_cap(s.ptr, &cap)
	return StrokeCap(cap), resultToError(res)
}

// SetStrokeJoin 设置笔触连接样式
func (s *Shape) SetStrokeJoin(join StrokeJoin) error {
	if err := checkNilShape(s); err != nil {
		return err
	}
	return resultToError(C.tvg_shape_set_stroke_join(s.ptr, C.Tvg_Stroke_Join(join)))
}

// GetStrokeJoin 获取笔触连接样式
func (s *Shape) GetStrokeJoin() (StrokeJoin, error) {
	if err := checkNilShape(s); err != nil {
		return 0, err
	}
	var join C.Tvg_Stroke_Join
	res := C.tvg_shape_get_stroke_join(s.ptr, &join)
	return StrokeJoin(join), resultToError(res)
}

// SetStrokeMiterlimit 设置笔触斜接限制
func (s *Shape) SetStrokeMiterlimit(miterlimit float32) error {
	if err := checkNilShape(s); err != nil {
		return err
	}
	return resultToError(C.tvg_shape_set_stroke_miterlimit(s.ptr, C.float(miterlimit)))
}

// GetStrokeMiterlimit 获取笔触斜接限制
func (s *Shape) GetStrokeMiterlimit() (float32, error) {
	if err := checkNilShape(s); err != nil {
		return 0, err
	}
	var miterlimit C.float
	res := C.tvg_shape_get_stroke_miterlimit(s.ptr, &miterlimit)
	return float32(miterlimit), resultToError(res)
}

// SetTrimpath 设置路径修剪
func (s *Shape) SetTrimpath(begin, end float32, simultaneous bool) error {
	if err := checkNilShape(s); err != nil {
		return err
	}
	return resultToError(C.tvg_shape_set_trimpath(s.ptr, C.float(begin), C.float(end), C.bool(simultaneous)))
}

// SetFillColor 设置填充颜色
func (s *Shape) SetFillColor(r, g, b, a uint8) error {
	if err := checkNilShape(s); err != nil {
		return err
	}
	return resultToError(C.tvg_shape_set_fill_color(s.ptr, C.uint8_t(r), C.uint8_t(g), C.uint8_t(b), C.uint8_t(a)))
}

// GetFillColor 获取填充颜色
func (s *Shape) GetFillColor() (r, g, b, a uint8, err error) {
	if err = checkNilShape(s); err != nil {
		return 0, 0, 0, 0, err
	}
	var cr, cg, cb, ca C.uint8_t
	res := C.tvg_shape_get_fill_color(s.ptr, &cr, &cg, &cb, &ca)
	return uint8(cr), uint8(cg), uint8(cb), uint8(ca), resultToError(res)
}

// SetFillRule 设置填充规则
func (s *Shape) SetFillRule(rule FillRule) error {
	if err := checkNilShape(s); err != nil {
		return err
	}
	return resultToError(C.tvg_shape_set_fill_rule(s.ptr, C.Tvg_Fill_Rule(rule)))
}

// GetFillRule 获取填充规则
func (s *Shape) GetFillRule() (FillRule, error) {
	if err := checkNilShape(s); err != nil {
		return 0, err
	}
	var rule C.Tvg_Fill_Rule
	res := C.tvg_shape_get_fill_rule(s.ptr, &rule)
	return FillRule(rule), resultToError(res)
}

// SetPaintOrder 设置笔触和填充的渲染顺序
func (s *Shape) SetPaintOrder(strokeFirst bool) error {
	if err := checkNilShape(s); err != nil {
		return err
	}
	return resultToError(C.tvg_shape_set_paint_order(s.ptr, C.bool(strokeFirst)))
}

// SetGradient 设置渐变填充
func (s *Shape) SetGradient(grad *Gradient) error {
	if err := checkNilShape(s); err != nil {
		return err
	}
	if err := checkNilGradient(grad); err != nil {
		return err
	}
	return resultToError(C.tvg_shape_set_gradient(s.ptr, grad.ptr))
}

// GetGradient 获取渐变填充
func (s *Shape) GetGradient() (*Gradient, error) {
	if err := checkNilShape(s); err != nil {
		return nil, err
	}
	var grad C.Tvg_Gradient
	res := C.tvg_shape_get_gradient(s.ptr, &grad)
	if res != C.TVG_RESULT_SUCCESS {
		return nil, resultToError(res)
	}
	return &Gradient{ptr: grad}, nil
}
