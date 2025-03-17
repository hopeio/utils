package image

import (
	"image"
	"testing"
)

func TestBitMask(t *testing.T) {
	rect := image.Rect(0, 0, 10, 10)
	bitMask := NewBitMask(rect)
	bitMask.Set(5, 5, true)
	if bit, ok := bitMask.Get(5, 5); !ok || !bit {
		t.Errorf("Expected value to be set to true, got false")
	}
	bitMask.Set(0, 0, true)
	if bit, ok := bitMask.Get(0, 0); !ok || !bit {
		t.Errorf("Expected boundary value to be set to true, got false")
	}
	bitMask.Set(9, 9, true)
	if bit, ok := bitMask.Get(9, 9); !ok || !bit {
		t.Errorf("Expected boundary value to be set to true, got false")
	}
}

func TestSetAndGetValue_OutOfBounds_NoPanic(t *testing.T) {
	rect := image.Rect(0, 0, 10, 10)
	bitMask := NewBitMask(rect)
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Expected no panic, got %v", r)
		}
	}()
	bitMask.Set(10, 10, true)
	bitMask.Get(10, 10)
}

func TestMask(t *testing.T) {
	rect := image.Rect(0, 0, 10, 10)
	mask := NewMask(rect)
	mask.Set(5, 5, 1)
	if bit, ok := mask.Get(5, 5); !ok || bit != 1 {
		t.Errorf("Expected value to be set to 1, got %d", bit)
	}
	mask.Set(0, 0, 2)
	if bit, ok := mask.Get(0, 0); !ok || bit != 2 {
		t.Errorf("Expected boundary value to be set to 2, got %d", bit)
	}
	mask.Set(9, 9, 3)
	if bit, ok := mask.Get(9, 9); !ok || bit != 3 {
		t.Errorf("Expected boundary value to be set to 3, got %d", bit)
	}
}
