
package main


import(
    "fmt"
    "database/sql"
    _"github.com/mattn/go-sqlite3"
)


func main(){
    fmt.Println("Starting")
    db, err := sql.Open("sqlite3", "./credDB.db")
    if err != nil{
        panic(err)
    }
    defer db.Close()

    if err = db.Ping(); err != nil{
        panic(err)
    }

    table := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL
        );`

    _, err = db.Exec(table)
    if err != nil{
        panic(err)
    }
    /*
    _, err = db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", "", "")
    if err != nil{
        panic(err)
    }
    */

    fmt.Println("Finsihed")
}
