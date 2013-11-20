// Copyright 2013 The winapi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package winapi

import (
	"strconv"
	"syscall"
	"unsafe"
)

func StringToUintptr(v string) uintptr {
	if v == "" {
		return 0
	}

	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(v)))
}

func UintptrToString(v uintptr) string {
	if v == 0 {
		return ""
	}

	return syscall.UTF16ToString((*[1 << 29]uint16)(unsafe.Pointer(v))[0:])
}

func UTF16PtrToString(v *uint16) string {
	return UintptrToString(uintptr(unsafe.Pointer(v)))
}

func PtrToBool(v uintptr) (ret bool) {
	if int(v) > 0 {
		ret = true
	} else {
		ret = false
	}

	return ret
}

func BoolToPtr(v bool) (ret uintptr) {
	if v {
		ret = 1
	} else {
		ret = 0
	}

	return
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
			id = StringToUintptr(name)
		} else {
			id = uintptr(idNumber)
		}
	} else {
		id = StringToUintptr(name)
	}

	return id
}

func ResourceIdToName(id int) string {
	return strconv.Itoa(id)
}
