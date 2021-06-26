package cgoutil

import (
	"syscall"
	"unsafe"
)

func IntPtr(n int) uintptr {
	return uintptr(n)
}

// GOARCH=386;CGO_ENABLED=1

func StrPtr(s string) uintptr {
	//return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s)))
	return uintptr(unsafe.Pointer(syscall.StringBytePtr(s)))
}
