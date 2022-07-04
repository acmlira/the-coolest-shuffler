// +build all

package repository

import (
	"io/ioutil"
	"testing"
	"the-coolest-shuffler/configs"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestRedis(t *testing.T) {
	type scenario struct {
		value        string
		defaultValue string
		expected     interface{}
		name         string
	}

	log.SetOutput(ioutil.Discard)

	putScenario := scenario{
		"95bf44f6-exx6-447f-80e5-ef6694f13e3x",
		"{\"deck_id\":\"95bf44f6-exx6-447f-80e5-ef6694f13e3x\",\"remaining\":1,\"shuffled\":false,\"cards\":[]}",
		"{\"deck_id\":\"95bf44f6-exx6-447f-80e5-ef6694f13e3x\",\"remaining\":1,\"shuffled\":false,\"cards\":[]}",
		"set and then get",
	}

	removeScenario := scenario{
		"95bf44f6-exx6-447f-80e5-ef6694f13e3x",
		"{\"deck_id\":\"95bf44f6-exx6-447f-80e5-ef6694f13e3x\",\"remaining\":1,\"shuffled\":false,\"cards\":[]}",
		"",
		"delete then get",
	}

	t.Run(putScenario.name, func(t *testing.T) {
		configs.Init()
		Cache(
			configs.GetRedisHost(),
			configs.GetRedisPort(),
			configs.GetRedisDatabase(),
			configs.GetRedisPassword())

		Set(putScenario.value, putScenario.defaultValue)
		value := Get(putScenario.value)
		assert.Equal(t, putScenario.expected, value)
	})

	t.Run(removeScenario.name, func(t *testing.T) {
		configs.Init()
		Cache(
			configs.GetRedisHost(),
			configs.GetRedisPort(),
			configs.GetRedisDatabase(),
			configs.GetRedisPassword())

		Del(removeScenario.value)
		value := Get(removeScenario.value)
		assert.Equal(t, removeScenario.expected, value)
	})
}
