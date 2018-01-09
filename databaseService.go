package main

import "database/sql"

type DatabaseService struct {
	configService ConfigService
}

func (ds *DatabaseService) GetDb() (*sql.DB, error) {
	if ds.configService == nil {
		ds.configService = &FileConfigService{File: "config"}
	}

	config := ds.configService.GetConfig()

	return sql.Open(config.Driver, config.DBString())
}