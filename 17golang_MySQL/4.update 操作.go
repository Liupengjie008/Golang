package main

import (
    "fmt"

    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
)

type Person struct {
    UserId   int    `db:"user_id"`
    Username string `db:"username"`
    Sex      string `db:"sex"`
    Email    string `db:"email"`
}

type Place struct {
    Country string `db:"country"`
    City    string `db:"city"`
    TelCode int    `db:"telcode"`
}

var Db *sqlx.DB

func init() {

    database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
    if err != nil {
        fmt.Println("open mysql failed,", err)
        return
    }

    Db = database
}

func main() {

    res, err := Db.Exec("update person set username=? where user_id=?", "stu0003", 1)
    if err != nil {
        fmt.Println("exec failed, ", err)
        return
	}
    row, err := res.RowsAffected()
    if err != nil {
        fmt.Println("rows failed, ",err)
    }
	fmt.Println("update succ:",row)

}


运行结果：

第一次：$ go run main.go 
update succ: 1
#受影响的行数 1

第二次：$ go run main.go 
update succ: 0
#受影响的行数 0