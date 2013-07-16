// Copyright 2013 The winapi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package winapi

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
