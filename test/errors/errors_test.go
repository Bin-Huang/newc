package errors

import (
	"testing"
)

func TestNewForbidden(t *testing.T) {
	value := NewForbidden("msg")
	if value.Status != 403 {
		t.Errorf("NewForbidden should calling init method")
	}
}
