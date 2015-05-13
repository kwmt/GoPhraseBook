// Created by cgo - DO NOT EDIT

package arc4random

import "unsafe"

import "syscall"

import _ "runtime/cgo"

type _ unsafe.Pointer

func _Cerrno(dst *error, x int) { *dst = syscall.Errno(x) }
type _Ctype_uint uint32

type _Ctype_u_int32_t _Ctype_uint

type _Ctype_void [0]byte

func _Cfunc_arc4random() _Ctype_u_int32_t
