package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
//Mysql ----------------------------------------------------------
type Mysql struct {
	DB *sql.DB
	execute mysqlExecute
	getStr mysqlGetStr
	getStrAry mysqlGetStrAry
}
//mysql constructor
func newMysql(id string, ps string, endpoint string, port int, schema string)(mysql Mysql, err error){
	DB, err := sql.Open("mysql", id+":"+ps+"@tcp("+endpoint+":"+tools.getStr.fromInt(port)+")/"+schema+"?multiStatements=true")
	if err != nil {
		return mysql, err
	}
	err = DB.Ping()
	if err != nil {
		return mysql, err
	}
	mysql.DB = DB
	mysql.execute.DB = DB
	mysql.getStr.DB = DB
	mysql.getStrAry.DB = DB
	return mysql, nil
}

//mysql.execute
type mysqlExecute struct {
	DB *sql.DB
}
//mysql.execute.query
func (mysql *mysqlExecute) query (queryStr string) (err error) {
	_, err = mysql.DB.Exec(queryStr)
	if err != nil {
		return err
	}
	return nil
}

type mysqlGetStr struct {
	DB *sql.DB
}
//mysql.getStr.paymentIDFromUID
func (mysql *mysqlGetStr) paymentIDFromUID (UIDStr string) (paymentIDStr string, err error){
	rows, err := mysql.DB.Query(`SELECT * FROM adiy.firebase_uid WHERE uid = "`+UIDStr+`"`)
	if err != nil {
		return "", err
	}
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	fmt.Printf("col = %v\n", columns)
	return "", nil
}

//mysql.getStrAry
type mysqlGetStrAry struct {
	DB *sql.DB
}
//mysql.getStrAry.query
func (mysql *mysqlGetStrAry) query (queryStr string)(resStrAry []string, err error){
	rows, err := mysql.DB.Query(queryStr)
	if err != nil {
		return nil, err
	}
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	return columns, nil
}