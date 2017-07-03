// Copyright 2015-2017 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kexec

import (
	"fmt"
	"io/ioutil"
	"os"
	"syscall"
	"unsafe"
)

// FileLoad loads the given kernel as the new kernel with the given ramfs and
// cmdline.
func FileLoad(kernel, ramfs *os.File, cmdline string) error {
	var flags uintptr
	var ramfsfd uintptr
	if ramfs != nil {
		ramfsfd = ramfs.Fd()
	} else {
		flags |= _KEXEC_FILE_NO_INITRAMFS
	}

	cmdPtr, err := syscall.BytePtrFromString(cmdline)
	if err != nil {
		return fmt.Errorf("could not use cmdline %q: %v", cmdline, err)
	}

	if _, _, errno := syscall.Syscall6(
		_SYS_KEXEC_FILE_LOAD,
		kernel.Fd(),
		ramfsfd,
		uintptr(len(cmdline)),
		uintptr(unsafe.Pointer(cmdPtr)),
		flags,
		0); errno != 0 {
		return fmt.Errorf("sys_kexec(%d, %d, %s, %x) = %v", kernel.Fd(), ramfsfd, cmdline, flags, errno)
	}
	return nil
}

// Reboot executes a kernel previously loaded with FileInit.
func Reboot() error {
	if err := syscall.Reboot(syscall.LINUX_REBOOT_CMD_KEXEC); err != nil {
		return fmt.Errorf("sys_reboot(..., kexec) = %v", err)
	}
	return nil
}

func CurrentKernelCmdline() (string, error) {
	procCmdline, err := ioutil.ReadFile("/proc/cmdline")
	if err != nil {
		return "", err
	}
	return string(procCmdline), nil
}
