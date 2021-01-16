package hack

import (
	"reflect"
	"unsafe"
)

func ToSlice(s string) (b []byte) {
	sheader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bheader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bheader.Data = sheader.Data
	bheader.Len = sheader.Len
	bheader.Cap = sheader.Len
	return
}

func ToString(b []byte) (s string) {
	if len(b) == 0 {
		return ""
	}

	sheader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bheader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sheader.Data = bheader.Data
	sheader.Len = bheader.Len
	return
}
