package main

import (
	"bytes"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net"
	"net/http"
	"os/exec"
	"time"
)



func main() {
	checkCommand()
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
	l, err := net.Listen("tcp4", ":12345")
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.Serve(l, nil))


}
func checkCommand()  {
	//cmd := exec.Command("python3", "-m", "http.server")
	//fmt.Println("Check this railway command", command)
	cmd := exec.Command("./mycheck", "11")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	// Wait for the process to finish or kill it after a timeout:
	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()
	select {
	case <-time.After(10 * time.Millisecond):
		if err := cmd.Process.Kill(); err != nil {
			log.Fatal("failed to kill process: ", err)
		}
		log.Println("process killed as timeout reached")
	case err := <-done:
		if err != nil {
			log.Fatalf("process finished with error = %v", err)
		}
		fmt.Println("Check return %s" , out.String())
		log.Print("process finished successfully")
	}

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
