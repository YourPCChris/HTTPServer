
package main

import (
    "fmt"
    "database/sql"
    _"github.com/mattn/go-sqlite3"
)

func checkUser(name, pass string) bool {
    //fmt.Println("checking")

    db, err := sql.Open("sqlite3", "./DB/credDB.db")
    if err != nil{
        panic(err)
    }
    defer db.Close()


    query := "SELECT id FROM users WHERE username = ? AND password = ?"
    var id int
    err = db.QueryRow(query, name, pass).Scan(&id)

    if err == sql.ErrNoRows{
        //fmt.Println("No user found")
        return false
    }else if err != nil {
        panic(err)
        return false
    }else {
        fmt.Println("User found")
        return true
    }
}
