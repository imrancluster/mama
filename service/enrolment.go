package service

import (
	"database/sql"
	"fmt"

	// _ mysql using for internal usages
	_ "github.com/go-sql-driver/mysql"
)

// EnrolmentDB provie db connection
var EnrolmentDB *sql.DB
var err error

func ini() {

	EnrolmentDB, err = sql.Open("mysql", "revamp:mytonic!23@tcp(mytonic-staging.ckvp0ck3llgr.ap-southeast-1.rds.amazonaws.com:3306)/dh_enrolment?charset=utf8")
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
