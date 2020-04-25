package gofly

import "database/sql"

type DatabaseService struct {
	ConfigService ConfigService
}

func (ds *DatabaseService) GetDb() (*sql.DB, error) {
	if ds.ConfigService == nil {
		ds.ConfigService = &FileConfigService{File: "config"}
	}

	config := ds.ConfigService.GetConfig()

	return sql.Open(config.Driver, config.DBString())
}
