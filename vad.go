package webrtcvad

/*
#cgo CFLAGS: -Iwebrtc
#include "webrtc/common_audio/vad/include/webrtc_vad.h"
void rtc_FatalMessage(const char* file, int line, const char* msg) {
}
*/
import "C"

import (
	"errors"
	"runtime"
	"unsafe"
)

func New() (*VAD, error) {
	var inst *C.struct_WebRtcVadInst

	inst = C.WebRtcVad_Create()
	if inst == nil {
		return nil, errors.New("failed to create VAD")
	}

	vad := &VAD{inst}
	runtime.SetFinalizer(vad, free)

	ret := C.WebRtcVad_Init(inst)
	if ret != 0 {
		return nil, errors.New("default mode could not be set")
	}

	return vad, nil
}

func free(vad *VAD) {
	C.WebRtcVad_Free(vad.inst)
}

type VAD struct {
	inst *C.struct_WebRtcVadInst
}

func (v *VAD) SetMode(mode int) error {
	ret := C.WebRtcVad_set_mode(v.inst, C.int(mode))
	if ret != 0 {
		return errors.New("mode could not be set")
	}
	return nil
}

func (v *VAD) Process(fs int, audioFrame []byte) (activeVoice bool, err error) {
	if len(audioFrame)%2 != 0 {
		return false, errors.New("audio frames must be 16bit little endian unsigned integers")
	}

	audioFramePtr := (*C.int16_t)(unsafe.Pointer(&audioFrame[0]))
	frameLen := C.ulong(len(audioFrame) / 2)

	ret := C.WebRtcVad_Process(v.inst, C.int(fs), audioFramePtr, frameLen)
	switch ret {
	case 0:
		return false, nil
	case 1:
		return true, nil
	default:
		return false, errors.New("processing error")
	}
}

func (v *VAD) ValidRateAndFrameLength(rate int, frameLength int) bool {
	ret := C.WebRtcVad_ValidRateAndFrameLength(C.int(rate), C.ulong(frameLength))
	if ret < 0 {
		return false
	}
	return true
}
