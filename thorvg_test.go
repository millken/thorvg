package thorvg

import (
	"testing"
)

// TestResultError tests the Result.Error() method
func TestResultError(t *testing.T) {
	tests := []struct {
		name     string
		result   Result
		hasError bool
	}{
		{"Success", ResultSuccess, false},
		{"InvalidArgument", ResultInvalidArgument, true},
		{"InsufficientCondition", ResultInsufficientCondition, true},
		{"FailedAllocation", ResultFailedAllocation, true},
		{"MemoryCorruption", ResultMemoryCorruption, true},
		{"NotSupported", ResultNotSupported, true},
		{"Unknown", ResultUnknown, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.result.Error()
			if tt.hasError && err == nil {
				t.Errorf("Expected error for %v, got nil", tt.result)
			}
			if !tt.hasError && err != nil {
				t.Errorf("Expected no error for %v, got %v", tt.result, err)
			}
		})
	}
}

// TestConstants tests that constants are defined
func TestConstants(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
	}{
		{"EngineTypeSW", EngineTypeSW},
		{"EngineTypeGL", EngineTypeGL},
		{"ColorSpaceARGB8888", ColorSpaceARGB8888},
		{"ColorSpaceABGR8888", ColorSpaceABGR8888},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value == nil {
				t.Errorf("Constant %s is nil", tt.name)
			}
		})
	}
}

// Note: Integration tests that require ThorVG C library to be installed
// should be in a separate file with build tags like:
// // +build integration
//
// This allows users to run basic tests without having ThorVG installed
