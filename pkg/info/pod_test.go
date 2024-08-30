package info

import (
	"os"
	"testing"

	"k8s.io/apimachinery/pkg/types"
)

func TestGetCurrentPodNamespaceAndName(t *testing.T) {
	tests := []struct {
		name           string
		namespaceEnv   string
		nameEnv        string
		wantErr        bool
		expectedResult *types.NamespacedName
	}{
		{
			name:         "BothSet",
			namespaceEnv: "namespace1",
			nameEnv:      "pod1",
			wantErr:      false,
			expectedResult: &types.NamespacedName{
				Namespace: "namespace1",
				Name:      "pod1",
			},
		},
		{
			name:         "NamespaceNotSet",
			namespaceEnv: "",
			nameEnv:      "pod1",
			wantErr:      true,
		},
		{
			name:         "NameNotSet",
			namespaceEnv: "namespace1",
			nameEnv:      "",
			wantErr:      true,
		},
		{
			name:         "BothNotSet",
			namespaceEnv: "",
			nameEnv:      "",
			wantErr:      true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			os.Setenv("POD_NAMESPACE", tc.namespaceEnv)
			os.Setenv("POD_NAME", tc.nameEnv)

			result, err := GetCurrentPodNamespaceAndName()
			if (err != nil) != tc.wantErr {
				t.Errorf("Expected error: %v, but got: %v", tc.wantErr, err)
			}
			if err == nil && !compareNamespacedNames(result, tc.expectedResult) {
				t.Errorf("Expected result: %v, but got: %v", tc.expectedResult, result)
			}
		})
	}
}

func compareNamespacedNames(a, b *types.NamespacedName) bool {
	return a.Namespace == b.Namespace && a.Name == b.Name
}

func TestGetCurrentPodInfo(t *testing.T) {
	tests := []struct {
		name           string
		namespaceEnv   string
		nameEnv        string
		wantErr        bool
		expectedResult string
	}{
		{
			name:           "BothSet",
			namespaceEnv:   "namespace1",
			nameEnv:        "pod1",
			wantErr:        false,
			expectedResult: "namespace1-pod1",
		},
		{
			name:         "NamespaceNotSet",
			namespaceEnv: "",
			nameEnv:      "pod1",
			wantErr:      true,
		},
		{
			name:         "NameNotSet",
			namespaceEnv: "namespace1",
			nameEnv:      "",
			wantErr:      true,
		},
		{
			name:         "BothNotSet",
			namespaceEnv: "",
			nameEnv:      "",
			wantErr:      true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			os.Setenv("POD_NAMESPACE", tc.namespaceEnv)
			os.Setenv("POD_NAME", tc.nameEnv)

			result, err := GetCurrentPodInfo()
			if (err != nil) != tc.wantErr {
				t.Errorf("Expected error: %v, but got: %v", tc.wantErr, err)
			}
			if err == nil && result != tc.expectedResult {
				t.Errorf("Expected result: %s, but got: %s", tc.expectedResult, result)
			}
		})
	}
}
