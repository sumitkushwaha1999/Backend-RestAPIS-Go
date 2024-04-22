package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
)

const (
	server                = "sql6.freesqldatabase.com"
	userName              = "sql6700764"
	password              = "QUQpnJqI3l"
	port                  = "3306"
	dataBase              = "sql6700764"
	maxOpenConntion       = 20
	maxIdleConnection     = 10
	maxIdleConnectionTime = time.Minute
)

var (
	DBClient *sql.DB
)

func InitDB() error {
	var err error
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", userName, password, server, port, dataBase)
	// Open a database connection
	db, err := sql.Open("mysql", connString)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println(connString)
	db.SetConnMaxIdleTime(maxIdleConnectionTime)
	db.SetMaxIdleConns(maxIdleConnection)
	db.SetMaxOpenConns(maxOpenConntion)
	fmt.Println(db)

	DBClient = db
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
