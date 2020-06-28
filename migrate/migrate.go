package migrate

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
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
