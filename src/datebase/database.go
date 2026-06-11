package datebase

import (
	"database/sql"
	"overtime_system_menagement/src/config"

	_ "github.com/go-sql-driver/mysql"
)

func Connection() (*sql.DB, error) {

	db, erro := sql.Open("mysql", config.StringConnectionDb)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, erro
}
