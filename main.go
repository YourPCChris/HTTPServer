
package main 

import 
(
    "fmt"
    "net/http"
    "html/template"
)


func login(w http.ResponseWriter, r *http.Request){
    if r.Method == http.MethodPost {
        r.ParseForm()
        username := r.FormValue("username")
        password := r.FormValue("password")

        //check account details against database

        if checkUser(username, password){
            cookie := http.Cookie{
                Name: "session",
                Value: "loggedin",
                Path: "/",
                HttpOnly: true,
            }
            
            http.SetCookie(w, &cookie)
            http.Redirect(w, r, "/Home", http.StatusSeeOther)
            return 
        }
        http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
    }
}

func requireLogin(w http.ResponseWriter, r *http.Request) bool {
    cookie, err := r.Cookie("session")

    if err != nil || cookie.Value != "loggedin" {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return false
    }
    return true
}

func session(w http.ResponseWriter, r *http.Request){
    if r.URL.Path == "/Hello"{
        w.Header().Set("Content-Type", "text/html")
        tmplt, err := template.ParseFiles("static/hello.html")
        if err != nil{
            http.Error(w, "1Internal Server Error", http.StatusInternalServerError)
            return
        }

        err = tmplt.Execute(w, nil)
        if err != nil{
            http.Error(w, "2Internal Server Error", http.StatusInternalServerError)
            return 
        }
    }else if r.URL.Path == "/"{
        if r.Method == http.MethodGet {
        tmplt, err:= template.ParseFiles("static/login.html")
        if err != nil {
            http.Error(w, "1Internal Error", http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "text/html; charset=utf-8")
        err = tmplt.Execute(w, nil)
        if err != nil {
            http.Error(w, "2Internal Error", http.StatusInternalServerError)
            return 
        }
    }else if r.Method == http.MethodPost {
        login(w, r)
    }
    }else if r.URL.Path == "/Home"{
        if !requireLogin(w, r) {
            return
        }

        tmplt, err := template.ParseFiles("static/home.html")

        if err != nil{
            http.Error(w, "1Internal Server Error", http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "text/html; charset=utf-8")
        err = tmplt.Execute(w, nil)
        if err != nil{
            http.Error(w, "2Internal Server Error", http.StatusInternalServerError)
            return
        }
    }else if r.URL.Path == "/Calc"{
        if  !requireLogin(w, r) {
            return
        }

        if r.Method == http.MethodGet{
            tmplt, err := template.ParseFiles("static/calc.html")
            if err != nil{
                http.Error(w, "1Internal Server Error", http.StatusInternalServerError)
                return
            }

            w.Header().Set("Content-Type", "text/html; charset=utf-8")
            err = tmplt.Execute(w, nil)
            if err != nil{
                http.Error(w, "2Internal Server Error", http.StatusInternalServerError)
            }
        }else if r.Method == http.MethodPost{
            if requireLogin(w, r){
                return 
            }

            r.ParseForm()
            num1 := r.FormValue("num1")
            num2 := r.FormValue("num2")
            op := r.FormValue("operation")
            answer, err := Calc(num1, num2, op)
            if err != nil { answer = 0 }
            
            finalText := fmt.Sprintf("You entered %s and %s with operation %s = %f", num1, num2, op, answer)

            w.Header().Set("Content-Type", "text/html")
            tmplt, err := template.ParseFiles("static/calcResult.html")
            if err != nil{
                http.Error(w, "1Internal Error", http.StatusInternalServerError)
                return
            }

            err = tmplt.Execute(w, finalText)
            if err != nil{
                http.Error(w, "2Internal Erro", http.StatusInternalServerError)
                return 
            }
        }else {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    }else if r.URL.Path == "/Games"{
        if !requireLogin(w, r) {
            return
        }

        if r.Method == http.MethodGet{
            tmplt, err := template.ParseFiles("static/games.html")
            if err != nil{
                http.Error(w, "1Games Internal Error", http.StatusInternalServerError)
                return
            }

            w.Header().Set("content-Type", "text/html")
            err = tmplt.Execute(w, nil)
            if err != nil{
                http.Error(w, "2Games Internal Error", http.StatusInternalServerError)
                return 
            }
        }
    }else if r.URL.Path == "/Games/TicTacToe"{
        if !requireLogin(w, r){
            return
        }

        if r.Method == http.MethodGet{
            tmplt, err := template.ParseFiles("static/TicTacToe.html")
            if err != nil{
                http.Error(w, "1TicTacToe Internal Error", http.StatusInternalServerError)
                return
            }

            w.Header().Set("Content-type", "text/html")
            err = tmplt.Execute(w, nil)
            if err != nil{
                //http.Error(w, "2TicTacToe Internal Error", http.StatusInternalServerError)
                fmt.Println("Failed")
                return 
            }
        }
    }
}


func main(){
    fmt.Println("We Ball")

    http.HandleFunc("/", session)
    fmt.Println("Listening on Port: 8080")

    //Enable File server
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Server Failed", err)
    }
}
