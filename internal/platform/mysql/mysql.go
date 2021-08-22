package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var c *client

const driver = "mysql"

type cfg struct {
	dbName          string
	username        string
	password        string
	host            string
	maxOpenConn     int
	maxIdleConn     int
	connMaxLifetime time.Duration
}

type client struct {
	db *sql.DB
}

func (m *client) connect(config cfg) error {
	uri := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", config.username, config.password, config.host, config.dbName)
	conn, err := sql.Open(driver, uri)
	if err != nil {
		return err
	}

	if err = conn.Ping(); err != nil {
		return err
	}

	m.db = conn
	conn.SetMaxOpenConns(config.maxOpenConn)
	conn.SetMaxIdleConns(config.maxIdleConn)
	conn.SetConnMaxLifetime(config.connMaxLifetime)

	return nil
}

func NewDefaultClient() *client {
	return c
}

func init() {
	c = &client{}
	err := c.connect(cfg{
		dbName:          "health",
		username:        "root",
		password:        "root",
		host:            "127.0.0.1:3306",
		maxOpenConn:     500,
		maxIdleConn:     10,
		connMaxLifetime: 1000 * time.Millisecond,
	})

	if err != nil {
		panic(err)
	}
}
