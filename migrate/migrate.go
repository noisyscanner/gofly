package migrate

import (
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/noisyscanner/gofly/gofly"
)

func getMigrationsPath() string {
	envVar := os.Getenv("MIGRATIONS_PATH")
	if envVar != "" {
		return envVar
	}
	return "/migrations"
}

func getMigrate(configService gofly.ConfigService) (*migrate.Migrate, error) {
	config := configService.GetConfig()
	migrationsPath := fmt.Sprintf("file://%s", getMigrationsPath())
	mysqlDbStr := fmt.Sprintf("mysql://%s", config.DBString())
	return migrate.New(migrationsPath, mysqlDbStr)
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

	return m.Down()
}
