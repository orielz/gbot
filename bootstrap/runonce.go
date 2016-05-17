package bootstrap

import (
	"os"
	"syscall"
	"unsafe"
)

var (
	modkernel32        = syscall.NewLazyDLL("kernel32.dll")
	procCreateMailslot = modkernel32.NewProc("CreateMailslotW")
)

// RunOnce insure only one instance
func RunOnce(guid string) {
	err := singleInstance(guid)
	if err != nil {
	}
}

func singleInstance(name string) error {
	ret, _, _ := procCreateMailslot.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(`\\.\mailslot\`+name))),
		0,
		0,
		0,
	)
	if int64(ret) == -1 {
		os.Exit(0)
	}
	return nil
}
