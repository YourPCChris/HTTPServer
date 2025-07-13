
package main


import(
    "fmt"
    "database/sql"
    _"github.com/mattn/go-sqlite3"
    "golang.org/x/crypto/bcrypt"
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
        password TEXT NOT NULL,
        isAdmin INTEGER NOT NULL
        );`

    _, err = db.Exec(table)
    if err != nil{
        panic(err)
    }

    /*
    _, err = db.Exec("DELETE FROM users WHERE username = ?", "")
    if err != nil{
        panic(err)
    }
    */

    /*
    firstPassword, err := bcrypt.GenerateFromPassword([]byte(""), bcrypt.DefaultCost)
    if err != nil{
        fmt.Println("Hash Failed")
        return
    }
    */

    /*
    _, err = db.Exec("INSERT INTO users (username, password, isAdmin) VALUES (?, ?, ?)", "Chris", string(firstPassword), true)
    if err != nil{
        panic(err)
    }
    */

    fmt.Println("Finsihed")
}
