package thorvg

/*
#include <stdlib.h>
#include <string.h>
#include "thorvg_capi.h"
*/
import "C"
import "unsafe"

type EngineOption byte

const (
	EngineOptionNone EngineOption = iota
	EngineOptionDefault
	EngineOptionSmartRender //启用自动部分（智能）渲染优化
)

// Result 表示 ThorVG 操作的结果
type Result interface {
	error
	IsSuccess() bool
	Code() byte
}

// result 是 Result 的内部实现
type result struct {
	code byte
}

// 预定义的结果常量
var (
	ResultSuccess               Result = &result{0}   // 正确执行
	ResultInvalidArgument       Result = &result{1}   // 参数错误
	ResultInsufficientCondition Result = &result{2}   // 条件不足
	ResultFailedAllocation      Result = &result{3}   // 内存分配失败
	ResultMemoryCorruption      Result = &result{4}   // 内存损坏
	ResultNotSupported          Result = &result{5}   // 不支持
	ResultUnknown               Result = &result{255} // 其他未知情况
)

// Error 实现 error 接口
func (r *result) Error() string {
	switch r.code {
	case 0:
		return "success"
	case 1:
		return "invalid argument"
	case 2:
		return "insufficient condition"
	case 3:
		return "failed allocation"
	case 4:
		return "memory corruption"
	case 5:
		return "not supported"
	case 255:
		return "unknown error"
	default:
		return "unknown result"
	}
}

// IsSuccess 检查结果是否成功
func (r *result) IsSuccess() bool {
	return r.code == 0
}

// Code 返回结果代码
func (r *result) Code() byte {
	return r.code
}

// IsError 检查结果是否表示错误
func IsError(r Result) bool {
	return !r.IsSuccess()
}

type Colorspace byte

const (
	ColorspaceABGR8888   Colorspace = 0
	ColorspaceARGB8888   Colorspace = 1
	ColorspaceABGR8888S  Colorspace = 2
	ColorspaceARGB8888S  Colorspace = 3
	ColorspaceGrayscale8 Colorspace = 4
	ColorspaceUnknown    Colorspace = 255
)

type MaskMethod byte

const (
	MaskNone       MaskMethod = 0  // 无遮罩
	MaskAlpha      MaskMethod = 1  // Alpha遮罩
	MaskInvAlpha   MaskMethod = 2  // 反Alpha遮罩
	MaskLuma       MaskMethod = 3  // 灰度遮罩
	MaskInvLuma    MaskMethod = 4  // 反灰度遮罩
	MaskAdd        MaskMethod = 5  // 叠加
	MaskSubtract   MaskMethod = 6  // 相减
	MaskIntersect  MaskMethod = 7  // 交集
	MaskDifference MaskMethod = 8  // 差异
	MaskLighten    MaskMethod = 9  // 取最大透明度
	MaskDarken     MaskMethod = 10 // 取最小透明度
)

type Type byte

const (
	TypeUndefined      Type = 0  // 未定义类型
	TypeShape          Type = 1  // 形状类型
	TypeScene          Type = 2  // 场景类型
	TypePicture        Type = 3  // 图片类型
	TypeText           Type = 4  // 文本类型
	TypeLinearGradient Type = 10 // 线性渐变类型
	TypeRadialGradient Type = 11 // 放射渐变类型
)

type BlendMethod byte

const (
	BlendNormal     BlendMethod = 0  // 普通混合（默认）
	BlendMultiply   BlendMethod = 1  // 正片叠底
	BlendScreen     BlendMethod = 2  // 滤色
	BlendOverlay    BlendMethod = 3  // 叠加
	BlendDarken     BlendMethod = 4  // 变暗
	BlendLighten    BlendMethod = 5  // 变亮
	BlendColorDodge BlendMethod = 6  // 颜色减淡
	BlendColorBurn  BlendMethod = 7  // 颜色加深
	BlendHardLight  BlendMethod = 8  // 强光
	BlendSoftLight  BlendMethod = 9  // 柔光
	BlendDifference BlendMethod = 10 // 差值
	BlendExclusion  BlendMethod = 11 // 排除
	BlendHue        BlendMethod = 12 // 色相（保留，未支持）
	BlendSaturation BlendMethod = 13 // 饱和度（保留，未支持）
	BlendColor      BlendMethod = 14 // 颜色（保留，未支持）
	BlendLuminosity BlendMethod = 15 // 明度（保留，未支持）
	BlendAdd        BlendMethod = 16 // 线性加深
)

type PathCommand uint8

