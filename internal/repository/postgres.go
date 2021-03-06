package repository

import (
	"strconv"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	database *gorm.DB
}

func NewDatabase(dsn string) *Database {
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error(err)
		return &Database{}
	}
	return &Database{database: database}
}

func (d *Database) Select(table string, target interface{}, filter map[string][]string) interface{} {
	var query *gorm.DB = d.database
	for k, v := range filter {
		if len(v) > 0 {
			query = query.Where(k + " IN ?", v)
		}
	}

	var result = query.Table(table).Find(&target)
	if result.Error != nil {
		log.Error(result.Error)
	} else {
		log.Debug("Retrived " + strconv.Itoa(int(result.RowsAffected)) + " objects")
	}
	return target
}
