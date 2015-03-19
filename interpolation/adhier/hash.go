package adhier

import (
	"reflect"
	"unsafe"
)

type hash struct {
	ni      int
	mapping map[string]bool
}

func newHash(ni uint) *hash {
	return &hash{
		ni:      int(ni),
		mapping: make(map[string]bool),
	}
}

func (h *hash) unique(indices []uint64) []uint64 {
	const (
		sizeOfUint64 = 8
	)

	ni := h.ni
	nn := len(indices) / ni
	nb := ni * sizeOfUint64

	var bytes []byte
	header := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	header.Cap = nb
	header.Len = nb

	offset := ((*reflect.SliceHeader)(unsafe.Pointer(&indices))).Data

	ns := 0

	for i, k := 0, 0; i < nn; i++ {
		header.Data = offset + uintptr(k*nb)
		key := string(bytes)

		if _, ok := h.mapping[key]; !ok {
			h.mapping[key] = true
			if k > ns {
				copy(indices[ns*ni:], indices[k*ni:])
				k = ns
			}
			ns++
		}

		k++
	}

	return indices[:ns*ni]
}
