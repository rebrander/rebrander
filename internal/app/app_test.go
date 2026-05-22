package app

import "testing"

func TestOptionsValidate(t *testing.T) {
	opts := Options{Source: "https://github.com/a/a", Target: "https://github.com/b/b", Mode: "safe"}
	if err := opts.validate(); err != nil {
		t.Fatalf("expected valid options: %v", err)
	}
}

func TestOptionsValidateRejectsMode(t *testing.T) {
	opts := Options{Source: "s", Target: "t", Mode: "unsafe"}
	if err := opts.validate(); err == nil {
		t.Fatal("expected invalid mode error")
	}
}
