
package main

import (
    "fmt"
    "database/sql"
    _"github.com/mattn/go-sqlite3"
    "golang.org/x/crypto/bcrypt"
)


func checkUser(name, pass string) (bool, bool){
    db, err := sql.Open("sqlite3", "./DB/credDB.db")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    query := "SELECT password FROM users WHERE username = ?"
    var storedPassword string
    qerr := db.QueryRow(query, name).Scan(&storedPassword)
    if qerr == sql.ErrNoRows{
        fmt.Println("User does not exist")
        return false, false
    }else if qerr != nil{
        fmt.Println("Failed to query database")
        return false, false 
    }else if qerr == nil{
        err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(pass))
        if err != nil{
            return false, false
        }else if err == nil{
            var isAdmin bool 
            query = "SELECT isAdmin FROM users WHERE username = ?"
            qerr = db.QueryRow(query, name).Scan(&isAdmin)
            return true, isAdmin
        }
    }
    return false, false 
}


func addUserToDB(name, pass string, admin bool) bool {
    var command string = "INSERT INTO users (username, password, isAdmin) VALUES (?, ?, ?)"

    db, err := sql.Open("sqlite3", "./DB/credDB.db")
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
