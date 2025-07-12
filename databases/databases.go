package databases

import (
	"log"
	"log/slog"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewSqlx() *sqlx.DB {
	// 実際にDBに接続する場合は書き換える
	const Dsn = "user:password@tcp(localhost:3306)/testdb?parseTime=true"

	db, err := sqlx.Connect("mysql", Dsn)
	if err != nil {
		log.Fatalln("Failed to connect to database:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalln("Failed to ping database:", err)
	}

	slog.Info("Successfully connected to database")

	return db
}
