package template

import (
	"github.com/agiledragon/gomonkey/v2"
	"github.com/magicsong/kidecar/pkg/info"
	corev1 "k8s.io/api/core/v1"
	"reflect"
	"testing"
)

type ConfigStruct struct {
	Field1 string `parse:"true"`
	Field2 int
}

func TestParseConfig(t *testing.T) {

	patchPod := gomonkey.ApplyFunc(
		info.GetCurrentPod,
		func() (*corev1.Pod, error) {
			return &corev1.Pod{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Env: []corev1.EnvVar{
								{Name: "ENV_VAR", Value: "value"},
							},
						},
					},
				},
			}, nil
		},
	)
	defer patchPod.Reset()

	tests := []struct {
		name           string
		inputConfig    interface{}
		expectedConfig interface{}
		expectedError  error
	}{
		{
			name: "SuccessfulParse",
			inputConfig: &ConfigStruct{
				Field1: "${POD:ENV_VAR}",
				Field2: 10,
			},
			expectedConfig: &ConfigStruct{
				Field1: "value",
				Field2: 10,
			},
			expectedError: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := ParseConfig(tc.inputConfig)
			if err != tc.expectedError {
				t.Errorf("Expected error: %v, but got: %v", tc.expectedError, err)
			}

			if err == nil {
				inputValue := tc.inputConfig
				expectedValue := tc.expectedConfig
				inputReflectValue := reflect.ValueOf(inputValue).Elem()
				expectedReflectValue := reflect.ValueOf(expectedValue).Elem()

				for i := 0; i < inputReflectValue.NumField(); i++ {
					if inputReflectValue.Field(i).Interface() != expectedReflectValue.Field(i).Interface() {
						t.Errorf("Field %s mismatch. Expected: %v, Got: %v", inputReflectValue.Type().Field(i).Name, expectedReflectValue.Field(i).Interface(), inputReflectValue.Field(i).Interface())
					}
				}
			}
		})
	}
}
