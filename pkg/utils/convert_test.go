package utils

import (
	"testing"
)

type TestStruct struct {
	Field1 string
	Field2 int
}

func TestConvertJsonObjectToStruct(t *testing.T) {
	tests := []struct {
		name    string
		source  interface{}
		target  interface{}
		wantErr bool
	}{
		{
			name:    "Success",
			source:  map[string]interface{}{"Field1": "value1", "Field2": 2},
			target:  &TestStruct{},
			wantErr: false,
		},
		{
			name:    "SourceNil",
			source:  nil,
			target:  &TestStruct{},
			wantErr: true,
		},
		{
			name:    "TargetNil",
			source:  map[string]interface{}{"Field1": "value1", "Field2": 2},
			target:  nil,
			wantErr: true,
		},
		{
			name:    "SourceInvalid",
			source:  123,
			target:  &TestStruct{},
			wantErr: true,
		},
		{
			name:    "TargetInvalid",
			source:  map[string]interface{}{"Field1": "value1", "Field2": 2},
			target:  TestStruct{},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := ConvertJsonObjectToStruct(tc.source, tc.target)
			if (err != nil) != tc.wantErr {
				t.Errorf("Expected error: %v, but got: %v", tc.wantErr, err)
			}
		})
	}
}
