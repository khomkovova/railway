package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net"
	"net/http"
)



func main() {
	go DownloadFirmware()
	fs := http.FileServer(http.Dir("public/js"))
	http.Handle("/js/", http.StripPrefix("/js/", fs))
	http.HandleFunc("/", IndexPage)
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/welcome", Welcome)
	http.HandleFunc("/signup", Signup)
	http.HandleFunc("/admin", Admin)
	http.HandleFunc("/api/signin", ApiSignin)
	http.HandleFunc("/api/signup", ApiSignup)
	http.HandleFunc("/api/railwayinfo", GetRailwayInfo)
	http.HandleFunc("/api/traincommand", GetTrainCommand)
	http.HandleFunc("/api/railwaycommand", GetRailwayCommand)
	http.HandleFunc("/api/settraincommand", SetTrainCommand)
	http.HandleFunc("/api/setrailwaycommand", SetRailwayCommand)
	l, err := net.Listen("tcp4", ":80")
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.Serve(l, nil))


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
