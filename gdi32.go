// Copyright 2013 The winapi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package winapi

import (
	"syscall"
	"unsafe"
)

var (
	modGdi32                  = syscall.NewLazyDLL("Gdi32.dll")
	procGetObjectW            = modGdi32.NewProc("GetObjectW")
	procMoveToEx              = modGdi32.NewProc("MoveToEx")
	procTextOutW              = modGdi32.NewProc("TextOutW")
	procPolyTextOutW          = modGdi32.NewProc("PolyTextOutW")
	procExtTextOutW           = modGdi32.NewProc("ExtTextOutW")
	procGetTextExtentPointW   = modGdi32.NewProc("GetTextExtentPointW")
	procGetTextExtentPoint32W = modGdi32.NewProc("GetTextExtentPoint32W")
	procCreatePolygonRgn      = modGdi32.NewProc("CreatePolygonRgn")
	procDPtoLP                = modGdi32.NewProc("DPtoLP")
	procLPtoDP                = modGdi32.NewProc("LPtoDP")
	procPolygon               = modGdi32.NewProc("Polygon")
	procPolyline              = modGdi32.NewProc("Polyline")
	procLineTo                = modGdi32.NewProc("LineTo")
	procPolyBezier            = modGdi32.NewProc("PolyBezier")
	procPolyBezierTo          = modGdi32.NewProc("PolyBezierTo")
	procPolylineTo            = modGdi32.NewProc("PolylineTo")
	procSetViewportExtEx      = modGdi32.NewProc("SetViewportExtEx")
	procSetViewportOrgEx      = modGdi32.NewProc("SetViewportOrgEx")
	procSetWindowExtEx        = modGdi32.NewProc("SetWindowExtEx")
	procSetWindowOrgEx        = modGdi32.NewProc("SetWindowOrgEx")
	procOffsetViewportOrgEx   = modGdi32.NewProc("OffsetViewportOrgEx")
	procOffsetWindowOrgEx     = modGdi32.NewProc("OffsetWindowOrgEx")
	procScaleViewportExtEx    = modGdi32.NewProc("ScaleViewportExtEx")
	procScaleWindowExtEx      = modGdi32.NewProc("ScaleWindowExtEx")
	procSetBitmapDimensionEx  = modGdi32.NewProc("SetBitmapDimensionEx")
	procSetBrushOrgEx         = modGdi32.NewProc("SetBrushOrgEx")
	procGdiFlush              = modGdi32.NewProc("GdiFlush")
)

func GetObject(h HANDLE, size int32, data uintptr) int32 {
	var ret uintptr
	ret, _, lastError = procGetObjectW.Call(uintptr(h), uintptr(size), data)

	return int32(ret)
}

func MoveToEx(hdc HDC, x int32, y int32, lppt *POINT) bool {
	var ret uintptr
	ret, _, lastError = procMoveToEx.Call(uintptr(hdc), uintptr(x), uintptr(y), uintptr(unsafe.Pointer(lppt)))

	return PtrToBool(ret)
}

func TextOut(hdc HDC, x int32, y int32, lpString string) bool {
	var ret uintptr
	ret, _, lastError = procTextOutW.Call(uintptr(hdc), uintptr(x), uintptr(y), StringToUintptr(lpString), uintptr(len(lpString)))

	return PtrToBool(ret)
}

func GetTextExtentPoint(hdc HDC, lpString string, lpsz *SIZE) bool {
	var ret uintptr
	ret, _, lastError = procGetTextExtentPointW.Call(uintptr(hdc), StringToUintptr(lpString), uintptr(len(lpString)), uintptr(unsafe.Pointer(lpsz)))

	return PtrToBool(ret)
}

func GetTextExtentPoint32(hdc HDC, lpString string, psizl *SIZE) bool {
	var ret uintptr
	ret, _, lastError = procGetTextExtentPoint32W.Call(uintptr(hdc), StringToUintptr(lpString), uintptr(len(lpString)), uintptr(unsafe.Pointer(psizl)))

	return PtrToBool(ret)
}

func CreatePolygonRgn(pts []POINT, iMode int32) HRGN {
	var ret uintptr
	ret, _, lastError = procCreatePolygonRgn.Call(uintptr(unsafe.Pointer(&pts[0])), uintptr(len(pts)), uintptr(iMode))

	return HRGN(ret)
}

func DPtoLP(hdc HDC, pts []POINT) bool {
	var ret uintptr
	ret, _, lastError = procDPtoLP.Call(uintptr(hdc), uintptr(unsafe.Pointer(&pts[0])), uintptr(len(pts)))

	return PtrToBool(ret)
}

func LPtoDP(hdc HDC, pts []POINT) bool {
	var ret uintptr
	ret, _, lastError = procLPtoDP.Call(uintptr(hdc), uintptr(unsafe.Pointer(&pts[0])), uintptr(len(pts)))

	return PtrToBool(ret)
}

