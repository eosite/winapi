// Copyright 2013 The winapi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package winapi

import (
	"syscall"
	"unicode/utf16"
	"unsafe"
)

var (
	modKernel32         = syscall.NewLazyDLL("Kernel32.dll")
	procGetLastError    = modKernel32.NewProc("GetLastError")
	procGetLocaleInfo   = modKernel32.NewProc("GetLocaleInfoW")
	procGetModuleHandle = modKernel32.NewProc("GetModuleHandleW")
	procLoadString      = modUser32.NewProc("LoadStringW")
)

func GetLastError() uint {
	ret, _, _ := procGetLastError.Call()

	return uint(ret)
}

func GetLocaleInfo(lcid LCID, lctype LCTYPE) []uint16 {
	buf := make([]uint16, 256)
	ret, _, _ := procGetLocaleInfo.Call(uintptr(lcid), uintptr(lctype),
		uintptr(unsafe.Pointer(&buf[0])), uintptr(256))

	if ret > 0 {
		return buf[:ret]
	} else {
		return nil
	}
}

func GetModuleHandle(moduleName string) HMODULE {
	var param uintptr = 0
	if moduleName != "" {
		param = GoStringToPtr(moduleName)
	}

	ret, _, _ := procGetModuleHandle.Call(param)

	return HMODULE(ret)
}

func LoadString(inst HINSTANCE, id uint) (ret string) {
	text := make([]uint16, 1024)
	r, _, _ := procLoadString.Call(uintptr(inst), uintptr(id),
		uintptr(unsafe.Pointer(&text[0])), 1024)

	if int(r) <= 0 {
		ret = ""
	} else {
		ret = string(utf16.Decode(text[0:r]))
	}

	return
}

type (
	LCID   uint32
	LCTYPE uint32
)

// Predefined locale ids
const (
	LOCALE_CUSTOM_DEFAULT     LCID = 0x0c00
	LOCALE_CUSTOM_UI_DEFAULT  LCID = 0x1400
	LOCALE_CUSTOM_UNSPECIFIED LCID = 0x1000
	LOCALE_INVARIANT          LCID = 0x007f
	LOCALE_USER_DEFAULT       LCID = 0x0400
	LOCALE_SYSTEM_DEFAULT     LCID = 0x0800
)

// Predefined LCType ids
const (
	LOCALE_SISO3166CTRYNAME  LCTYPE = 0x5a
	LOCALE_SISO3166CTRYNAME2 LCTYPE = 0x68
	LOCALE_SISO639LANGNAME   LCTYPE = 0x59
	LOCALE_SISO639LANGNAME2  LCTYPE = 0x67
)
