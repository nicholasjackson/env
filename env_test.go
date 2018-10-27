package env

import (
	"fmt"
	"os"
	"testing"

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

func TestIntSetEnv(t *testing.T) {
	envs = make([]envVar, 0)
	cleanup := setEnv("nic", "1")
	defer cleanup()

	n := Integer("nic", true, 0, "something")
	Parse()

	assert.Equal(t, 1, *n)
}

func TestIntError(t *testing.T) {
	envs = make([]envVar, 0)
	cleanup := setEnv("nic", "a")
	defer cleanup()

	Integer("nic", false, 0, "something")
	err := Parse()

	assert.Contains(t, "expected: nic type: integer got: a", err.Error())
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
	Integer("TIMEOUT", true, 12, "Timeout duration in seconds")
	h := Help()

	fmt.Println(h)
}
