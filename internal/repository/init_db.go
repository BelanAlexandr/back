package repository

import (
	"database/sql"

	"github.com/BelanAlexandr/back/internal/config"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB(cfg *config.Config) {
	connStr := cfg.ConnectionString
	base, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	// driver, err := postgres.WithInstance(base, &postgres.Config{})
	// if err != nil {
	// 	log.Fatalf("Не удалось создать драйвер миграций: %v", err)
	// }

	// m, err := migrate.NewWithDatabaseInstance(
	// 	"file://migrate",
	// 	"postgres",
	// 	driver,
	// )
	// if err != nil {
	// 	log.Fatalf("Ошибка инициализации мигратора: %v", err)
	// }
	// err = m.Up()
	// if err != nil && err != migrate.ErrNoChange {
	// 	log.Fatalf("Ошибка применения миграций: %v", err)
	// }
	db = base
}
