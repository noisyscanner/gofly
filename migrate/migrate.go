package migrate

import (
	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/noisyscanner/gofly/gofly"
)

func getMigrate(configService gofly.ConfigService) (*migrate.Migrate, error) {
	config := configService.GetConfig()
	return migrate.New("file:///migrations/", config.DBString())
}

func Up(configService gofly.ConfigService) error {
	m, err := getMigrate(configService)
	if err != nil {
		return err
	}

	return m.Up()
}

func Down(configService gofly.ConfigService) error {
	m, err := getMigrate(configService)
	if err != nil {
		return err
	}

	return m.Up()
}
