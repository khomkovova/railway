package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"database/sql"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)


var PrivateKey, err2 = rsa.GenerateKey(rand.Reader, 2048)
var PublicKey = &PrivateKey.PublicKey

type CredentialsSignin struct {
	Password string `json:"password"`
	Username string `json:"username"`
}
type Commands struct {
	Speed int `json:"speed"`
	Direction int `json:"direction"`
}
type CommandsRailway struct {
	Firstswitch int `json:"firstswitch"`
	Secondswitch int `json:"secondswitch"`
}

type CommandsAll struct {
	Train Commands `json:"train"`
	Railway CommandsRailway `json:"railway"`
}
type CredentialsRegistration struct {
	Password string `json:"password"`
	Username string `json:"username"`
	Email string `json:"email"`
}


var commands Commands
var commandsRailway  CommandsRailway
var commandsAll CommandsAll


func IndexPage(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("public/index.html")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write([]byte(data))
}

func Signin(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadFile("public/signin.html")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Write([]byte(data))

}
func ApiSignin(w http.ResponseWriter, r *http.Request)  {
	var creds CredentialsSignin

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var Db, err2 = sql.Open("mysql", "root:Remidolov12345@@/railway?charset=utf8")
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var username string
	err = Db.QueryRow("SELECT username FROM users WHERE username=? AND password=?", creds.Username, creds.Password).Scan(&username)
	if err != nil || username == ""{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	cookie, err3 := encodeCookie(creds.Username)
	if err3 == false{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
}

func Signup(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("public/signup.html")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write([]byte(data))
}

func ApiSignup(w http.ResponseWriter, r *http.Request)  {

	var creds CredentialsRegistration
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("Signup with =", creds)
	var Db, err2 = sql.Open("mysql", "root:Remidolov12345@@/railway?charset=utf8")
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := Db.Query("INSERT INTO users (username,password)" + "VALUES ('" + creds.Username + "','" + creds.Password + "')" )
	if err != nil{
		id.Close()
		w.Write([]byte("The username using or this credentials is not correct"))
		return
	}

	w.Write([]byte("You success registered"))
	w.WriteHeader(http.StatusOK)

}

func Admin(w http.ResponseWriter, r *http.Request)  {
	user := decodeCookie(r)
	if user == ""{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if user != "admin"{
		http.Redirect(w, r, "/welcome", http.StatusSeeOther)
		return
	}
		data, err := ioutil.ReadFile("public/admin.html")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Write([]byte(data))

}

func SetTrainCommand(w http.ResponseWriter, r *http.Request)  {
	user := decodeCookie(r)
	if user == ""{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if user != "admin"{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	jd := json.NewDecoder(r.Body)

	err := jd.Decode(&commands)
	if err !=nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if (commands.Speed < 0 || commands.Speed > 8) || (commands.Direction != 0 && commands.Direction != 1) {
		w.Write([]byte("Your commands are bad"))
		w.WriteHeader(http.StatusOK)
		return

	}
	fmt.Println("Train commands = ", commands)
	w.Write([]byte("Your commands were send"))
	w.WriteHeader(http.StatusOK)



}

func SetRailwayCommand(w http.ResponseWriter, r *http.Request)  {
	user := decodeCookie(r)
	if user == ""{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if user != "admin"{
		http.Redirect(w, r, "/welcome", http.StatusSeeOther)
		return
	}

	jd := json.NewDecoder(r.Body)
	var commandsRailwayTest CommandsRailway
	err := jd.Decode(&commandsRailwayTest)
	if err !=nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("Railway commands = ", commandsRailwayTest)
	if (commandsRailwayTest.Secondswitch != 0 && commandsRailwayTest.Secondswitch != 1) || (commandsRailwayTest.Firstswitch != 0 && commandsRailwayTest.Firstswitch != 1) {
		w.Write([]byte("Your commands are bad"))
		w.WriteHeader(http.StatusOK)
		return

	}
	if !checkCommandRailway(commandsRailwayTest){
		w.Write([]byte("Railway upgrading!!!<br>Change direction can be dangerous.<br>You could write to support team for solving this problem"))
		w.WriteHeader(http.StatusOK)
		return
	}
	commandsRailway = commandsRailwayTest
	if err !=nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write([]byte("Your commands send"))
	w.WriteHeader(http.StatusOK)



}

func Welcome(w http.ResponseWriter, r *http.Request) {
	if !getStatusUser(r){
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
		data, err := ioutil.ReadFile("public/welcome.html")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Write([]byte(data))

}

func GetRailwayInfo(w http.ResponseWriter, r *http.Request)  {
	if !getStatusUser(r){
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	commandsAll.Train = commands
	commandsAll.Railway = commandsRailway
	data, err := json.Marshal(commandsAll)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write([]byte(data))
}

func GetTrainCommand(w http.ResponseWriter, r *http.Request){
	var token string
	for _, cookie := range r.Cookies() {
		token = cookie.Value
	}
	if token != "train1234567890"{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.Write([]byte(string(commands.Speed + '0')+string(commands.Direction + '0')))
}

func GetRailwayCommand(w http.ResponseWriter, r *http.Request){
	var token string
	for _, cookie := range r.Cookies() {
		token = cookie.Value
	}
	if token != "railway1234567890"{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.Write([]byte(string(commandsRailway.Firstswitch + '0')+string(commandsRailway.Secondswitch + '0')))


}

func checkCommandRailway(command CommandsRailway) bool {
	return true
	//cmd := exec.Command("python3", "-m", "http.server")
	commandStr := "./mycheck" + " " + string(command.Firstswitch + '0') + string(command.Secondswitch + '0')
	fmt.Println(commandStr)
	cmd := exec.Command("./mycheck", string(command.Firstswitch + '0') + string(command.Secondswitch + '0'))
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
			return false
		}
		log.Println("process killed as timeout reached")
		return false
	case err := <-done:
		if err != nil {
			log.Fatalf("process finished with error = %v", err)
			return false
		}
		log.Print("process finished successfully")
		fmt.Println("Check return %s" , out.String())
		if out.String() != "True"{
			return false
		}
		return true
	}

}
func DownloadFirmware() {
	time.Sleep(5 * time.Second)
	fmt.Println("Download start")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://35.211.111.79:80/downloadfirmware", nil)
	if err != nil{
		DownloadFirmware()
		return
	}
	c := http.Cookie{Name:"token",Value:"server1234567890"}
	req.AddCookie(&c)
	resp, err := client.Do(req)
	if err != nil {
		DownloadFirmware()
		return
	}

	f, err := os.OpenFile("./"+"mycheck", os.O_WRONLY|os.O_CREATE, 0766)
	if err != nil {
		DownloadFirmware()
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		DownloadFirmware()
		return
	}
	f.Write(body)
	f.Close()
	fmt.Println("Download firmware success")
	DownloadFirmware()
}

func decodeCookie(r *http.Request) string {
	var token string
	for _, cookie := range r.Cookies() {
		token = cookie.Value
	}
	if token == ""{
		return ""
	}
	sDec, _ := b64.StdEncoding.DecodeString(token)
	label := []byte("")
	hash := sha256.New()
	plainText, err := rsa.DecryptOAEP(hash, rand.Reader, PrivateKey, sDec, label)
	if err != nil{
		return ""
	}
	user := string(plainText)
	fmt.Println("Decode token = ",user)
	return user
}

func encodeCookie(user string) (http.Cookie, bool)  {
	message := []byte(user)
	label := []byte("")
	hash := sha256.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, PublicKey, message, label)
	if err != nil{
		return http.Cookie{}, false
	}
	sEnc := b64.StdEncoding.EncodeToString(ciphertext)
	cookie := http.Cookie{Name:"token", Value:sEnc}
	cookie.Path = "/"
	return cookie, true

}

func getStatusUser(r *http.Request) bool {
	user := decodeCookie(r)
	if user == ""{
		return false
	}
	var Db, err2 = sql.Open("mysql", "root:Remidolov12345@@/railway?charset=utf8")
	if err2 != nil {
		return false
	}
	var username string
	err := Db.QueryRow("SELECT username FROM users WHERE username=?", user).Scan(&username)
	if err != nil || username == ""{
		return false
	}
	return true
}

