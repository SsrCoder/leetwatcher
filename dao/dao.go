package dao

import (
	"github.com/SsrCoder/leetwatcher/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Dao struct {
	mysql  *gorm.DB
	config *config.DBConfig
}

func New(config *config.DBConfig) (*Dao, error) {
	m, err := newMySQLClient(config.MySQL)
	if err != nil {
		return nil, err
	}

	return &Dao{
		mysql:  m,
		config: config,
	}, nil
}

func newMySQLClient(config *config.MySQLConfig) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(config.URL))
}
