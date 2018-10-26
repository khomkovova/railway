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

func Signin(w http.ResponseWriter, r *http.Request) {
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
	http.Redirect(w, r, "/welcome", http.StatusSeeOther)
}

func Signup(w http.ResponseWriter, r *http.Request)  {
	var creds CredentialsRegistration
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(creds)
	var Db, err2 = sql.Open("mysql", "root:Remidolov12345@@/railway?charset=utf8")
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := Db.Query("INSERT INTO users (username,password)" + "VALUES ('" + creds.Username + "','" + creds.Password + "')" )
	fmt.Println(id)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//Db.Close()
	//cookie, err3 := encodeCookie(creds.Username)
	//if err3 == false{
	//	fmt.Println("error")
	//	return
	//}
	//http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/signin", http.StatusSeeOther)

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

	jd := json.NewDecoder(r.Body)

	err := jd.Decode(&commandsRailway)
	jd.DisallowUnknownFields()
	err2 := jd.Decode(&commands)
	fmt.Println(commandsRailway)
	fmt.Println(commands)
	if err != nil && err2 !=nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("command=", commands)
	w.WriteHeader(http.StatusOK)



}

func SetTrainCommand(w http.ResponseWriter, r *http.Request)  {
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

	err := jd.Decode(&commands)
	fmt.Println(commands)
	if err !=nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("command=", commands)
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
	if !checkCommandRailway(commandsRailwayTest){
		w.Write([]byte("Now upgrade railway and choose direction can be dangeros we can write to support team https://supportrailway.com"))
		w.WriteHeader(http.StatusOK)
		return
	}
	commandsRailway = commandsRailwayTest
	fmt.Println(commandsRailway)
	if err !=nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("command=", commandsRailway)
	w.WriteHeader(http.StatusOK)



}

func Welcome(w http.ResponseWriter, r *http.Request) {
	user := decodeCookie(r)
	if user == ""{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	var Db, err2 = sql.Open("mysql", "root:Remidolov12345@@/railway?charset=utf8")
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var username string
	err := Db.QueryRow("SELECT username FROM users WHERE username=?", user).Scan(&username)
	if err != nil || username == ""{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	commandsAll.Train = commands
	commandsAll.Railway = commandsRailway
	data, err := json.Marshal(commandsAll)
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
	cmd := exec.Command("./mycheck", string(command.Firstswitch +'0')+string(command.Secondswitch +'0') )
	//cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error exc commands")
		return false
	}
	fmt.Println("Check return %s" , out.String())
	if out.String() != "True"{
		return false
	}
	return true

}
func DownloadFirmware() {
	time.Sleep(5 * time.Second)
	fmt.Println("Run")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://127.0.0.1:12346/downloadfirmware", nil)
	if err != nil{
		DownloadFirmware()
		return
	}
	c := http.Cookie{Name:"token",Value:"server1234567890"}
	req.AddCookie(&c)
	resp, err := client.Do(req)
	fmt.Println(resp.Body)
	//resp, err := http.Get("http://127.0.0.1/downloadfirmware")

	f, err := os.OpenFile("./"+"mycheck", os.O_WRONLY|os.O_CREATE, 0766)
	if err != nil {
		//fmt.Println("124")
		DownloadFirmware()
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		DownloadFirmware()
		return
	}
	fmt.Println(body)
	f.Write(body)
	f.Close()
	DownloadFirmware()
}

func decodeCookie(r *http.Request) string {
	var token string
	for _, cookie := range r.Cookies() {
		token = cookie.Value
	}
	if token == ""{
		fmt.Println("error")
		return ""
	}
	sDec, _ := b64.StdEncoding.DecodeString(token)
	label := []byte("")
	hash := sha256.New()
	plainText, err := rsa.DecryptOAEP(hash, rand.Reader, PrivateKey, sDec, label)
	if err != nil{
		fmt.Println("error")
		return ""
	}
	user := string(plainText)
	fmt.Println("token=",token)
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
	return cookie, true

}

