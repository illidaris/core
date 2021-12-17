package core

import (
	"context"
	"testing"
)

func TestMetaData(t *testing.T) {
	const right = "TraceId"
	if TraceID.String() != right {
		t.Errorf(TraceID.String())
	}
}

func TestMetaData_CtxGet(t *testing.T) {
	const v = "123456aaa"
	ctx := context.Background()
	ctx = TraceID.CtxSet(ctx, v)
	if res := TraceID.CtxGet(ctx); res != v {
		t.Errorf("%s!=%s", res, v)
	}
}

func TestMetaData_CtxGetString(t *testing.T) {
	const v = "123456aaa"
	ctx := context.Background()
	ctx = TraceID.CtxSetString(ctx, v)
	if res := TraceID.CtxGetString(ctx); res != v {
		t.Errorf("%s!=%s", res, v)
	}
}
