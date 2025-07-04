package utils_test

import (
	"testing"

	"aidanhonan.com/backend/utils"
)

var testEnvName string = "TEST_ENV_VAR_UNUSED"

func TestGetEnv_HappyPath(t *testing.T) {
	expectedEnv := "test"
	t.Setenv(testEnvName, expectedEnv)

	actualEnv := utils.GetEnv(testEnvName, "default")

	if expectedEnv != actualEnv {
		t.Errorf(`Expected "%s" but found "%s"`, expectedEnv, actualEnv)
	}
}

func TestGetEnv_UsesDefault(t *testing.T) {
	expectedEnv := "default"

	actualEnv := utils.GetEnv(testEnvName, "default")

	if expectedEnv != actualEnv {
		t.Errorf(`Expected "%s" but found "%s"`, expectedEnv, actualEnv)
	}
}

func TestGetEnv_PanicWithNoDefault(t *testing.T) {
	// GetEnv should panic and call recover()
	defer func() {
		if recover() == nil {
			t.Errorf("GetEnv() didn't panic when passed non-existant env var")
		}
	}()

	utils.GetEnv(testEnvName, "")
}
