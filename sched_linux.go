// Package sched contains functions for set/get process scheduling parameters
package sched

import (
	"syscall"
	"unsafe"
)

type Policy uintptr

const (
	Other Policy = 0
	FIFO  Policy = 1
	RR    Policy = 2
	Batch Policy = 3
	Idle  Policy = 5
)

func (p Policy) String() string {
	switch p {
	case Other:
		return "Other"
	case FIFO:
		return "FIFO"
	case RR:
		return "RR"
	case Batch:
		return "Batch"
	case Idle:
		return "Idle"
	}
	return "Unknown"
}

func (p Policy) MinPriority() int {
	r0, _, e := syscall.RawSyscall(
		syscall.SYS_SCHED_GET_PRIORITY_MIN,
		uintptr(p),
		0, 0,
	)
	if e != 0 {
		panic(e)
	}
	return int(r0)
}

func (p Policy) MaxPriority() int {
	r0, _, e := syscall.RawSyscall(
		syscall.SYS_SCHED_GET_PRIORITY_MAX,
		uintptr(p),
		0, 0,
	)
	if e != 0 {
		panic(e)
	}
	return int(r0)
}

type Param struct {
	Priority int
}

func SetPolicy(pid int, policy Policy, param *Param) error {
	_, _, e := syscall.RawSyscall(
		syscall.SYS_SCHED_SETSCHEDULER,
		uintptr(pid),
		uintptr(policy),
		uintptr(unsafe.Pointer(param)),
	)
	if e != 0 {
		return e
	}
	return nil
}

func GetPolicy(pid int) (Policy, error) {
	r0, _, e := syscall.RawSyscall(
		syscall.SYS_SCHED_GETSCHEDULER,
		uintptr(pid),
		0, 0,
	)
	if e != 0 {
		return 0, e
	}
	return Policy(r0), nil
}

func SetParam(pid int, param *Param) error {
	_, _, e := syscall.RawSyscall(
		syscall.SYS_SCHED_SETPARAM,
		uintptr(pid),
		uintptr(unsafe.Pointer(param)),
		0,
	)
	if e != 0 {
		return e
	}
	return nil

}

func GetParam(pid int, param *Param) error {
	_, _, e := syscall.RawSyscall(
		syscall.SYS_SCHED_GETPARAM,
		uintptr(pid),
		uintptr(unsafe.Pointer(param)),
		0,
	)
	if e != 0 {
		return e
	}
	return nil
}
