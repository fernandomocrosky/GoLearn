package banco

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Conectar() (*sql.DB, error) {
	user := "root"
	password := ""
	dbName := "devbook"
	stringConexao := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", user, password, dbName)

	db, erro := sql.Open("mysql", stringConexao)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
