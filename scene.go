package thorvg

/*
#include <stdlib.h>
#include <string.h>
#include "thorvg_capi.h"
*/
import "C"

// NewScene 创建一个新的 Scene 对象
func NewScene() *Scene {
	paint := C.tvg_scene_new()
	if paint == nil {
		return nil
	}
	return &Scene{Paint{ptr: paint}}
}

// Push 将画布添加到场景中
func (s *Scene) Push(paint interface{}) error {
	if err := checkNilScene(s); err != nil {
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
	default:
		return ErrInvalidArgument
	}
	if p == nil || p.ptr == nil {
		return ErrInvalidArgument
	}
	return resultToError(C.tvg_scene_push(s.ptr, p.ptr))
}

// PushAt 将画布添加到场景中的指定位置
func (s *Scene) PushAt(target, at interface{}) error {
	if err := checkNilScene(s); err != nil {
		return err
	}
	var t, a *Paint
	switch v := target.(type) {
	case *Paint:
		t = v
	case *Shape:
		t = &v.Paint
	case *Scene:
		t = &v.Paint
	default:
		return ErrInvalidArgument
	}
	switch v := at.(type) {
	case *Paint:
		a = v
	case *Shape:
		a = &v.Paint
	case *Scene:
		a = &v.Paint
	case nil:
		a = nil
	default:
		return ErrInvalidArgument
	}
	if t == nil || t.ptr == nil {
		return ErrInvalidArgument
	}
	var cat C.Tvg_Paint
	if a != nil {
		cat = a.ptr
	}
	return resultToError(C.tvg_scene_push_at(s.ptr, t.ptr, cat))
}

// Remove 从场景中移除画布
func (s *Scene) Remove(paint interface{}) error {
	if err := checkNilScene(s); err != nil {
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
	case nil:
		p = nil
	default:
		return ErrInvalidArgument
	}
	var cpaint C.Tvg_Paint
	if p != nil {
		cpaint = p.ptr
	}
	return resultToError(C.tvg_scene_remove(s.ptr, cpaint))
}

// ResetEffects 重置所有场景效果
func (s *Scene) ResetEffects() error {
	if err := checkNilScene(s); err != nil {
		return err
	}
	return resultToError(C.tvg_scene_reset_effects(s.ptr))
}

// PushEffectGaussianBlur 添加高斯模糊效果
func (s *Scene) PushEffectGaussianBlur(sigma float64, direction, border, quality int) error {
	if err := checkNilScene(s); err != nil {
		return err
	}
	if err := checkPositiveValue(sigma, "sigma"); err != nil {
		return err
	}
	return resultToError(C.tvg_scene_push_effect_gaussian_blur(s.ptr, C.double(sigma), C.int(direction), C.int(border), C.int(quality)))
}

// PushEffectDropShadow 添加阴影效果
func (s *Scene) PushEffectDropShadow(r, g, b, a uint8, angle, distance, sigma float64, quality int) error {
	if err := checkNilScene(s); err != nil {
		return err
	}
	if err := checkPositiveValue(sigma, "sigma"); err != nil {
		return err
	}
	return resultToError(C.tvg_scene_push_effect_drop_shadow(s.ptr, C.int(r), C.int(g), C.int(b), C.int(a), C.double(angle), C.double(distance), C.double(sigma), C.int(quality)))
}

// PushEffectFill 添加填充效果
func (s *Scene) PushEffectFill(r, g, b, a uint8) error {
	if err := checkNilScene(s); err != nil {
		return err
	}
	return resultToError(C.tvg_scene_push_effect_fill(s.ptr, C.int(r), C.int(g), C.int(b), C.int(a)))
}

// PushEffectTint 添加色调效果
func (s *Scene) PushEffectTint(blackR, blackG, blackB, whiteR, whiteG, whiteB uint8, intensity float64) error {
	if err := checkNilScene(s); err != nil {
		return err
	}
	return resultToError(C.tvg_scene_push_effect_tint(s.ptr, C.int(blackR), C.int(blackG), C.int(blackB), C.int(whiteR), C.int(whiteG), C.int(whiteB), C.double(intensity)))
}

// PushEffectTritone 添加三色调效果
func (s *Scene) PushEffectTritone(shadowR, shadowG, shadowB, midtoneR, midtoneG, midtoneB, highlightR, highlightG, highlightB uint8, blend uint8) error {
	if err := checkNilScene(s); err != nil {
		return err
	}
	return resultToError(C.tvg_scene_push_effect_tritone(s.ptr, C.int(shadowR), C.int(shadowG), C.int(shadowB), C.int(midtoneR), C.int(midtoneG), C.int(midtoneB), C.int(highlightR), C.int(highlightG), C.int(highlightB), C.int(blend)))
}
