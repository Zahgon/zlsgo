//go:build go1.18
// +build go1.18

package zarray

import (
	"math/bits"
	"unsafe"
)

const (
	// hash input allowed sizes
	byteSize = 1 << iota
	wordSize
	dwordSize
	qwordSize
	owordSize
)

const (
	prime1 uint64 = 11400714785074694791
	prime2 uint64 = 14029467366897019727
	prime3 uint64 = 1609587929392839161
	prime4 uint64 = 9650029242287828579
	prime5 uint64 = 2870177450012600261
)

var prime1v = prime1

func u64(b []byte) uint64 { _ = "STUB: not implemented"; return 0 }
func u32(b []byte) uint32 { _ = "STUB: not implemented"; return 0 }

func round(acc, input uint64) uint64 { _ = "STUB: not implemented"; return 0 }

func mergeRound(acc, val uint64) uint64 { _ = "STUB: not implemented"; return 0 }

func rol1(x uint64) uint64  { _ = "STUB: not implemented"; return 0 }
func rol7(x uint64) uint64  { _ = "STUB: not implemented"; return 0 }
func rol11(x uint64) uint64 { _ = "STUB: not implemented"; return 0 }
func rol12(x uint64) uint64 { _ = "STUB: not implemented"; return 0 }
func rol18(x uint64) uint64 { _ = "STUB: not implemented"; return 0 }
func rol23(x uint64) uint64 { _ = "STUB: not implemented"; return 0 }
func rol27(x uint64) uint64 { _ = "STUB: not implemented"; return 0 }
func rol31(x uint64) uint64 { _ = "STUB: not implemented"; return 0 }

var (
	byteHasher = func(key uint8) uintptr {
		h := prime5 + 1
		h ^= uint64(key) * prime5
		h = bits.RotateLeft64(h, 11) * prime1
		h ^= h >> 33
		h *= prime2
		h ^= h >> 29
		h *= prime3
		h ^= h >> 32
		return uintptr(h)
	}

	wordHasher = func(key uint16) uintptr {
		h := prime5 + 2
		h ^= (uint64(key) & 0xff) * prime5
		h = bits.RotateLeft64(h, 11) * prime1
		h ^= ((uint64(key) >> 8) & 0xff) * prime5
		h = bits.RotateLeft64(h, 11) * prime1
		h ^= h >> 33
		h *= prime2
		h ^= h >> 29
		h *= prime3
		h ^= h >> 32
		return uintptr(h)
	}

	dwordHasher = func(key uint32) uintptr {
		h := prime5 + 4
		h ^= uint64(key) * prime1
		h = bits.RotateLeft64(h, 23)*prime2 + prime3
		h ^= h >> 33
		h *= prime2
		h ^= h >> 29
		h *= prime3
		h ^= h >> 32
		return uintptr(h)
	}

	float32Hasher = func(key float32) uintptr {
		h := prime5 + 4
		h ^= uint64(*(*uint32)(unsafe.Pointer(&key))) * prime1
		h = bits.RotateLeft64(h, 23)*prime2 + prime3
		h ^= h >> 33
		h *= prime2
		h ^= h >> 29
		h *= prime3
		h ^= h >> 32
		return uintptr(h)
	}

	qwordHasher = func(key uint64) uintptr {
		k1 := key * prime2
		k1 = bits.RotateLeft64(k1, 31)
		k1 *= prime1
		h := (prime5 + 8) ^ k1
		h = bits.RotateLeft64(h, 27)*prime1 + prime4
		h ^= h >> 33
		h *= prime2
		h ^= h >> 29
		h *= prime3
		h ^= h >> 32
		return uintptr(h)
	}

	float64Hasher = func(key float64) uintptr {
		k1 := *(*uint64)(unsafe.Pointer(&key)) * prime2
		k1 = bits.RotateLeft64(k1, 31)
		k1 *= prime1
		h := (prime5 + 8) ^ k1
		h = bits.RotateLeft64(h, 27)*prime1 + prime4
		h ^= h >> 33
		h *= prime2
		h ^= h >> 29
		h *= prime3
		h ^= h >> 32
		return uintptr(h)
	}

	complex64Hasher = func(key complex64) uintptr {
		k1 := *(*uint64)(unsafe.Pointer(&key)) * prime2
		k1 = bits.RotateLeft64(k1, 31)
		k1 *= prime1
		h := (prime5 + 8) ^ k1
		h = bits.RotateLeft64(h, 27)*prime1 + prime4
		h ^= h >> 33
		h *= prime2
		h ^= h >> 29
		h *= prime3
		h ^= h >> 32
		return uintptr(h)
	}
)

func (m *Maper[K, V]) setDefaultHasher() {
	_ = "STUB: not implemented"
	// default hash functions
	return
}

// word hasher

// dword hasher

// qword hasher

// byte hasher

// word hasher

// dword hasher

// custom float32 dword hasher

// qword hasher

// custom float64 qword hasher

// custom complex64 qword hasher

// oword hasher, key size -> 16 bytes
