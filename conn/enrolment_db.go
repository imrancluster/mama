package conn

import (
	"database/sql"
	"fmt"

	// _ mysql connetion
	_ "github.com/go-sql-driver/mysql"
	// postgres conn
	"github.com/imrancluster/mama/config"
	// postgres conn
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// EnrolmentDB provie db connection
var db *sql.DB

// EnrolmentMySQLClient struct
type EnrolmentMySQLClient struct {
	*sql.DB
}

// CorePostgresClient struct
type CorePostgresClient struct {
	*sql.DB
}

// Conn is an instance *sql.DB
var Conn EnrolmentMySQLClient

// CoreCon is an instance *sql.DB
var CoreCon CorePostgresClient

// ConnectEnrolmentDB to provide db connection
func ConnectEnrolmentDB() error {

	cfg := config.EnrolmentDBConfig()
	// "revamp:mytonic!23@tcp(mytonic-staging.ckvp0ck3llgr.ap-southeast-1.rds.amazonaws.com:3306)/dh_enrolment?charset=utf8"
	dbSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?multiStatements=true&charset=utf8",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)

	c, err := sql.Open("mysql", dbSourceName)
	check(err)

	Conn = EnrolmentMySQLClient{
		DB: c,
	}

	return nil
}

// ConnectCoreDB to provide db connection
func ConnectCoreDB() error {

	cfg := config.CoreDBConfig()
	dbSourceName := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Name,
		cfg.Password,
	)

	c, err := sql.Open("postgres", dbSourceName)
	check(err)

	CoreCon = CorePostgresClient{
		DB: c,
	}

	return nil
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// EnrolmentDB to get db connections
func EnrolmentDB() EnrolmentMySQLClient {
	return Conn
}

// CoreDB to get db connections
func CoreDB() CorePostgresClient {
	return CoreCon
}
