package main

import (
        "database/sql"
        "fmt"
        "net/http"
        _ "github.com/go-sql-driver/mysql"
)

//User is struct
type User struct {
        ID   int
        Name string
}

func main() {
        http.HandleFunc("/",HelloHandler)
        http.ListenAndServe(":8080",nil)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) 
	w.Write([]byte("{

        db, err := sql.Open("mysql", "root:Shun!1007@tcp(127.0.0.1:3306)/sample_db")
        if err != nil {
                panic(err.Error())
        }
        defer db.Close()

        rows, err := db.Query("SELECT * FROM users")
        if err != nil {
                panic(err.Error())
        }
        defer rows.Close()

        for rows.Next() {
                var user User
                err := rows.Scan(&user.ID, &user.Name)
                if err != nil {
                        panic(err.Error())
                }
                fmt.Println(user.ID, user.Name)
        }

        err = rows.Err()
        if err != nil {
                panic(err.Error())
        }"))
