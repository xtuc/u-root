// Copyright 2015-2017 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kexec

// Syscall number for kexec_file_load(2).
const _SYS_KEXEC_FILE_LOAD = 320

// kexec_file_load(2) syscall flags.
const (
	_KEXEC_FILE_UNLOAD       = 0x1
	_KEXEC_FILE_ON_CRASH     = 0x2
	_KEXEC_FILE_NO_INITRAMFS = 0x4
)
