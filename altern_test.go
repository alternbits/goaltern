package altern

import "testing"

func TestIsStatusSuccess(t *testing.T) {
	status, err := IsStatusSuccess()
	if err != nil {
		t.Fatalf("Failed to get status: %v", err)
	}

	if !status {
		t.Error("Expected status to be success but got failure")
	}
}
