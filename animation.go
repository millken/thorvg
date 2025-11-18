package thorvg

/*
#include <stdlib.h>
#include <string.h>
#include "thorvg_capi.h"
*/
import "C"

// NewAnimation 创建一个新的 Animation 对象
func NewAnimation() *Animation {
	paint := C.tvg_animation_new()
	if paint == nil {
		return nil
	}
	return &Animation{ptr: paint}
}

// SetFrame 设置动画的当前帧
func (a *Animation) SetFrame(frame float32) error {
	if err := checkNilAnimation(a); err != nil {
		return err
	}
	if err := checkPositiveValue(frame, "frame"); err != nil {
		return err
	}
	return resultToError(C.tvg_animation_set_frame(a.ptr, C.float(frame)))
}

// GetPicture 获取与动画关联的图片对象
func (a *Animation) GetPicture() *Picture {
	if a == nil || a.ptr == nil {
		return nil
	}
	c := C.tvg_animation_get_picture(a.ptr)
	if c == nil {
		return nil
	}
	return &Picture{Paint{ptr: c}}
}

// GetFrame 获取动画的当前帧
func (a *Animation) GetFrame() (float32, error) {
	if err := checkNilAnimation(a); err != nil {
		return 0, err
	}
	var cFrame C.float
	res := C.tvg_animation_get_frame(a.ptr, &cFrame)
	return float32(cFrame), resultToError(res)
}

// GetTotalFrame 获取动画的总帧数
func (a *Animation) GetTotalFrame() (float32, error) {
	if err := checkNilAnimation(a); err != nil {
		return 0, err
	}
	var cTotal C.float
	res := C.tvg_animation_get_total_frame(a.ptr, &cTotal)
	return float32(cTotal), resultToError(res)
}

// GetDuration 获取动画的持续时间（秒）
func (a *Animation) GetDuration() (float32, error) {
	if err := checkNilAnimation(a); err != nil {
		return 0, err
	}
	var cDuration C.float
	res := C.tvg_animation_get_duration(a.ptr, &cDuration)
	return float32(cDuration), resultToError(res)
}

// SetSegment 设置动画的播放段
func (a *Animation) SetSegment(begin, end float32) error {
	if err := checkNilAnimation(a); err != nil {
		return err
	}
	if err := checkPositiveValue(begin, "begin"); err != nil {
		return err
	}
	if err := checkPositiveValue(end, "end"); err != nil {
		return err
	}
	if begin > end {
		return ErrInvalidArgument
	}
	return resultToError(C.tvg_animation_set_segment(a.ptr, C.float(begin), C.float(end)))
}

// GetSegment 获取动画的播放段
func (a *Animation) GetSegment() (begin, end float32, err error) {
	if err = checkNilAnimation(a); err != nil {
		return
	}
	var cBegin, cEnd C.float
	res := C.tvg_animation_get_segment(a.ptr, &cBegin, &cEnd)
	err = resultToError(res)
	begin = float32(cBegin)
	end = float32(cEnd)
	return
}

// Release 释放 Animation 资源
func (a *Animation) Release() error {
	if a == nil || a.ptr == nil {
		return nil
	}
	res := C.tvg_animation_del(a.ptr)
	a.ptr = nil
	return resultToError(res)
}
