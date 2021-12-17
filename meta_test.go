package core

import "testing"

func TestMetaData(t *testing.T) {
	if TraceID.String() != "TraceId" {
		t.Errorf(TraceID.String())
	}
}
