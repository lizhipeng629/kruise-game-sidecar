package template

import (
	"fmt"
	"os"
	"regexp"

	corev1 "k8s.io/api/core/v1"
)

// Expression format:
// ${SELF:VAR_NAME}: Indicates the environment variable of the sidecar itself.
// ${POD:VAR_NAME}: Indicates the environment variable of the Pod.

const (
	pattern = `\$\{(SELF|POD):([^}]+)\}`
)

func ReplaceValue(value string, container *corev1.Container) (string, error) {
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(value)
	if matches == nil {
		return value, nil
	}

	envType := matches[1]
	envName := matches[2]

	var envValue string
	var found bool
	if envType == "SELF" {
		envValue, found = os.LookupEnv(envName)
	} else if envType == "POD" {
		// Search from the environment variables of the container
		for _, envVar := range container.Env {
			if envVar.Name == envName {
				envValue = envVar.Value
				found = true
				break
			}
		}
	} else {
		return "", fmt.Errorf("unknown environment variable type: %s", envType)
	}

	if !found {
		return "", fmt.Errorf("environment variable %s not found", envName)
	}

	return envValue, nil
}
