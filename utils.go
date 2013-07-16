// Copyright 2013 The winapi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package winapi

import (
	"errors"
	"fmt"
	"strconv"
	"unicode/utf16"
	"unsafe"
)

//GetLastError to error
func LastError(pfx string) error {
	v := GetLastError()
	var s string
	if pfx == "" {
		s = fmt.Sprintf("Last error:%d(0x%x)", v, v)
	} else {
		s = fmt.Sprintf("%s Last error:%d(0x%x)", pfx, v, v)
	}

	return errors.New(s)
}

func GoStringToPtr(v string) uintptr {
	if v == "" {
		return 0
	}

	u := utf16.Encode([]rune(v))
	u = append(u, 0)

	return uintptr(unsafe.Pointer(&u[0]))
}

func PtrToGoString(v uintptr) string {
	if v == 0 {
		return ""
	}

	vp := (*[1 << 29]uint16)(unsafe.Pointer(v))
	size := 0
	for ; vp[size] != 0; size++ {
	}

	return string(utf16.Decode(vp[:size]))
}

func BoolToPtr(v bool) (ret uintptr) {
	if v {
		ret = uintptr(1)
	} else {
		ret = uintptr(0)
	}

	return
}

func PtrToBool(v uintptr) (ret bool) {
	if int(v) > 0 {
		ret = true
	} else {
		ret = false
	}

	return ret
}

func allNumber(s string) bool {
	for _, v := range s {
		if !(v >= '0' && v <= '9') {
			return false
		}
	}

	return true
}

func resourceNameToPtr(name string) uintptr {
	number := allNumber(name)
	var id uintptr
	if number {
		idNumber, err := strconv.Atoi(name)
		if err != nil {
			id = GoStringToPtr(name)
		} else {
			id = uintptr(idNumber)
		}
	} else {
		id = GoStringToPtr(name)
	}

	return id
}

func ResourceIdToName(id int) string {
	return strconv.Itoa(id)
}
