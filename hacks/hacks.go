// Hacks needed for generics.
// (Must be a separate package because the linkname hack doesn't work in go2go.)
package hacks // import "generic/hacks"

import "unsafe"

// RuntimeHash returns the hash of x used by the Go runtime's maps,
// or panics if key is unhashable.
func RuntimeHash(key interface{}, seed uintptr) uintptr {	
	type eface struct { t, v unsafe.Pointer }
	e := (*eface)(unsafe.Pointer(&key))
	if e.v == nil {
		return 0
	}
	return typehash(e.t, e.v, seed)
}

//go:linkname typehash reflect.typehash
func typehash(t, p unsafe.Pointer, h uintptr) uintptr
