package database

import (
	"fmt"
	"time"

	"database/sql"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type Options struct {
	Host           string
	Username       string
	Password       string
	Database       string
	Port           int
	MaxConnections int
}

type SQL struct {
	DB      *sql.DB
	prepare map[string]*sql.Stmt
}

var MySQL *SQL

func New(opt Options) {

	mysql, err := sql.Open("mysql", connectionString(opt))
	if err != nil {
		panic(err)
	}

	mysql.SetConnMaxLifetime(time.Minute * 3)
	mysql.SetMaxOpenConns(opt.MaxConnections)
	mysql.SetMaxIdleConns(opt.MaxConnections)

	err = mysql.Ping()
	if err != nil {
		panic(err)
	}

	MySQL = &SQL{
		DB:      mysql,
		prepare: make(map[string]*sql.Stmt),
	}

}

func (sql *SQL) Exec(event string, args ...interface{}) (*sql.Rows, *mysql.MySQLError) {

	if sql.prepare[event] == nil {
		prepare, err := sql.DB.Prepare(event)
		if err != nil {
			sqlErr, ok := err.(*mysql.MySQLError)
			if ok {
				return nil, sqlErr
			}
			return nil, &mysql.MySQLError{}
		}
		sql.prepare[event] = prepare
	}

	rows, err := sql.prepare[event].Query(args...)
	if err != nil {
		sqlErr, ok := err.(*mysql.MySQLError)
		if ok {
			return nil, sqlErr
		}
		return nil, &mysql.MySQLError{}
	}

	return rows, nil

}

func connectionString(opts Options) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", opts.Username, opts.Password, opts.Host, opts.Port, opts.Database)
}
