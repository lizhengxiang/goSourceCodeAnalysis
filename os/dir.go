// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

// Readdir reads the contents of the directory associated with file and
// returns a slice of up to n FileInfo values, as would be returned
// by Lstat, in directory order. Subsequent calls on the same file will yield
// further FileInfos.
//
// If n > 0, Readdir returns at most n FileInfo structures. In this case, if
// Readdir returns an empty slice, it will return a non-nil error
// explaining why. At the end of a directory, the error is io.EOF.
//
// If n <= 0, Readdir returns all the FileInfo from the directory in
// a single slice. In this case, if Readdir succeeds (reads all
// the way to the end of the directory), it returns the slice and a
// nil error. If it encounters an error before the end of the
// directory, Readdir returns the FileInfo read until that point
// and a non-nil error.

//Readdir 读取与文件关联的目录的内容，并以目录顺序返回最多 n 个 FileInfo 值的片段，Lstat 将返回该片段。随后对相同文件的调用将产生更多的 FileInfos 。
//如果 n> 0，Readdir 最多返回 n 个 FileInfo 结构。在这种情况下，如果 Readdir 返回一个空片段，它将返回一个非零错误来解释原因。在目录结尾处，错误是 io.EOF 。
//如果 n<=0，则 Readdir 从单个切片中的目录返回所有 FileInfo 。在这种情况下，如果 Readdir 成功（一直读取到目录的末尾），
//它将返回切片并返回一个零错误。如果在目录结束之前遇到错误，则 Readdir 将返回 FileInfo 读取直到该点，并出现非零错误。

func (f *File) Readdir(n int) ([]FileInfo, error) {
	if f == nil {
		return nil, ErrInvalid
	}
	return f.readdir(n)
}

// Readdirnames reads the contents of the directory associated with file
// and returns a slice of up to n names of files in the directory,
// in directory order. Subsequent calls on the same file will yield
// further names.
//
// If n > 0, Readdirnames returns at most n names. In this case, if
// Readdirnames returns an empty slice, it will return a non-nil error
// explaining why. At the end of a directory, the error is io.EOF.
//
// If n <= 0, Readdirnames returns all the names from the directory in
// a single slice. In this case, if Readdirnames succeeds (reads all
// the way to the end of the directory), it returns the slice and a
// nil error. If it encounters an error before the end of the
// directory, Readdirnames returns the names read until that point and
// a non-nil error.

//Readdirnames 读取并返回目录f中的一段名称。
//如果 n>0，Readdirnames 最多返回 n 个名称。在这种情况下，如果 Readdirnames 返回一个空片段，它将返回一个非零错误来解释原因。在目录结尾处，错误是 io.EOF 。
//如果n <= 0，则 Readdirnames 将返回单个切片中目录中的所有名称。在这种情况下，如果 Readdirnames 成功（一直读到目录的末尾），
//它将返回切片并返回一个零错误。如果在目录结束之前遇到错误，则 Readdirnames 将返回直到该点时读取的名称以及非零错误。

func (f *File) Readdirnames(n int) (names []string, err error) {
	if f == nil {
		return nil, ErrInvalid
	}
	return f.readdirnames(n)
}
