package core

import (
	"context"
	"testing"
)

func TestMetaData(t *testing.T) {
	const right = "traceId"
	if TraceID.String() != right {
		t.Errorf(TraceID.String())
	}
}

func TestMetaData_Get(t *testing.T) {
	const v = "123456aaa"
	ctx := context.Background()
	ctx = TraceID.Set(ctx, v)
	if res := TraceID.Get(ctx); res != v {
		t.Errorf("%s!=%s", res, v)
	}
}

func TestMetaData_GetString(t *testing.T) {
	const v = "123456aaa"
	ctx := context.Background()
	ctx = TraceID.SetString(ctx, v)
	if res := TraceID.GetString(ctx); res != v {
		t.Errorf("%s!=%s", res, v)
	}
}
