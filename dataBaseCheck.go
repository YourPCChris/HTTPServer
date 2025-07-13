
package main

import (
    "fmt"
    "database/sql"
    _"github.com/mattn/go-sqlite3"
    "golang.org/x/crypto/bcrypt"
)

func checkUser(name, pass string) (bool, bool){
    //fmt.Println("checking")

    db, err := sql.Open("sqlite3", "./DB/credDB.db")
    if err != nil{
        panic(err)
    }
    defer db.Close()


    query := "SELECT isAdmin FROM users WHERE username = ? AND password = ?"
    var isAdmin bool 
    newHashedPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
    if err != nil {
        fmt.Println("Failed to hash password")
        return false, false
    }

    err = db.QueryRow(query, name, string(newHashedPass)).Scan(&isAdmin)

    if err == sql.ErrNoRows{
        //fmt.Println("No user found")
        return false, false
    }else if err != nil {
        panic(err)
        return false, false
    }else {
        fmt.Println("User found")
        return true, isAdmin
    }
}

func addUserToDB(name, pass string, admin bool) bool {
    var command string = "INSERT INTO users (username, password, isAdmin) VALUES (?, ?, ?)"

    db, err := sql.Open("sql3", "/DB/credDB.db")
    if err != nil{
        fmt.Println("Failed to open DB")
        return false;
    }
    defer db.Close()

    newHashedPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
    _, err = db.Exec(command, name, string(newHashedPass), admin)
    if err != nil{
        fmt.Println("Failed to add user")
        return false
    }

    return true;
}
