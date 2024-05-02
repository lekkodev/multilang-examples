// Generated by Lekko. DO NOT EDIT.
package lekkoexamples

import (
	"context"

	client "github.com/lekkodev/go-sdk/client"
	configv1beta1 "github.com/lekkodev/multilang-examples/golang/lekko/proto/examples/config/v1beta1"
)

type LekkoClient struct {
	client.Client
}

// tunable boolean
func (c *LekkoClient) GetBooleanTunable() bool {
	ctx := context.Background()

	result, err := c.GetBool(ctx, "examples", "boolean-tunable")
	if err == nil {
		return result
	}
	return getBooleanTunable()
}

// tunable number
func (c *LekkoClient) GetNumberTunable() float64 {
	ctx := context.Background()

	result, err := c.GetFloat(ctx, "examples", "number-tunable")
	if err == nil {
		return result
	}
	return getNumberTunable()
}

// tunable string
func (c *LekkoClient) GetStringTunable() string {
	ctx := context.Background()

	result, err := c.GetString(ctx, "examples", "string-tunable")
	if err == nil {
		return result
	}
	return getStringTunable()
}

// test boolean operators
func (c *LekkoClient) GetTestBooleanOperators(isTest bool) string {
	ctx := context.Background()
	ctx = client.Add(ctx, "is_test", isTest)
	result, err := c.GetString(ctx, "examples", "test-boolean-operators")
	if err == nil {
		return result
	}
	return getTestBooleanOperators(isTest)
}

func (c *LekkoClient) GetTestLogicalOperators(env string, isTest bool, version float64) string {
	ctx := context.Background()
	ctx = client.Add(ctx, "env", env)
	ctx = client.Add(ctx, "is_test", isTest)
	ctx = client.Add(ctx, "version", version)
	result, err := c.GetString(ctx, "examples", "test-logical-operators")
	if err == nil {
		return result
	}
	return getTestLogicalOperators(env, isTest, version)
}

// test number operators
func (c *LekkoClient) GetTestNumberOperators(version float64) string {
	ctx := context.Background()
	ctx = client.Add(ctx, "version", version)
	result, err := c.GetString(ctx, "examples", "test-number-operators")
	if err == nil {
		return result
	}
	return getTestNumberOperators(version)
}

// test string operators
func (c *LekkoClient) GetTestStringOperators(env string) string {
	ctx := context.Background()
	ctx = client.Add(ctx, "env", env)
	result, err := c.GetString(ctx, "examples", "test-string-operators")
	if err == nil {
		return result
	}
	return getTestStringOperators(env)
}

// tunable interface
func (c *LekkoClient) GetTunableInterface(env string) *configv1beta1.TunableStruct {
	ctx := context.Background()
	ctx = client.Add(ctx, "env", env)
	result := &configv1beta1.TunableStruct{}
	err := c.GetProto(ctx, "examples", "tunable-interface", result)
	if err == nil {
		return result
	}
	return getTunableInterface(env)
}
