package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"flag"
)

var envs []envVar
var help = flag.Bool("help", false, "--help to show env help")

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

// String captures the value of an environment varialbe
// name: the name of the environment variable
// required: if set to true and environment variable does not exist an error will be raised
// defaultValue: the default value to return if the environment variable is not set
// help: help string related to the variable
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

// Int something
func Int(name string, required bool, defaultValue int, help string) *int {
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

// Float64 something
func Float64(name string, required bool, defaultValue float64, help string) *float64 {
	v := new(float64)

	envs = append(envs, envVar{
		v,
		name,
		"float",
		required,
		defaultValue,
		help,
		func(a interface{}, b string) error {
			v, err := strconv.ParseFloat(b, 64)
			if err != nil {
				a = nil
				return err
			}

			*a.(*float64) = float64(v)

			return nil
		},
		func(a interface{}, b interface{}) {
			*a.(*float64) = b.(float64)
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

// Duration something
func Duration(name string, required bool, defaultValue time.Duration, help string) *time.Duration {
	v := new(time.Duration)

	envs = append(envs, envVar{
		v,
		name,
		"duration",
		required,
		defaultValue,
		help,
		func(a interface{}, b string) error {
			v, err := time.ParseDuration(b)
			if err != nil {
				a = nil
				return err
			}

			*a.(*time.Duration) = v

			return nil
		},
		func(a interface{}, b interface{}) {
			*a.(*time.Duration) = b.(time.Duration)
		},
		new(string),
	})

	return v
}

// Parse something
func Parse() error {
	// Parse the main flags package to enable the --help option
	flag.Parse()
	if *help == true {
		fmt.Println("Configuration values are set using environment variables, for info please see the following list.")
		fmt.Println("")
		fmt.Println(Help())

		os.Exit(0)
	}

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
