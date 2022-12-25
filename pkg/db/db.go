// Пакет для работы с БД приложения GoNews.
package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	username = "root"
	password = "root"
	hostname = "127.0.0.1"
	port     = 3306
	dbName   = "newsdb"
)

// База данных.
type DB struct {
	pool *sql.DB
}

type PostComment struct {
	Id      int
	Comment string
}

func New() (*DB, error) {

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		username,
		password,
		hostname,
		port,
		dbName,
	)

	log.Println("connString: ", connString)

	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, fmt.Errorf("mysql err: %s", err)
	}

	return &DB{
		pool: db,
	}, nil
}

func (db *DB) Validate(id int) error {
	query := fmt.Sprintf("UPDATE `comment` SET `is_validate` = '1' WHERE `comment`.`id` = %d",
		id,
	)

	_, err := db.pool.Query(query)
	if err != nil {

		return err
	}

	return nil
}
