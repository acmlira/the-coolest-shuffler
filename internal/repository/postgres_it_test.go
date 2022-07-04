// +build all

package repository

import (
	"io/ioutil"
	"testing"

	"the-coolest-shuffler/configs"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestPostgres(t *testing.T) {
	log.SetOutput(ioutil.Discard)

	t.Run("verify connection", func(t *testing.T) {
		configs.Init()
		Database(configs.GetPostgresDSN())
		var result int
		GetDatabase().Raw("SELECT 1+1 AS result").Scan(&result)
		assert.Equal(t, 2, result)
	})
}
