// Copyright 2013 The winapi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package winapi

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
