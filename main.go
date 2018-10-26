package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"log"
	"net/http"
)

var cache redis.Conn

func main() {
	go DownloadFirmware()
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/welcome", Welcome)
	http.HandleFunc("/signup", Signup)
	//http.HandleFunc("/admin", Admin)
	http.HandleFunc("/traincommand", GetTrainCommand)
	http.HandleFunc("/railwaycommand", GetRailwayCommand)
	http.HandleFunc("/settraincommand", SetTrainCommand)
	http.HandleFunc("/setrailwaycommand", SetRailwayCommand)
	//http.HandleFunc("/upload", Upload)
	//http.HandleFunc("/refresh", Refresh)
	// start the server on port 8000
	log.Fatal(http.ListenAndServe(":12345", nil))

}

func TestDb(){
	db, err := sql.Open("mysql", "root:Remidolov12345@@/test?charset=utf8")
	if err != nil{
		return
	}
	rows, err := db.Query("SELECT * FROM test")
	if err != nil{
		return
	}
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		if err != nil{
			return
		}
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}
}
