// Copyright 2013 The winapi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package winapi

type (
	HANDLE    uintptr
	HWND      HANDLE
	HMENU     HANDLE
	HMODULE   HANDLE
	HINSTANCE HANDLE
	HDC       HANDLE
	HRGN      HANDLE
	HBRUSH    HANDLE
	HICON     HANDLE
	HCURSOR   HANDLE
	HPEN      HANDLE
	HPALETTE  HANDLE
	HBITMAP   HANDLE
	WPARAM    uintptr
	LPARAM    uintptr
	UINT      uint32
	BOOL      int32
	DWORD     uint32
	LRESULT   int
	COLORREF  uint32
	LANGID    uint16
)

type Point struct {
	X, Y int32
}

type Rect struct {
	Left, Top, Right, Bottom int32
}
