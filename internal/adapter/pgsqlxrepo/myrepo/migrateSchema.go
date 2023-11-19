// Миграция схемы
package myrepo

import (
	"os"
	"path/filepath"
	"strings"

	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func (r *Repository) MigrateSchema() error {

	// Путь к каталогу с миграциями
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exePath := filepath.Dir(ex) + `/../../`
	exePath = strings.Replace(exePath, `\`, "/", -1)
	//	fmt.Println(exePath)

	driver, err := pgx.WithInstance(r.GetConn(), &pgx.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+exePath+"db/migration",
		"postgres", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
