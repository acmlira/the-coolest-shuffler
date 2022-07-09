package api

import (
	"io/ioutil"
	"testing"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestRequiredUUID(t *testing.T) {
	type scenario struct {
		value      string
		result     interface{}
		shouldFail bool
		name       string
	}

	scenarios := []scenario{
		{"036aa0d3-cadb-42fa-a234-d2253bc8de88", uuid.MustParse("036aa0d3-cadb-42fa-a234-d2253bc8de88"), false, "uuid ok"},
		{"-- not an uuid --", nil, true, "uuid not ok"},
		{"", nil, true, "uuid missing"},
	}

	log.SetOutput(ioutil.Discard)

	for _, it := range scenarios {
		t.Run(it.name, func(t *testing.T) {
			value, err := requiredUUID(it.value, "")
			if it.shouldFail {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, it.result, value)
			}
		})
	}
}

func TestOptionalInt(t *testing.T) {
	type scenario struct {
		value        string
		defaultValue int
		expected     interface{}
		name         string
	}

	scenarios := []scenario{
		{"123", 722, 123, "int ok"},
		{"-1", 321, -1, "negative int ok"},
		{"-- not a int --", 312, 312, "int defaults"},
		{"", 213, 213, "int missing"},
	}

	for _, it := range scenarios {
		t.Run(it.name, func(t *testing.T) {
			value := optionalInt(it.value, "", it.defaultValue)
			assert.Equal(t, it.expected, value)
		})
	}
}

func TestOptionalBool(t *testing.T) {
	type scenario struct {
		value        string
		defaultValue bool
		expected     interface{}
		name         string
	}

	scenarios := []scenario{
		{"true", false, true, "true ok"},
		{"t", false, true, "t ok"},
		{"TRUE", false, true, "TRUE ok"},
		{"T", false, true, "T ok"},
		{"1", false, true, "1 ok"},
		{"false", true, false, "false ok"},
		{"f", true, false, "f ok"},
		{"FALSE", true, false, "FALSE ok"},
		{"F", true, false, "F ok"},
		{"0", true, false, "0 ok"},
		{"-- not a bool --", true, true, "bool defaults"},
		{"", false, false, "bool missing"},
	}

	for _, it := range scenarios {
		t.Run(it.name, func(t *testing.T) {
			value := optionalBool(it.value, "", it.defaultValue)
			assert.Equal(t, it.expected, value)
		})
	}
}

func TestOptionalStringList(t *testing.T) {
	type scenario struct {
		value    string
		expected interface{}
		name     string
	}

	scenarios := []scenario{
		{"AS,2C,4H", []string{"AS", "2C", "4H"}, "string list ok"},
		{"", []string{}, "string list missing"},
	}

	for _, it := range scenarios {
		t.Run(it.name, func(t *testing.T) {
			value := optionalStringList(it.value, "")
			assert.Equal(t, it.expected, value)
		})
	}
}
