package main

import "testing"

func TestHostname(t *testing.T) {
	result := hostname()
	if result == "" {
		t.Error() // to indicate test failed
	}
}
