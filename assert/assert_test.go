package assert

import (
	"testing"
)

type mtWorker struct {
	hasError bool
}

func (m *mtWorker) Errorf(format string, args ...interface{}) {
	m.hasError = true
}

func TestEquals(t *testing.T) {
	var tw *mtWorker
	tw = new(mtWorker)
	Equals(tw, 1, 1)
	if tw.hasError {
		t.Errorf("should not hasError")
	}

	tw = new(mtWorker)
	Equals(tw, "a", "a")
	if tw.hasError {
		t.Errorf("should not hasError")
	}

	tw = new(mtWorker)
	Equals(tw, 1, 2)
	if !tw.hasError {
		t.Errorf("should hasError")
	}

	tw = new(mtWorker)
	Equals(tw, "a", "b")
	if !tw.hasError {
		t.Errorf("should hasError")
	}

	tw = new(mtWorker)
	Equals(tw, "a", 1)
	if !tw.hasError {
		t.Errorf("should hasError")
	}
}

func TestNotEquals(t *testing.T) {
	var tw *mtWorker
	tw = new(mtWorker)
	NotEquals(tw, 1, 1)
	if !tw.hasError {
		t.Errorf("should not hasError")
	}

	tw = new(mtWorker)
	NotEquals(tw, "a", "a")
	if !tw.hasError {
		t.Errorf("should not hasError")
	}

	tw = new(mtWorker)
	NotEquals(tw, 1, 2)
	if tw.hasError {
		t.Errorf("should hasError")
	}

	tw = new(mtWorker)
	NotEquals(tw, "a", "b")
	if tw.hasError {
		t.Errorf("should hasError")
	}

	tw = new(mtWorker)
	NotEquals(tw, "a", 1)
	if tw.hasError {
		t.Errorf("should hasError")
	}
}

func TestNil(t *testing.T) {
	var tw *mtWorker
	tw = new(mtWorker)
	Nil(tw, 1)
	if !tw.hasError {
		t.Errorf("should  hasError")
	}

	tw = new(mtWorker)
	Nil(tw, nil)
	if tw.hasError {
		t.Errorf("should not hasError")
	}
}

func TestNotNil(t *testing.T) {
	var tw *mtWorker
	tw = new(mtWorker)
	NotNil(tw, 1)
	if tw.hasError {
		t.Errorf("should not hasError")
	}

	tw = new(mtWorker)
	NotNil(tw, nil)
	if !tw.hasError {
		t.Errorf("should hasError")
	}
}
