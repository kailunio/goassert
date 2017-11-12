package assert

import "testing"

func TestIsNil(t *testing.T) {

	// testing point of string
	var s = "abc"
	var ps = &s
	if IsNil(ps) != false {
		t.Error("IsNil assert failed")
	}

	ps = nil
	if IsNil(ps) != true {
		t.Error("IsNil assert failed")
	}

	// testing slice
	var slice []int = nil
	if IsNil(slice) != true {
		t.Error("IsNil assert failed")
	}

	slice = make([]int, 10)
	if IsNil(slice) != false {
		t.Error("IsNil assert failed")
	}

	// testing map
	var mmap map[string]string = nil
	if IsNil(mmap) != true {
		t.Error("IsNil assert failed")
	}

	mmap = make(map[string]string, 10)
	if IsNil(mmap) != false {
		t.Error("IsNil assert failed")
	}

	// testing func
	var ffunc = func() { }
	if IsNil(ffunc) != false {
		t.Error("IsNil assert failed")
	}

	ffunc = nil
	if IsNil(ffunc) != true {
		t.Error("IsNil assert failed")
	}
}

func TestIsEquals(t *testing.T) {
	// equals nil
	if !IsEquals(nil, nil) {
		t.Error("IsEquals assert failed")
	}

	if IsEquals("", nil) {
		t.Error("IsEquals assert failed")
	}
}

