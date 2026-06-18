package repository

import (
	"database/sql"
	"log"

	"github.com/BelanAlexandr/back/internal/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB(cfg *config.Config) {
	connStr := cfg.ConnectionString
	base, err := sql.Open("postgres", connStr)
	log.Println(connStr)
	log.Println("Подключение к базе данных по DB_CONN_STR...")

	if err != nil {
		log.Fatalf("Ошибка открытия БД: %v", err)
	}

	if err := base.Ping(); err != nil {
		log.Fatalf("База данных недоступна: %v", err)
	}

	// Миграции
	driver, err := postgres.WithInstance(base, &postgres.Config{})
	if err != nil {
		log.Fatalf("Не удалось создать драйвер миграций: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrate", "postgres", driver)
	if err != nil {
		log.Fatalf("Ошибка инициализации мигратора: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Ошибка применения миграций: %v", err)
	}

	db = base
	log.Println("База успешно инициализирована, миграции применены!")
}
