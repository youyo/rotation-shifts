package db

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
)

func NewConnection() (*dbr.Session, error) {
	dsn := buildDsn(
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"))
	c, err := dbr.Open("mysql", dsn, nil)
	if err != nil {
		return nil, err
	}
	s := c.NewSession(nil)
	return s, nil
}

func buildDsn(user, password, host, port, database string) string {
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?parseTime=true&loc=Local"
	return dsn
}
