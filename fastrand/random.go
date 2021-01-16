package fastrand

import (
	_ "unsafe" // 必须导入
)

// Uint32 返回无锁化生成的 uint32 值
//go:linkname Uint32 runtime.fastrand
func Uint32() uint32

// Uint64 返回基于 Uint32 生成的 uint64 值
func Uint64() uint64 {
	return uint64(Uint32())<<32 | uint64(Uint32())
}

func Uint32N(n uint32) uint32 {
	if n&(n-1) == 0 {
		return Uint32() & (n - 1)
	}

	return Uint32() % n
}

func Uint64N(n uint64) uint64 {
	if n&(n-1) == 0 {
		return Uint64() & (n - 1)
	}

	return Uint64() % n
}
