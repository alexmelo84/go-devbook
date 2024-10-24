package db

import (
	"api/src/config"
	"database/sql"

	_"github.com/go-sql-driver/mysql" // Driver MySQL
)

// Conectar conecta ao banco de dados
func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConexaoBd)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
