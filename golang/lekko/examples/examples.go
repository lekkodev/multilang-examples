package lekkoexamples

import (
	"strings"

	configv1beta1 "github.com/lekkodev/test-multilang/lekko/proto/examples/config/v1beta1"
	"golang.org/x/exp/slices"
)

// tunable boolean
func getBooleanTunable() bool {
	return true
}

// tunable number
func getNumberTunable() float64 {
	return 42
}

// tunable string
func getStringTunable() string {
	return "foo"
}

// test boolean operators
func getTestBooleanOperators(isTest bool) string {
	if isTest == true {
		return "true"
	} else if isTest == false {
		return "false"
	}
	return "default"
}

// test number operators
func getTestNumberOperators(version float64) string {
	if version == 1 {
		return "equals"
	} else if version != 2 {
		return "not equals"
	} else if version > 3 {
		return "greater"
	} else if version >= 4 {
		return "greater or equals"
	} else if version < 5 {
		return "less"
	} else if version <= 6 {
		return "less or equals"
	} else if slices.Contains([]float64{1, 3, 5}, version) {
		return "in"
	}
	return "default"
}

// test string operators
func getTestStringOperators(env string) string {
	if env == "prod" {
		return "equals"
	} else if env != "dev" {
		return "not equals"
	} else if strings.Contains(env, "test") {
		return "test"
	} else if strings.HasPrefix(env, "staging") {
		return "staging"
	} else if strings.HasSuffix(env, "beta") {
		return "beta"
	} else if slices.Contains([]string{"special1", "special2"}, env) {
		return "special"
	}
	return "default"
}

// tunable interface
func getTunableInterface() *configv1beta1.TunableStruct {
	return &configv1beta1.TunableStruct{
		BooleanField: true,
		NumberField:  42,
		StringField:  "default",
	}
}
