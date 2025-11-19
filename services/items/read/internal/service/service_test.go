package service

import "testing"

func TestSanity(t *testing.T) {
	if 2+2 != 4 {
		t.Fatal("expected math to work")
	}
}
