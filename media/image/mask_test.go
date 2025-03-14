package image

import (
	"image"
	"testing"
)

func TestSet_ValidPoint_ValueSetCorrectly(t *testing.T) {
	rect := image.Rect(0, 0, 10, 10)
	bitMask := NewBitMask(rect)
	bitMask.Set(5, 5, true)
	if bit, ok := bitMask.Get(5, 5); !ok || !bit {
		t.Errorf("Expected value to be set to true, got false")
	}
}

func TestSetAndGetValueAtBoundary_CorrectBehavior(t *testing.T) {
	rect := image.Rect(0, 0, 10, 10)
	bitMask := NewBitMask(rect)
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
