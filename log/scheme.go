package log

import (
	"github.com/ydb-platform/ydb-go-sdk/v3/trace"
)

func Scheme(l Logger, details trace.Details) (t trace.Scheme) {
	if details&trace.SchemeEvents == 0 {
		return
	}
	_ = l.WithName(`scheme`)
	return t
}
