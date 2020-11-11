package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
//Mysql ----------------------------------------------------------
type Mysql struct {
	DB *sql.DB
	execute mysqlExecute
	getStrAry mysqlGetStrAry
}
//Mysql constructor
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
	mysql.getStrAry.DB = DB
	mysql.execute.DB = DB
	return mysql, nil
}

//mysql.execute
type mysqlExecute struct {
	DB *sql.DB
}
//Mysql.execute.query
func (mysql *mysqlExecute) query (queryStr string) (err error) {
	_, err = mysql.DB.Exec(queryStr)
	if err != nil {
		return err
	}
	return nil
}

//mysql.getStrAry
type mysqlGetStrAry struct {
	DB *sql.DB
}
//Mysql.getStrAry.query
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