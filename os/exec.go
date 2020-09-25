// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import (
	"internal/testlog"
	"runtime"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

// Process stores the information about a process created by StartProcess.
type Process struct {
	Pid    int
	handle uintptr      // handle is accessed atomically on Windows
	isdone uint32       // process has been successfully waited on, non zero if true //进程已成功等待，如果为true，则非零
	sigMu  sync.RWMutex // avoid race between wait and signal //避免等待和信号之间的竞争
}

func newProcess(pid int, handle uintptr) *Process {
	p := &Process{Pid: pid, handle: handle}
	runtime.SetFinalizer(p, (*Process).Release)
	return p
}

func (p *Process) setDone() {
	atomic.StoreUint32(&p.isdone, 1)
}

func (p *Process) done() bool {
	return atomic.LoadUint32(&p.isdone) > 0
}

// ProcAttr holds the attributes that will be applied to a new process
// started by StartProcess.
// ProcAttr保留将应用于由StartProcess启动的新流程的属性。
type ProcAttr struct {
	// If Dir is non-empty, the child changes into the directory before
	// creating the process.
	//如果Dir为非空，则子进程将在创建进程之前进入目录。
	Dir string
	// If Env is non-nil, it gives the environment variables for the
	// new process in the form returned by Environ.
	// If it is nil, the result of Environ will be used.
	//如果Env不为nil，则以Environ返回的形式为新进程提供环境变量。 如果为nil，则将使用Environ的结果。
	Env []string
	// Files specifies the open files inherited by the new process. The
	// first three entries correspond to standard input, standard output, and
	// standard error. An implementation may support additional entries,
	// depending on the underlying operating system. A nil entry corresponds
	// to that file being closed when the process starts.

	//文件指定新进程继承的打开文件。 前三个条目对应于标准输入，标准输出和标准错误。
	//一个实现可能支持其他条目，取决于基础操作系统。 nil条目对应于该进程开始时关闭的文件。
	Files []*File

	// Operating system-specific process creation attributes.
	// Note that setting this field means that your program
	// may not execute properly or even compile on some
	// operating systems.
	Sys *syscall.SysProcAttr
}

// A Signal represents an operating system signal.
// The usual underlying implementation is operating system-dependent:
// on Unix it is syscall.Signal.

// Signal表示操作系统信号。 通常的底层实现依赖于操作系统：在Unix上是syscall.Signal。

type Signal interface {
	String() string
	Signal() // to distinguish from other Stringers  //与其他纵梁相区别
}

// Getpid returns the process id of the caller.
func Getpid() int { return syscall.Getpid() }

// Getppid returns the process id of the caller's parent.
func Getppid() int { return syscall.Getppid() }

// FindProcess looks for a running process by its pid.
//
// The Process it returns can be used to obtain information
// about the underlying operating system process.
//
// On Unix systems, FindProcess always succeeds and returns a Process
// for the given pid, regardless of whether the process exists.
func FindProcess(pid int) (*Process, error) {
	return findProcess(pid)
}

// StartProcess starts a new process with the program, arguments and attributes
// specified by name, argv and attr. The argv slice will become os.Args in the
// new process, so it normally starts with the program name.
//
// If the calling goroutine has locked the operating system thread
// with runtime.LockOSThread and modified any inheritable OS-level
// thread state (for example, Linux or Plan 9 name spaces), the new
// process will inherit the caller's thread state.
//
// StartProcess is a low-level interface. The os/exec package provides
// higher-level interfaces.
//
// If there is an error, it will be of type *PathError.
func StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error) {
	testlog.Open(name)
	return startProcess(name, argv, attr)
}

// Release releases any resources associated with the Process p,
// rendering it unusable in the future.
// Release only needs to be called if Wait is not.

// Release释放与Process p相关的所有资源，//使其将来无法使用。 //只有在没有Wait的情况下才需要调用Release
func (p *Process) Release() error {
	return p.release()
}

// Kill causes the Process to exit immediately. Kill does not wait until
// the Process has actually exited. This only kills the Process itself,
// not any other processes it may have started.
func (p *Process) Kill() error {
	return p.kill()
}

// Wait waits for the Process to exit, and then returns a
// ProcessState describing its status and an error, if any.
// Wait releases any resources associated with the Process.
// On most operating systems, the Process must be a child
// of the current process or an error will be returned.

//等待等待Process退出，然后返回
//ProcessState描述其状态和错误（如果有）。
//等待释放与流程关联的所有资源。
//在大多数操作系统上，该Process必须是当前进程的子代
//否则将返回错误。

func (p *Process) Wait() (*ProcessState, error) {
	return p.wait()
}

// Signal sends a signal to the Process.
// Sending Interrupt on Windows is not implemented.
func (p *Process) Signal(sig Signal) error {
	return p.signal(sig)
}

// UserTime returns the user CPU time of the exited process and its children.
// UserTime返回已退出进程及其子进程的用户CPU时间。
func (p *ProcessState) UserTime() time.Duration {
	return p.userTime()
}

// SystemTime returns the system CPU time of the exited process and its children.
func (p *ProcessState) SystemTime() time.Duration {
	return p.systemTime()
}

// Exited reports whether the program has exited.
func (p *ProcessState) Exited() bool {
	return p.exited()
}

// Success reports whether the program exited successfully,
// such as with exit status 0 on Unix.
func (p *ProcessState) Success() bool {
	return p.success()
}

// Sys returns system-dependent exit information about
// the process. Convert it to the appropriate underlying
// type, such as syscall.WaitStatus on Unix, to access its contents.
func (p *ProcessState) Sys() interface{} {
	return p.sys()
}

// SysUsage returns system-dependent resource usage information about
// the exited process. Convert it to the appropriate underlying
// type, such as *syscall.Rusage on Unix, to access its contents.
// (On Unix, *syscall.Rusage matches struct rusage as defined in the
// getrusage(2) manual page.)
func (p *ProcessState) SysUsage() interface{} {
	return p.sysUsage()
}
