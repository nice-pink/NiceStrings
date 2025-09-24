package datasize

import (
	"testing"
)

func TestDataSize_String(t *testing.T) {
	tests := []struct {
		name     string
		value    int64
		unit     string
		expected string
	}{
		{
			name:     "Bytes",
			value:    1024,
			unit:     Bytes,
			expected: "1024Bytes",
		},
		{
			name:     "KB",
			value:    5,
			unit:     KB,
			expected: "5KB",
		},
		{
			name:     "MB",
			value:    10,
			unit:     MB,
			expected: "10MB",
		},
		{
			name:     "GB",
			value:    2,
			unit:     GB,
			expected: "2GB",
		},
		{
			name:     "TB",
			value:    1,
			unit:     TB,
			expected: "1TB",
		},
		{
			name:     "PB",
			value:    3,
			unit:     PB,
			expected: "3PB",
		},
		{
			name:     "Zero value",
			value:    0,
			unit:     MB,
			expected: "0MB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &DataSize{
				Value: tt.value,
				Unit:  tt.unit,
			}
			result := ds.String()
			if result != tt.expected {
				t.Errorf("String() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestDataSize_ToBytes(t *testing.T) {
	tests := []struct {
		name     string
		value    int64
		unit     string
		expected int64
	}{
		{
			name:     "Bytes to bytes",
			value:    1024,
			unit:     Bytes,
			expected: 1024,
		},
		{
			name:     "KB to bytes",
			value:    1,
			unit:     KB,
			expected: 1024,
		},
		{
			name:     "MB to bytes",
			value:    1,
			unit:     MB,
			expected: 1024 * 1024,
		},
		{
			name:     "GB to bytes",
			value:    1,
			unit:     GB,
			expected: 1024 * 1024 * 1024,
		},
		{
			name:     "TB to bytes",
			value:    1,
			unit:     TB,
			expected: 1024 * 1024 * 1024 * 1024,
		},
		{
			name:     "PB to bytes",
			value:    1,
			unit:     PB,
			expected: 1024 * 1024 * 1024 * 1024 * 1024,
		},
		{
			name:     "Multiple KB to bytes",
			value:    5,
			unit:     KB,
			expected: 5 * 1024,
		},
		{
			name:     "Multiple MB to bytes",
			value:    10,
			unit:     MB,
			expected: 10 * 1024 * 1024,
		},
		{
			name:     "Zero value",
			value:    0,
			unit:     GB,
			expected: 0,
		},
		{
			name:     "Large value",
			value:    1000,
			unit:     TB,
			expected: 1000 * 1024 * 1024 * 1024 * 1024,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &DataSize{
				Value: tt.value,
				Unit:  tt.unit,
			}
			result := ds.ToBytes()
			if result != tt.expected {
				t.Errorf("ToBytes() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFromString(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    *DataSize
		expectError bool
	}{
		{
			name:  "Valid KB string",
			input: "1024KB",
			expected: &DataSize{
				Value: 1024,
				Unit:  KB,
			},
			expectError: false,
		},
		{
			name:  "Valid MB string",
			input: "5MB",
			expected: &DataSize{
				Value: 5,
				Unit:  MB,
			},
			expectError: false,
		},
		{
			name:  "Valid GB string",
			input: "2GB",
			expected: &DataSize{
				Value: 2,
				Unit:  GB,
			},
			expectError: false,
		},
		{
			name:  "Valid TB string",
			input: "1TB",
			expected: &DataSize{
				Value: 1,
				Unit:  TB,
			},
			expectError: false,
		},
		{
			name:  "Valid PB string",
			input: "3PB",
			expected: &DataSize{
				Value: 3,
				Unit:  PB,
			},
			expectError: false,
		},
		{
			name:        "Invalid string - no unit",
			input:       "1024",
			expected:    nil,
			expectError: true,
		},
		{
			name:        "Invalid string - unknown unit",
			input:       "1024EB",
			expected:    nil,
			expectError: true,
		},
		{
			name:        "Empty string",
			input:       "",
			expected:    nil,
			expectError: true,
		},
		{
			name:        "Invalid string - just unit",
			input:       "KB",
			expected:    nil,
			expectError: true,
		},
		{
			name:        "Invalid string - non-numeric",
			input:       "abcMB",
			expected:    nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := FromString(tt.input)

			if tt.expectError {
				if err == nil {
					t.Errorf("FromString() expected error but got none")
				}
				if result != nil {
					t.Errorf("FromString() expected nil result but got %v", result)
				}
			} else {
				if err != nil {
					t.Errorf("FromString() unexpected error: %v", err)
				}
				if result == nil {
					t.Errorf("FromString() expected result but got nil")
				} else if result.Value != tt.expected.Value || result.Unit != tt.expected.Unit {
					t.Errorf("FromString() = %v, want %v", result, tt.expected)
				}
			}
		})
	}
}

func TestForSuffix(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		suffix      string
		expected    *DataSize
		expectError bool
	}{
		{
			name:   "Valid KB suffix",
			input:  "1024KB",
			suffix: KB,
			expected: &DataSize{
				Value: 1024,
				Unit:  KB,
			},
			expectError: false,
		},
		{
			name:   "Valid MB suffix",
			input:  "5MB",
			suffix: MB,
			expected: &DataSize{
				Value: 5,
				Unit:  MB,
			},
			expectError: false,
		},
		{
			name:   "Valid GB suffix",
			input:  "2GB",
			suffix: GB,
			expected: &DataSize{
				Value: 2,
				Unit:  GB,
			},
			expectError: false,
		},
		{
			name:   "Valid TB suffix",
			input:  "1TB",
			suffix: TB,
			expected: &DataSize{
				Value: 1,
				Unit:  TB,
			},
			expectError: false,
		},
		{
			name:   "Valid PB suffix",
			input:  "3PB",
			suffix: PB,
			expected: &DataSize{
				Value: 3,
				Unit:  PB,
			},
			expectError: false,
		},
		{
			name:        "Invalid input - non-numeric",
			input:       "abcKB",
			suffix:      KB,
			expected:    nil,
			expectError: true,
		},
		{
			name:        "Empty input",
			input:       "",
			suffix:      KB,
			expected:    nil,
			expectError: true,
		},
		{
			name:   "Zero value",
			input:  "0MB",
			suffix: MB,
			expected: &DataSize{
				Value: 0,
				Unit:  MB,
			},
			expectError: false,
		},
		{
			name:   "Large value",
			input:  "999999999GB",
			suffix: GB,
			expected: &DataSize{
				Value: 999999999,
				Unit:  GB,
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ForSuffix(tt.input, tt.suffix)

			if tt.expectError {
				if err == nil {
					t.Errorf("ForSuffix() expected error but got none")
				}
				if result != nil {
					t.Errorf("ForSuffix() expected nil result but got %v", result)
				}
			} else {
				if err != nil {
					t.Errorf("ForSuffix() unexpected error: %v", err)
				}
				if result == nil {
					t.Errorf("ForSuffix() expected result but got nil")
				} else if result.Value != tt.expected.Value || result.Unit != tt.expected.Unit {
					t.Errorf("ForSuffix() = %v, want %v", result, tt.expected)
				}
			}
		})
	}
}

func TestConstants(t *testing.T) {
	tests := []struct {
		name     string
		constant string
		expected string
	}{
		{"Bytes constant", Bytes, "Bytes"},
		{"KB constant", KB, "KB"},
		{"MB constant", MB, "MB"},
		{"GB constant", GB, "GB"},
		{"TB constant", TB, "TB"},
		{"PB constant", PB, "PB"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant != tt.expected {
				t.Errorf("Constant %s = %v, want %v", tt.name, tt.constant, tt.expected)
			}
		})
	}
}

func TestMultipliers(t *testing.T) {
	tests := []struct {
		name       string
		multiplier int64
		expected   int64
	}{
		{"MultiBytes", MultiBytes, 1},
		{"MultiKB", MultiKB, 1024},
		{"MultiMB", MultiMB, 1024 * 1024},
		{"MultiGB", MultiGB, 1024 * 1024 * 1024},
		{"MultiTB", MultiTB, 1024 * 1024 * 1024 * 1024},
		{"MultiPB", MultiPB, 1024 * 1024 * 1024 * 1024 * 1024},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.multiplier != tt.expected {
				t.Errorf("Multiplier %s = %v, want %v", tt.name, tt.multiplier, tt.expected)
			}
		})
	}
}

func TestDataSize_EdgeCases(t *testing.T) {
	t.Run("Negative value", func(t *testing.T) {
		ds := &DataSize{Value: -1, Unit: MB}
		result := ds.ToBytes()
		expected := int64(-1024 * 1024)
		if result != expected {
			t.Errorf("ToBytes() with negative value = %v, want %v", result, expected)
		}
	})

	t.Run("Maximum int64 value", func(t *testing.T) {
		ds := &DataSize{Value: 9223372036854775807, Unit: Bytes}
		result := ds.ToBytes()
		if result != 9223372036854775807 {
			t.Errorf("ToBytes() with max int64 = %v, want 9223372036854775807", result)
		}
	})

	t.Run("String with negative value", func(t *testing.T) {
		ds := &DataSize{Value: -5, Unit: GB}
		result := ds.String()
		expected := "-5GB"
		if result != expected {
			t.Errorf("String() with negative value = %v, want %v", result, expected)
		}
	})
}
