package main

import "testing"

func TestHostname(t *testing.T) {
	result := hostname()
	if result == "" {
		t.Errorf("hostname failed: result is empty") // to indicate test failed
	} else {
		t.Logf("hostname success, got %v", result)
	}
}
