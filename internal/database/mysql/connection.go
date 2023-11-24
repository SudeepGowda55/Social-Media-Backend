package mysql

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"log"
	"net/url"

	"github.com/go-sql-driver/mysql"
)

func databaseInstance() *sql.DB {

	mysql.RegisterTLSConfig("tidb", &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: "gateway01.ap-southeast-1.prod.aws.tidbcloud.com",
	})

	dbURI := "bQuCs15fm2qm7ps.root:Uxe3kohIGOqEKZCE@tcp(gateway01.ap-southeast-1.prod.aws.tidbcloud.com:4000)/sp500insight?tls=tidb"

	conn, _ := url.Parse(dbURI)

	db, err := sql.Open("mysql", conn.String())

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Successfully connected to database!")
	}

	// var dbName string
	err = db.Ping()
	if err != nil {
		log.Fatal("failed to execute query", err)
	}
	// fmt.Println(dbName)

	return db
}

var DbClient *sql.DB = databaseInstance()
