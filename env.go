package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var envs []envVar

func init() {
	envs = make([]envVar, 0)
}

type envVar struct {
	value        interface{}
	name         string
	varType      string
	required     bool
	defaultValue interface{}
	help         string
	setValue     func(interface{}, string) error
	setDefault   func(interface{}, interface{})
	envValue     *string
}

// String something
func String(name string, required bool, defaultValue, help string) *string {
	v := new(string)

	envs = append(envs, envVar{
		v,
		name,
		"string",
		required,
		defaultValue,
		help,
		func(a interface{}, b string) error {
			*a.(*string) = b
			return nil
		},
		func(a interface{}, b interface{}) {
			*a.(*string) = b.(string)
		},
		new(string),
	})

	return v
}

// Integer something
func Integer(name string, required bool, defaultValue int, help string) *int {
	v := new(int)

	envs = append(envs, envVar{
		v,
		name,
		"integer",
		required,
		defaultValue,
		help,
		func(a interface{}, b string) error {
			v, err := strconv.ParseInt(b, 0, 64)
			if err != nil {
				a = nil
				return err
			}

			*a.(*int) = int(v)

			return nil
		},
		func(a interface{}, b interface{}) {
			*a.(*int) = b.(int)
		},
		new(string),
	})

	return v
}

// Bool something
func Bool(name string, required bool, defaultValue bool, help string) *bool {
	v := new(bool)

	envs = append(envs, envVar{
		v,
		name,
		"boolean",
		required,
		defaultValue,
		help,
		func(a interface{}, b string) error {
			v, err := strconv.ParseBool(b)
			if err != nil {
				a = nil
				return err
			}

			*a.(*bool) = bool(v)

			return nil
		},
		func(a interface{}, b interface{}) {
			*a.(*bool) = b.(bool)
		},
		new(string),
	})

	return v
}

// Parse something
func Parse() error {
	errors := make([]string, 0)

	for _, e := range envs {
		err := processEnvVar(e)
		if err != nil {
			errors = append(errors, fmt.Sprintf("expected: %s type: %s got: %s", e.name, e.varType, *e.envValue))
		}
	}

	if len(errors) > 0 {
		errString := strings.Join(errors, "\n")
		return fmt.Errorf(errString)
	}

	return nil
}

func processEnvVar(e envVar) error {
	*e.envValue = os.Getenv(e.name)
	if *e.envValue == "" && !e.required {
		e.setDefault(e.value, e.defaultValue)
		return nil
	}

	err := e.setValue(e.value, *e.envValue)
	if err != nil {
		return err
	}

	return nil
}

// Help is help
func Help() string {
	h := make([]string, 1)
	h[0] = "Environment variables:"

	for _, e := range envs {
		def := fmt.Sprintf("'%v'", e.defaultValue)
		if def == "''" {
			def = "no default"
		}

		h = append(h, "  "+e.name+"  default: "+def)
		h = append(h, "       "+e.help)
	}

	return strings.Join(h, "\n")
}