func Polygon(hdc HDC, pts []POINT) bool {
	var ret uintptr
	ret, _, lastError = procPolygon.Call(uintptr(hdc), uintptr(unsafe.Pointer(&pts[0])), uintptr(len(pts)))

	return PtrToBool(ret)
}

func Polyline(hdc HDC, pts []POINT) bool {
	var ret uintptr
	ret, _, lastError = procPolyline.Call(uintptr(hdc), uintptr(unsafe.Pointer(&pts[0])), uintptr(len(pts)))

	return PtrToBool(ret)
}

func LineTo(hdc HDC, x int32, y int32) bool {
	var ret uintptr
	ret, _, lastError = procLineTo.Call(uintptr(hdc), uintptr(x), uintptr(y))

	return PtrToBool(ret)
}

func PolyBezier(hdc HDC, pts []POINT) bool {
	var ret uintptr
	ret, _, lastError = procPolyBezier.Call(uintptr(hdc), uintptr(unsafe.Pointer(&pts[0])), uintptr(len(pts)))

	return PtrToBool(ret)
}

func PolyBezierTo(hdc HDC, pts []POINT) bool {
	var ret uintptr
	ret, _, lastError = procPolyBezierTo.Call(uintptr(hdc), uintptr(unsafe.Pointer(&pts[0])), uintptr(len(pts)))

	return PtrToBool(ret)
}

func PolylineTo(hdc HDC, pts []POINT) bool {
	var ret uintptr
	ret, _, lastError = procPolylineTo.Call(uintptr(hdc), uintptr(unsafe.Pointer(&pts[0])), uintptr(len(pts)))

	return PtrToBool(ret)
}

func SetViewportExtEx(hdc HDC, x int32, y int32, lpsz *SIZE) bool {
	var ret uintptr
	ret, _, lastError = procSetViewportExtEx.Call(uintptr(hdc), uintptr(x), uintptr(y), uintptr(unsafe.Pointer(lpsz)))

	return PtrToBool(ret)
}

func SetViewportOrgEx(hdc HDC, x int32, y int32, lppt *POINT) bool {
	var ret uintptr
	ret, _, lastError = procSetViewportOrgEx.Call(uintptr(hdc), uintptr(x), uintptr(y), uintptr(unsafe.Pointer(lppt)))

	return PtrToBool(ret)
}

func SetWindowExtEx(hdc HDC, x int32, y int32, lpsz *SIZE) bool {
	var ret uintptr
	ret, _, lastError = procSetWindowExtEx.Call(uintptr(hdc), uintptr(x), uintptr(y), uintptr(unsafe.Pointer(lpsz)))

	return PtrToBool(ret)
}

func SetWindowOrgEx(hdc HDC, x int32, y int32, lppt *POINT) bool {
	var ret uintptr
	ret, _, lastError = procSetWindowOrgEx.Call(uintptr(hdc), uintptr(x), uintptr(y), uintptr(unsafe.Pointer(lppt)))

	return PtrToBool(ret)
}

func OffsetViewportOrgEx(hdc HDC, x int32, y int32, lppt *POINT) bool {
	var ret uintptr
	ret, _, lastError = procOffsetViewportOrgEx.Call(uintptr(hdc), uintptr(x), uintptr(y), uintptr(unsafe.Pointer(lppt)))

	return PtrToBool(ret)
}

func OffsetWindowOrgEx(hdc HDC, x int32, y int32, lppt *POINT) bool {
	var ret uintptr
	ret, _, lastError = procOffsetWindowOrgEx.Call(uintptr(hdc), uintptr(x), uintptr(y), uintptr(unsafe.Pointer(lppt)))

	return PtrToBool(ret)
}

func ScaleViewportExtEx(hdc HDC, xn int32, dx int32, yn int32, yd int32, lpsz *SIZE) bool {
	var ret uintptr
	ret, _, lastError = procScaleViewportExtEx.Call(uintptr(hdc), uintptr(xn), uintptr(dx), uintptr(yn), uintptr(yd), uintptr(unsafe.Pointer(lpsz)))

	return PtrToBool(ret)
}

func ScaleWindowExtEx(hdc HDC, xn int32, xd int32, yn int32, yd int32, lpsz *SIZE) bool {
	var ret uintptr
	ret, _, lastError = procScaleWindowExtEx.Call(uintptr(hdc), uintptr(xn), uintptr(xd), uintptr(yn), uintptr(yd), uintptr(unsafe.Pointer(lpsz)))

	return PtrToBool(ret)
}

func SetBitmapDimensionEx(hbm HBITMAP, w int32, h int32, lpsz *SIZE) bool {
	var ret uintptr
	ret, _, lastError = procSetBitmapDimensionEx.Call(uintptr(hbm), uintptr(w), uintptr(h), uintptr(unsafe.Pointer(lpsz)))

	return PtrToBool(ret)
}

