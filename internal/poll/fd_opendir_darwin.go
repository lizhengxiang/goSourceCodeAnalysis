// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package poll

import (
	"syscall"
	_ "unsafe" // for go:linkname
)

// OpenDir returns a pointer to a DIR structure suitable for
// ReadDir. In case of an error, the name of the failed
// syscall is returned along with a syscall.Errno.

// OpenDir返回一个指向
//ReadDir的DIR结构的指针。如果发生错误，将返回失败的
//系统调用的名称以及syscall.Errno。

func (fd *FD) OpenDir() (uintptr, string, error) {
	// fdopendir(3) takes control of the file descriptor,
	// so use a dup.
	fd2, call, err := fd.Dup()
	if err != nil {
		return 0, call, err
	}
	dir, err := fdopendir(fd2)
	if err != nil {
		syscall.Close(fd2)
		return 0, "fdopendir", err
	}
	return dir, "", nil
}

// Implemented in syscall/syscall_darwin.go.
//go:linkname fdopendir syscall.fdopendir
func fdopendir(fd int) (dir uintptr, err error)
