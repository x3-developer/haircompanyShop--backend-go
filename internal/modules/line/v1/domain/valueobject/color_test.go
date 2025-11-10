package valueobject

import (
	"testing"
)

func TestNewColor_Valid(t *testing.T) {
	valid := []string{
		"#000000",
		"#FFFFFF",
		"#12AB9F",
		"#a1b2c3",
	}

	for _, c := range valid {
		color, err := NewColorVO(c)
		if err != nil {
			t.Fatalf("expected no error for %q, got %v", c, err)
		}
		if color.String() != c {
			t.Fatalf("expected String() to return %q, got %q", c, color.String())
		}
	}
}

func TestNewColor_Invalid(t *testing.T) {
	invalid := []string{
		"",
		"000000",
		"#12345",
		"#1234567",
		"#ZZZZZZ",
		"#12AGFF",
	}

	for _, c := range invalid {
		_, err := NewColorVO(c)
		if err == nil {
			t.Fatalf("expected error for %q, got nil", c)
		}
	}
}
