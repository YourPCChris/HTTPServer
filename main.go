
package main 

import 
(
    "fmt"
    "net/http"
    "html/template"
)

func session(w http.ResponseWriter, r *http.Request){
    //fmt.Fprintf(w, "Your Request: %s\n", r.URL.Path)
    //fmt.Fprintf(w, "Your Method: %s\n", r.Method)

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
        //Games 
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
