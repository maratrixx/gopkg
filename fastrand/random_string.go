package fastrand

import "strings"

const (
	UpperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LowerCase = "abcdefghijklmnopqrstuvwxyz"
	Number    = "0123456789"
	AllCase   = UpperCase + LowerCase + Number
)

func RandString(size int) string {
	var builder strings.Builder
	builder.Grow(size)
	caseLen := len(AllCase)
	for i := 0; i < size; i++ {
		builder.WriteByte(AllCase[Uint32N(uint32(caseLen))])
	}
	return builder.String()
}