const (
	PathClose   PathCommand = 0 // 结束当前子路径并闭合
	PathMoveTo  PathCommand = 1 // 移动到新起点
	PathLineTo  PathCommand = 2 // 画直线到指定点
	PathCubicTo PathCommand = 3 // 画三次贝塞尔曲线到指定点
)

type FillRule byte

const (
	FillRuleNonZero FillRule = 0 // 非零环绕规则
	FillRuleEvenOdd FillRule = 1 // 奇偶规则
)

type StrokeCap byte

const (
	StrokeCapButt   StrokeCap = 0 // 笔触结束于端点
	StrokeCapRound  StrokeCap = 1 // 笔触以半圆结束
	StrokeCapSquare StrokeCap = 2 // 笔触以方形结束
)

type StrokeJoin byte

const (
	StrokeJoinMiter StrokeJoin = 0 // 尖角连接
	StrokeJoinRound StrokeJoin = 1 // 圆角连接
	StrokeJoinBevel StrokeJoin = 2 // 斜角连接
)

type FillSpread byte

const (
	FillSpreadPad     FillSpread = 0 // 剩余区域用最近的渐变端点色填充
	FillSpreadReflect FillSpread = 1 // 渐变区域外反射填充
	FillSpreadRepeat  FillSpread = 2 // 渐变区域外重复填充
)

type SceneEffect byte

const (
	SceneEffectClearAll     SceneEffect = 0 // 重置所有场景效果
	SceneEffectGaussianBlur SceneEffect = 1 // 高斯模糊效果
	SceneEffectDropShadow   SceneEffect = 2 // 阴影效果
	SceneEffectFill         SceneEffect = 3 // 填充效果
	SceneEffectTint         SceneEffect = 4 // 色调效果
	SceneEffectTritone      SceneEffect = 5 // 三色调效果
)

type TextWrap byte

const (
	TextWrapNone        TextWrap = 0 // 不换行
	TextWrapCharacter   TextWrap = 1 // 按字符换行
	TextWrapWord        TextWrap = 2 // 按单词换行
	TextWrapSmart       TextWrap = 3 // 智能换行
	TextWrapEllipsis    TextWrap = 4 // 省略号截断
	TextWrapHyphenation TextWrap = 5 // 连字符换行（保留）
)

type Matrix struct {
	E11, E12, E13 float32
	E21, E22, E23 float32
	E31, E32, E33 float32
}

func (m Matrix) CPtr() *C.Tvg_Matrix {
	return &C.Tvg_Matrix{
		e11: C.float(m.E11),
		e12: C.float(m.E12),
		e13: C.float(m.E13),
		e21: C.float(m.E21),
		e22: C.float(m.E22),
		e23: C.float(m.E23),
		e31: C.float(m.E31),
		e32: C.float(m.E32),
		e33: C.float(m.E33),
	}
}

// ColorStop 表示渐变中的一个颜色点
type ColorStop struct {
	Offset float32 // 颜色在渐变中的相对位置
	R      uint8   // 红色通道 [0,255]
	G      uint8   // 绿色通道 [0,255]
	B      uint8   // 蓝色通道 [0,255]
	A      uint8   // alpha通道 [0,255]
}

// Cptr 返回对应的 *C.Tvg_Color_Stop 指针
func (c *ColorStop) Cptr() *C.Tvg_Color_Stop {
	return (*C.Tvg_Color_Stop)(unsafe.Pointer(c))
}

// Point 表示二维空间中的点
type Point struct {
	X float32
	Y float32
}

// Cptr 返回对应的 *C.Tvg_Point 指针
func (p *Point) Cptr() *C.Tvg_Point {
	return (*C.Tvg_Point)(unsafe.Pointer(p))
}

// Gradient 表示渐变对象
type Gradient struct {
	ptr C.Tvg_Gradient
}

// Scene 表示场景对象，用于组合多个画布
type Scene struct {
	Paint
}

// Saver 表示保存器对象，用于将画布保存到文件
type Saver struct {
	ptr C.Tvg_Saver
}

// Animation 表示动画对象，用于控制动画播放
type Animation struct {
	ptr C.Tvg_Animation
}

// Accessor 表示访问器对象，用于遍历场景树
type Accessor struct {
	ptr C.Tvg_Accessor
}

// Text 表示文本对象，用于渲染文本
type Text struct {
	Paint
}

// Picture 表示图片对象，用于加载和显示图片
type Picture struct {
	Paint
}

// Canvas 表示画布对象，用于渲染绘制对象
type Canvas struct {
	ptr C.Tvg_Canvas
}

func toCStr(s string) uintptr {
	cstr := make([]byte, len(s)+1)
	copy(cstr, s)
	return uintptr(unsafe.Pointer(&cstr[0]))
}