func SetBrushOrgEx(hdc HDC, x int32, y int32, lppt *POINT) bool {
	var ret uintptr
	ret, _, lastError = procSetBrushOrgEx.Call(uintptr(hdc), uintptr(x), uintptr(y), uintptr(unsafe.Pointer(lppt)))

	return PtrToBool(ret)
}

func GdiFlush() bool {
	var ret uintptr
	ret, _, lastError = procGdiFlush.Call()

	return PtrToBool(ret)
}

const LF_FACESIZE = 32

// Font weight constants
const (
	FW_DONTCARE   = 0
	FW_THIN       = 100
	FW_EXTRALIGHT = 200
	FW_ULTRALIGHT = FW_EXTRALIGHT
	FW_LIGHT      = 300
	FW_NORMAL     = 400
	FW_REGULAR    = 400
	FW_MEDIUM     = 500
	FW_SEMIBOLD   = 600
	FW_DEMIBOLD   = FW_SEMIBOLD
	FW_BOLD       = 700
	FW_EXTRABOLD  = 800
	FW_ULTRABOLD  = FW_EXTRABOLD
	FW_HEAVY      = 900
	FW_BLACK      = FW_HEAVY
)

// Charset constants
const (
	ANSI_CHARSET        = 0
	DEFAULT_CHARSET     = 1
	SYMBOL_CHARSET      = 2
	SHIFTJIS_CHARSET    = 128
	HANGEUL_CHARSET     = 129
	HANGUL_CHARSET      = 129
	GB2312_CHARSET      = 134
	CHINESEBIG5_CHARSET = 136
	GREEK_CHARSET       = 161
	TURKISH_CHARSET     = 162
	HEBREW_CHARSET      = 177
	ARABIC_CHARSET      = 178
	BALTIC_CHARSET      = 186
	RUSSIAN_CHARSET     = 204
	THAI_CHARSET        = 222
	EASTEUROPE_CHARSET  = 238
	OEM_CHARSET         = 255
	JOHAB_CHARSET       = 130
	VIETNAMESE_CHARSET  = 163
	MAC_CHARSET         = 77
)

// Font output precision constants
const (
	OUT_DEFAULT_PRECIS   = 0
	OUT_STRING_PRECIS    = 1
	OUT_CHARACTER_PRECIS = 2
	OUT_STROKE_PRECIS    = 3
	OUT_TT_PRECIS        = 4
	OUT_DEVICE_PRECIS    = 5
	OUT_RASTER_PRECIS    = 6
	OUT_TT_ONLY_PRECIS   = 7
	OUT_OUTLINE_PRECIS   = 8
	OUT_PS_ONLY_PRECIS   = 10
)

// Font clipping precision constants
const (
	CLIP_DEFAULT_PRECIS   = 0
	CLIP_CHARACTER_PRECIS = 1
	CLIP_STROKE_PRECIS    = 2
	CLIP_MASK             = 15
	CLIP_LH_ANGLES        = 16
	CLIP_TT_ALWAYS        = 32
	CLIP_EMBEDDED         = 128
)

// Font output quality constants
const (
	DEFAULT_QUALITY        = 0
	DRAFT_QUALITY          = 1
	PROOF_QUALITY          = 2
	NONANTIALIASED_QUALITY = 3
	ANTIALIASED_QUALITY    = 4
	CLEARTYPE_QUALITY      = 5
)

// Font pitch constants
const (
	DEFAULT_PITCH  = 0
	FIXED_PITCH    = 1
	VARIABLE_PITCH = 2
)

// Font family constants
const (
	FF_DECORATIVE = 80
	FF_DONTCARE   = 0
	FF_MODERN     = 48
	FF_ROMAN      = 16
	FF_SCRIPT     = 64
	FF_SWISS      = 32
)

type LOGFONT struct {
	LfHeight         int32
	LfWidth          int32
	LfEscapement     int32
	LfOrientation    int32
	LfWeight         int32
	LfItalic         byte
	LfUnderline      byte
	LfStrikeOut      byte
	LfCharSet        byte
	LfOutPrecision   byte
	LfClipPrecision  byte
	LfQuality        byte
	LfPitchAndFamily byte
	LfFaceName       [LF_FACESIZE]uint16
}

func RGB(r, g, b byte) COLORREF {
	return COLORREF(r) | (COLORREF(g) << 8) | (COLORREF(b) << 16)
}

func (p COLORREF) GetRValue() byte {
	return byte(p)
}

func (p COLORREF) GetGValue() byte {
	return byte(p >> 8)
}

func (p COLORREF) GetBValue() byte {
	return byte(p >> 16)
}

type SIZE struct {
	Cx, Cy int32
}

type POINT struct {
	X, Y int32
}

type RECT struct {
	Left, Top, Right, Bottom int32
}
