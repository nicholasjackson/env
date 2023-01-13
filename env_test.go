package env

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func setEnv(name, value string) func() {
	os.Setenv(name, value)

	return func() {
		os.Unsetenv(name)
	}
}

func TestStringSetEnv(t *testing.T) {
	envs = make([]envVar, 0)
	cleanup := setEnv("nic", "is awesome")
	defer cleanup()

	n := String("nic", false, "", "something")
	Parse()

	assert.Equal(t, "is awesome", *n)
}

func TestEmptyStringRequiredTrueEnv(t *testing.T) {
	envs = make([]envVar, 0)
	testEnv := "nic"
	os.Unsetenv(testEnv)
	required := true

	String(testEnv, required, "is awesome", "something")
	err := Parse()
	assert.Error(t, err)
}

func TestEmptyStringRequiredFalseEnv(t *testing.T) {
	envs = make([]envVar, 0)
	testEnv := "nic"
	os.Unsetenv(testEnv)
	required := false

	n := String(testEnv, required, "is awesome", "something")
	err := Parse()

	assert.Equal(t, "is awesome", *n)
	assert.NoError(t, err)
}

func TestIntSetEnv(t *testing.T) {
	envs = make([]envVar, 0)
	cleanup := setEnv("nic", "1")
	defer cleanup()

	n := Int("nic", true, 0, "something")
	Parse()

	assert.Equal(t, 1, *n)
}

func TestIntError(t *testing.T) {
	envs = make([]envVar, 0)
	cleanup := setEnv("nic", "a")
	defer cleanup()

	Int("nic", false, 0, "something")
	err := Parse()

	assert.Contains(t, "expected: nic type: integer got: a", err.Error())
}

func TestFloat64SetEnv(t *testing.T) {
	envs = make([]envVar, 0)
	cleanup := setEnv("nic", "1.1")
	defer cleanup()

	n := Float64("nic", true, 0, "something")
	Parse()

	assert.Equal(t, 1.1, *n)
}

func TestFloat64Error(t *testing.T) {
	envs = make([]envVar, 0)
	cleanup := setEnv("nic", "a")
	defer cleanup()

	Float64("nic", false, 0, "something")
	err := Parse()

	assert.Contains(t, "expected: nic type: float got: a", err.Error())
}

func TestBoolSetEnv(t *testing.T) {
	envs = make([]envVar, 0)
	cleanup := setEnv("nic", "true")
	defer cleanup()

	n := Bool("nic", false, false, "something")
	Parse()

	assert.Equal(t, true, *n)
}

func TestBoolError(t *testing.T) {
	envs = make([]envVar, 0)
	cleanup := setEnv("nic", "a")
	defer cleanup()

	Bool("nic", false, false, "something")
	err := Parse()

	assert.Contains(t, "expected: nic type: boolean got: a", err.Error())
}

func TestDurationSetEnv(t *testing.T) {
	envs = make([]envVar, 0)
	cleanup := setEnv("nic", "10s")
	defer cleanup()

	n := Duration("nic", false, 1*time.Second, "something")
	Parse()

	assert.Equal(t, 10*time.Second, *n)
}

func TestDurationError(t *testing.T) {
	envs = make([]envVar, 0)
	cleanup := setEnv("nic", "test")
	defer cleanup()

	Duration("nic", false, 1*time.Second, "something")
	err := Parse()

	assert.Contains(t, "expected: nic type: duration got: test", err.Error())
}

func TestSetsDefault(t *testing.T) {
	envs = make([]envVar, 0)

	n := String("nic", false, "is unset", "something")
	Parse()

	assert.Equal(t, "is unset", *n)
}

func TestHelp(t *testing.T) {
	envs = make([]envVar, 0)
	String("SERVER_URI", true, "localhost:8181", "URI for upstream server, i.e. localhost:8181")
	String("API_KEY", true, "", "API key for upstream server")
	Int("TIMEOUT", true, 12, "Timeout duration in seconds")
	h := Help()

	fmt.Println(h)
}
