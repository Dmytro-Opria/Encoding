package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"fmt"
)
var (
	db *sql.DB
)

func main() {
	mysqlStr := "root:123@tcp(127.0.0.1:3306)/user_information?parseTime=true&charset=utf8"
	Init(mysqlStr)
	money := 20.5
	for _, v := range getId() {
		addMoney(v, money)
	}
}

func Init(connection string) {

	var err error

	if db, err = sql.Open("mysql", connection); err != nil {
		fmt.Println("Can`t open connection")
		os.Exit(1)
	}

	if err = db.Ping(); err != nil {
		fmt.Println("Can`t ping user_information db")
		os.Exit(1)
	}

	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(64)
}

func getId() []int{
	sqlStr := "SELECT id FROM users WHERE name='Alex'"

	row, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println("Can`t get id", err)
		return []int{}
	}
	defer row.Close()

	ids := make([]int, 0, 1)

	for row.Next() {
		id := 0
		row.Scan(&id)
		ids = append(ids, id)
	}

	return ids
}

func addMoney(id int, money float64)(error){
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	stmt, err := tx.Prepare(`UPDATE money SET money=money+? WHERE id=?`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(money,id)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}

	return err
}
