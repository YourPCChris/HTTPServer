package main

import (
  "net/http"
)


func Login(w http.ResponseWriter, r *http.Request){
    if r.Method == http.MethodPost {
        r.ParseForm()
        username := r.FormValue("username")
        password := r.FormValue("password")
         

        if userPresent, _:= checkUser(username, string(password)); userPresent{
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

 func RequireLogin(w http.ResponseWriter, r *http.Request) bool{
     cookie, err := r.Cookie("session")

     if err != nil || cookie.Value != "loggedin"{
         http.Redirect(w, r, "/", http.StatusSeeOther)
         return false
     }
     return true;
}

